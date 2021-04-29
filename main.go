package main

import (
	"github.com/ChimeraCoder/gojson"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	jsonData, err := ioutil.ReadFile("./jsonData.json")
	if err != nil {
		log.Fatal(err)
	}

	stringJson := string(jsonData)
	//
	//var jsonMap map[string]interface{}
	//
	//json.Unmarshal(jsonData, &jsonMap)
	//
	//fmt.Println("Getting JSON Structure and Data")
	//fmt.Printf("%v", jsonMap)

	if output, err := gojson.Generate(strings.NewReader(stringJson), gojson.ParseJson, "jsonStruct", "main", []string{"json"}, false, true); err != nil {
		log.Fatal(err)
	} else {
		err := ioutil.WriteFile("models.go", output, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
