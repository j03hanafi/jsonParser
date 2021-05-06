package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// Setting up log file
	// set permission to read/write log file
	// read/write to existing log file, if there is none it will create new log file
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Found error in log ", err)
	}
	log.SetOutput(file)

	// Setting up HTTP Listener and Handler
	// router will handle any request at any endpoint available in server()
	router := pathHandler()
	// listen to specific address and handler
	address := "localhost:6010"
	log.Println("Server started at", address)
	server := http.ListenAndServe(address, router)
	if server != nil {
		log.Fatal(server.Error())
	}
}

// return HTTP handler
func pathHandler() *mux.Router {

	// create new handler instance
	router := mux.NewRouter()

	// Endpoints, Handler function, and HTTP request Method
	router.HandleFunc("/iso20022", parseIso).Methods("POST")

	return router
}

func parseIso(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	ipReq := getIP(r)
	log.Printf("[Conn: %v. Time: %v. Elapsed: %.6fs] Received new ISO20022 Request\n", ipReq, time.Now().Format("15:04:05"), time.Since(start).Seconds())

	// Get request body JSON
	body, _ := ioutil.ReadAll(r.Body)
	//fmt.Printf("%+v", string(body))
	var request Iso20022
	var response Response

	err := json.Unmarshal(body, &request)
	if err != nil {
		response.Message = fmt.Sprintf("Error unmarshal JSON: %s", err.Error())
		log.Printf(response.Message)
		responseFormatter(w, response, http.StatusInternalServerError)
		return
	}
	fmt.Printf("%+v\n", request)

	doc, err := json.MarshalIndent(request.BusMsg.Document, "", "  ")
	if err != nil {
		response.Message = fmt.Sprintf("Error MarshalIndent JSON: %s", err.Error())
		log.Printf(response.Message)
		responseFormatter(w, response, http.StatusInternalServerError)
		return
	}
	//log.Printf("\n\nDocument: %s\n\n", string(doc))

	// save as file
	filename := fmt.Sprintf("parsed/%v@%v", ipReq, time.Now().Format("15:04:05"))
	CreateFile(filename, string(doc))

	response.Message = "Parsing Success"
	responseFormatter(w, response, http.StatusOK)

}

// Return ip address for client request
func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

// Response formatter
func responseFormatter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Create file for request/response
func CreateFile(fileName string, content string) string {

	log.Println("Creating new file")

	if !strings.Contains(fileName, ".json") {
		fileName += ".json"
	}

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}

	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		log.Fatalf("Failed writing to file: %s", err)
	}

	log.Println("File created!")
	return fileName
}
