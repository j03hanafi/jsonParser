package main

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type Response struct {
	Message string `json:"Message"`
}

type Iso20022 struct {
	BusMsg BusMsg `json:"BusMsg"`
}

type BusMsg struct {
	AppHdr   AppHdr   `json:"AppHdr"`
	Document Document `json:"Document"`
}

type AppHdr struct {
	BizMsgIdr string `json:"BizMsgIdr"`
	MsgDefIdr string `json:"MsgDefIdr"`
	CreDt     string `json:"CreDt"`
}

type AccountIdentification4Choice struct {
	IBAN *IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 IBAN,omitempty" json:"IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Othr,omitempty" json:"Othr,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    *ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type ActiveCurrencyAndAmount struct {
	Value float64             `xml:",chardata" json:"Value,string"`
	Ccy   *ActiveCurrencyCode `xml:"Ccy,attr" json:"Ccy"`
}

// ActiveCurrencyCode Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                       `xml:",chardata" json:"Value,string"`
	Ccy   *ActiveOrHistoricCurrencyCode `xml:"Ccy,attr" json:"Ccy"`
}

// ActiveOrHistoricCurrencyCode Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// AddressType2Code May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    *AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

// AnyBICDec2014Identifier Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// BICFIDec2014Identifier Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId *FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FinInstnId" json:"FinInstnId"`
	BrnchId    *BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 BrnchId,omitempty" json:"BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      *Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Id,omitempty" json:"Id,omitempty"`
	LEI     *LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 LEI,omitempty" json:"LEI,omitempty"`
	Nm      *Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr *PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PstlAdr,omitempty" json:"PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   *AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Id" json:"Id"`
	Tp   *CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	Ccy  *ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Ccy,omitempty" json:"Ccy,omitempty"`
	Nm   *Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Nm,omitempty" json:"Nm,omitempty"`
	Prxy *ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prxy,omitempty" json:"Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    *ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type CategoryPurpose1Choice struct {
	Cd    *ExternalCategoryPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

// ChargeBearerType1Code May be one of DEBT, CRED, SHAR, SLEV
type ChargeBearerType1Code string

type Charges7 struct {
	Amt *ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Amt" json:"Amt"`
	Agt *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Agt" json:"Agt"`
}

// ClearingChannel2Code May be one of RTGS, RTNS, MPNS, BOOK
type ClearingChannel2Code string

type ClearingSystemIdentification2Choice struct {
	Cd    *ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd"`
	Prtry *Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry"`
}

type ClearingSystemIdentification3Choice struct {
	Cd    *ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ClrSysId,omitempty" json:"ClrSysId,omitempty"`
	MmbId    *Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 MmbId" json:"MmbId"`
}

type Contact4 struct {
	NmPrfx    *NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 NmPrfx,omitempty" json:"NmPrfx,omitempty"`
	Nm        *Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Nm,omitempty" json:"Nm,omitempty"`
	PhneNb    *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PhneNb,omitempty" json:"PhneNb,omitempty"`
	MobNb     *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 MobNb,omitempty" json:"MobNb,omitempty"`
	FaxNb     *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FaxNb,omitempty" json:"FaxNb,omitempty"`
	EmailAdr  *Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 EmailAdr,omitempty" json:"EmailAdr,omitempty"`
	EmailPurp *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 EmailPurp,omitempty" json:"EmailPurp,omitempty"`
	JobTitl   *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 JobTitl,omitempty" json:"JobTitl,omitempty"`
	Rspnsblty *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Rspnsblty,omitempty" json:"Rspnsblty,omitempty"`
	Dept      *Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Dept,omitempty" json:"Dept,omitempty"`
	Othr      []*OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Othr,omitempty" json:"Othr,omitempty"`
	PrefrdMtd *PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PrefrdMtd,omitempty" json:"PrefrdMtd,omitempty"`
}

// CountryCode Must match the pattern [A-Z]{2,2}
type CountryCode string

// CreditDebitCode May be one of CRDT, DBIT
type CreditDebitCode string

type CreditTransferMandateData1 struct {
	MndtId       *Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 MndtId,omitempty" json:"MndtId,omitempty"`
	Tp           *MandateTypeInformation2   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	DtOfSgntr    *ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DtOfSgntr,omitempty" json:"DtOfSgntr,omitempty"`
	DtOfVrfctn   *ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DtOfVrfctn,omitempty" json:"DtOfVrfctn,omitempty"`
	ElctrncSgntr *Max10KBinary              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ElctrncSgntr,omitempty" json:"ElctrncSgntr,omitempty"`
	FrstPmtDt    *ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FrstPmtDt,omitempty" json:"FrstPmtDt,omitempty"`
	FnlPmtDt     *ISODate                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FnlPmtDt,omitempty" json:"FnlPmtDt,omitempty"`
	Frqcy        *Frequency36Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Frqcy,omitempty" json:"Frqcy,omitempty"`
	Rsn          *MandateSetupReason1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Rsn,omitempty" json:"Rsn,omitempty"`
}

type CreditTransferTransaction43 struct {
	PmtId             *PaymentIdentification13                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PmtId" json:"PmtId"`
	PmtTpInf          *PaymentTypeInformation28                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	IntrBkSttlmAmt    *ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 IntrBkSttlmAmt" json:"IntrBkSttlmAmt"`
	IntrBkSttlmDt     *ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	SttlmPrty         *Priority3Code                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SttlmPrty,omitempty" json:"SttlmPrty,omitempty"`
	SttlmTmIndctn     *SettlementDateTimeIndication1                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SttlmTmIndctn,omitempty" json:"SttlmTmIndctn,omitempty"`
	SttlmTmReq        *SettlementTimeRequest2                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SttlmTmReq,omitempty" json:"SttlmTmReq,omitempty"`
	AccptncDtTm       *ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 AccptncDtTm,omitempty" json:"AccptncDtTm,omitempty"`
	PoolgAdjstmntDt   *ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PoolgAdjstmntDt,omitempty" json:"PoolgAdjstmntDt,omitempty"`
	InstdAmt          *ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstdAmt,omitempty" json:"InstdAmt,omitempty"`
	XchgRate          float64                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 XchgRate,omitempty" json:"XchgRate,omitempty"`
	ChrgBr            *ChargeBearerType1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ChrgBr" json:"ChrgBr"`
	ChrgsInf          []*Charges7                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ChrgsInf,omitempty" json:"ChrgsInf,omitempty"`
	MndtRltdInf       *CreditTransferMandateData1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 MndtRltdInf,omitempty" json:"MndtRltdInf,omitempty"`
	PrvsInstgAgt1     *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PrvsInstgAgt1,omitempty" json:"PrvsInstgAgt1,omitempty"`
	PrvsInstgAgt1Acct *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PrvsInstgAgt1Acct,omitempty" json:"PrvsInstgAgt1Acct,omitempty"`
	PrvsInstgAgt2     *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PrvsInstgAgt2,omitempty" json:"PrvsInstgAgt2,omitempty"`
	PrvsInstgAgt2Acct *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PrvsInstgAgt2Acct,omitempty" json:"PrvsInstgAgt2Acct,omitempty"`
	PrvsInstgAgt3     *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PrvsInstgAgt3,omitempty" json:"PrvsInstgAgt3,omitempty"`
	PrvsInstgAgt3Acct *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PrvsInstgAgt3Acct,omitempty" json:"PrvsInstgAgt3Acct,omitempty"`
	InstgAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstdAgt,omitempty" json:"InstdAgt,omitempty"`
	IntrmyAgt1        *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 IntrmyAgt1,omitempty" json:"IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct    *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 IntrmyAgt1Acct,omitempty" json:"IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2        *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 IntrmyAgt2,omitempty" json:"IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct    *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 IntrmyAgt2Acct,omitempty" json:"IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3        *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 IntrmyAgt3,omitempty" json:"IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct    *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 IntrmyAgt3Acct,omitempty" json:"IntrmyAgt3Acct,omitempty"`
	UltmtDbtr         *PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	InitgPty          *PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InitgPty,omitempty" json:"InitgPty,omitempty"`
	Dbtr              *PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Dbtr" json:"Dbtr"`
	DbtrAcct          *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DbtrAcct,omitempty" json:"DbtrAcct,omitempty"`
	DbtrAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DbtrAgt" json:"DbtrAgt"`
	DbtrAgtAcct       *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DbtrAgtAcct,omitempty" json:"DbtrAgtAcct,omitempty"`
	CdtrAgt           *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdtrAgt" json:"CdtrAgt"`
	CdtrAgtAcct       *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdtrAgtAcct,omitempty" json:"CdtrAgtAcct,omitempty"`
	Cdtr              *PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cdtr" json:"Cdtr"`
	CdtrAcct          *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdtrAcct,omitempty" json:"CdtrAcct,omitempty"`
	UltmtCdtr         *PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 UltmtCdtr,omitempty" json:"UltmtCdtr,omitempty"`
	InstrForCdtrAgt   []*InstructionForCreditorAgent3               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstrForCdtrAgt,omitempty" json:"InstrForCdtrAgt,omitempty"`
	InstrForNxtAgt    []*InstructionForNextAgent1                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstrForNxtAgt,omitempty" json:"InstrForNxtAgt,omitempty"`
	Purp              *Purpose2Choice                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Purp,omitempty" json:"Purp,omitempty"`
	RgltryRptg        []*RegulatoryReporting3                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RgltryRptg,omitempty" json:"RgltryRptg,omitempty"`
	Tax               *TaxInformation8                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tax,omitempty" json:"Tax,omitempty"`
	RltdRmtInf        []*RemittanceLocation7                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RltdRmtInf,omitempty" json:"RltdRmtInf,omitempty"`
	RmtInf            *RemittanceInformation16                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RmtInf,omitempty" json:"RmtInf,omitempty"`
	SplmtryData       []*SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type CreditorReferenceInformation2 struct {
	Tp  *CreditorReferenceType2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	Ref *Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Ref,omitempty" json:"Ref,omitempty"`
}

type CreditorReferenceType1Choice struct {
	Cd    *DocumentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type CreditorReferenceType2 struct {
	CdOrPrtry *CreditorReferenceType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdOrPrtry" json:"CdOrPrtry,omitempty"`
	Issr      *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Issr,omitempty" json:"Issr"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     *ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 BirthDt" json:"BirthDt"`
	PrvcOfBirth *Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PrvcOfBirth,omitempty" json:"PrvcOfBirth,omitempty"`
	CityOfBirth *Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CityOfBirth" json:"CityOfBirth"`
	CtryOfBirth *CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CtryOfBirth" json:"CtryOfBirth"`
}

type DatePeriod2 struct {
	FrDt *ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FrDt" json:"FrDt"`
	ToDt *ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ToDt" json:"ToDt"`
}

type DiscountAmountAndType1 struct {
	Tp  *DiscountAmountType1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Amt" json:"Amt"`
}

type DiscountAmountType1Choice struct {
	Cd    *ExternalDiscountAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type Document struct {
	FIToFICstmrCdtTrf *FIToFICustomerCreditTransferV09 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FIToFICstmrCdtTrf" json:"FIToFICstmrCdtTrf"`
}

type DocumentAdjustment1 struct {
	Amt       *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Amt" json:"Amt"`
	CdtDbtInd *CreditDebitCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdtDbtInd,omitempty" json:"CdtDbtInd,omitempty"`
	Rsn       *Max4Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Rsn,omitempty" json:"Rsn,omitempty"`
	AddtlInf  *Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 AddtlInf,omitempty" json:"AddtlInf,omitempty"`
}

type DocumentLineIdentification1 struct {
	Tp     *DocumentLineType1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	Nb     *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Nb,omitempty" json:"Nb,omitempty"`
	RltdDt *ISODate           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RltdDt,omitempty" json:"RltdDt,omitempty"`
}

type DocumentLineInformation1 struct {
	Id   []*DocumentLineIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Id" json:"Id"`
	Desc *Max2048Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Desc,omitempty" json:"Desc,omitempty"`
	Amt  *RemittanceAmount3             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Amt,omitempty" json:"Amt,omitempty"`
}

type DocumentLineType1 struct {
	CdOrPrtry *DocumentLineType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdOrPrtry" json:"CdOrPrtry"`
	Issr      *Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Issr,omitempty" json:"Issr,omitempty"`
}

type DocumentLineType1Choice struct {
	Cd    *ExternalDocumentLineType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

// DocumentType3Code May be one of RADM, RPIN, FXDR, DISP, PUOR, SCOR
type DocumentType3Code string

// DocumentType6Code May be one of MSIN, CNFA, DNFA, CINV, CREN, DEBN, HIRI, SBIN, CMCN, SOAC, DISP, BOLD, VCHR, AROI, TSUT, PUOR
type DocumentType6Code string

// Exact2NumericText Must match the pattern [0-9]{2}
type Exact2NumericText string

// Exact4AlphaNumericText Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// ExternalAccountIdentification1Code May be no more than 4 items long
type ExternalAccountIdentification1Code string

// ExternalCashAccountType1Code May be no more than 4 items long
type ExternalCashAccountType1Code string

// ExternalCashClearingSystem1Code May be no more than 3 items long
type ExternalCashClearingSystem1Code string

// ExternalCategoryPurpose1Code May be no more than 4 items long
type ExternalCategoryPurpose1Code string

// ExternalClearingSystemIdentification1Code May be no more than 5 items long
type ExternalClearingSystemIdentification1Code string

// ExternalCreditorAgentInstruction1Code May be no more than 4 items long
type ExternalCreditorAgentInstruction1Code string

// ExternalDiscountAmountType1Code May be no more than 4 items long
type ExternalDiscountAmountType1Code string

// ExternalDocumentLineType1Code May be no more than 4 items long
type ExternalDocumentLineType1Code string

// ExternalFinancialInstitutionIdentification1Code May be no more than 4 items long
type ExternalFinancialInstitutionIdentification1Code string

// ExternalGarnishmentType1Code May be no more than 4 items long
type ExternalGarnishmentType1Code string

// ExternalLocalInstrument1Code May be no more than 35 items long
type ExternalLocalInstrument1Code string

// ExternalMandateSetupReason1Code May be no more than 4 items long
type ExternalMandateSetupReason1Code string

// ExternalOrganisationIdentification1Code May be no more than 4 items long
type ExternalOrganisationIdentification1Code string

// ExternalPersonIdentification1Code May be no more than 4 items long
type ExternalPersonIdentification1Code string

// ExternalProxyAccountType1Code May be no more than 4 items long
type ExternalProxyAccountType1Code string

// ExternalPurpose1Code May be no more than 4 items long
type ExternalPurpose1Code string

// ExternalServiceLevel1Code May be no more than 4 items long
type ExternalServiceLevel1Code string

// ExternalTaxAmountType1Code May be no more than 4 items long
type ExternalTaxAmountType1Code string

type FIToFICustomerCreditTransferV09 struct {
	GrpHdr      *GroupHeader93                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 GrpHdr" json:"GrpHdr"`
	CdtTrfTxInf []*CreditTransferTransaction43 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdtTrfTxInf" json:"CdtTrfTxInf"`
	SplmtryData []*SupplementaryData1          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SplmtryData,omitempty" json:"SplmtryData,omitempty"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    *ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 BICFI,omitempty" json:"BICFI,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ClrSysMmbId,omitempty" json:"ClrSysMmbId,omitempty"`
	LEI         *LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 LEI,omitempty" json:"LEI,omitempty"`
	Nm          *Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PstlAdr,omitempty" json:"PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Othr,omitempty" json:"Othr,omitempty"`
}

type Frequency36Choice struct {
	Tp     *Frequency6Code      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	Prd    *FrequencyPeriod1    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prd,omitempty" json:"Prd,omitempty"`
	PtInTm *FrequencyAndMoment1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PtInTm,omitempty" json:"PtInTm,omitempty"`
}

// Frequency6Code May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, ADHO, INDA, FRTN
type Frequency6Code string

type FrequencyAndMoment1 struct {
	Tp     *Frequency6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp" json:"Tp"`
	PtInTm *Exact2NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PtInTm" json:"PtInTm"`
}

type FrequencyPeriod1 struct {
	Tp        *Frequency6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp" json:"Tp"`
	CntPerPrd float64         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CntPerPrd" json:"CntPerPrd"`
}

type Garnishment3 struct {
	Tp                *GarnishmentType1                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp" json:"Tp"`
	Grnshee           *PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Grnshee,omitempty" json:"Grnshee,omitempty"`
	GrnshmtAdmstr     *PartyIdentification135            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 GrnshmtAdmstr,omitempty" json:"GrnshmtAdmstr,omitempty"`
	RefNb             *Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RefNb,omitempty" json:"RefNb,omitempty"`
	Dt                *ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Dt,omitempty" json:"Dt,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RmtdAmt,omitempty" json:"RmtdAmt,omitempty"`
	FmlyMdclInsrncInd bool                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FmlyMdclInsrncInd,omitempty" json:"FmlyMdclInsrncInd,omitempty"`
	MplyeeTermntnInd  bool                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 MplyeeTermntnInd,omitempty" json:"MplyeeTermntnInd,omitempty"`
}

type GarnishmentType1 struct {
	CdOrPrtry *GarnishmentType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdOrPrtry" json:"CdOrPrtry"`
	Issr      *Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Issr,omitempty" json:"Issr,omitempty"`
}

type GarnishmentType1Choice struct {
	Cd    *ExternalGarnishmentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      *Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Id" json:"Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SchmeNm,omitempty" json:"SchmeNm,omitempty"`
	Issr    *Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Issr,omitempty" json:"Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      *Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Id" json:"Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SchmeNm,omitempty" json:"SchmeNm,omitempty"`
	Issr    *Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Issr,omitempty" json:"Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      *Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Id" json:"Id"`
	Issr    *Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Issr" json:"Issr"`
	SchmeNm *Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SchmeNm,omitempty" json:"SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Id" json:"Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SchmeNm,omitempty" json:"SchmeNm,omitempty"`
	Issr    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Issr,omitempty" json:"Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      *Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Id" json:"Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SchmeNm,omitempty" json:"SchmeNm,omitempty"`
	Issr    *Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Issr,omitempty" json:"Issr,omitempty"`
}

type GroupHeader93 struct {
	MsgId             *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 MsgId" json:"MsgId"`
	CreDtTm           *ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CreDtTm" json:"CreDtTm"`
	BtchBookg         bool                                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 BtchBookg,omitempty" json:"BtchBookg,omitempty"`
	NbOfTxs           *Max15NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 NbOfTxs" json:"NbOfTxs"`
	CtrlSum           float64                                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CtrlSum,omitempty" json:"CtrlSum,omitempty"`
	TtlIntrBkSttlmAmt *ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TtlIntrBkSttlmAmt,omitempty" json:"TtlIntrBkSttlmAmt,omitempty"`
	IntrBkSttlmDt     *ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 IntrBkSttlmDt,omitempty" json:"IntrBkSttlmDt,omitempty"`
	SttlmInf          *SettlementInstruction7                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SttlmInf" json:"SttlmInf"`
	PmtTpInf          *PaymentTypeInformation28                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PmtTpInf,omitempty" json:"PmtTpInf,omitempty"`
	InstgAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstgAgt,omitempty" json:"InstgAgt,omitempty"`
	InstdAgt          *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstdAgt,omitempty" json:"InstdAgt,omitempty"`
}

// IBAN2007Identifier Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

// Instruction4Code May be one of PHOA, TELA
type Instruction4Code string

type InstructionForCreditorAgent3 struct {
	Cd       *ExternalCreditorAgentInstruction1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	InstrInf *Max140Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstrInf,omitempty" json:"InstrInf,omitempty"`
}

type InstructionForNextAgent1 struct {
	Cd       *Instruction4Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	InstrInf *Max140Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstrInf,omitempty" json:"InstrInf,omitempty"`
}

// LEIIdentifier Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LocalInstrument2Choice struct {
	Cd    *ExternalLocalInstrument1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type MandateClassification1Choice struct {
	Cd    *MandateClassification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

// MandateClassification1Code May be one of FIXE, USGB, VARI
type MandateClassification1Code string

type MandateSetupReason1Choice struct {
	Cd    *ExternalMandateSetupReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type MandateTypeInformation2 struct {
	SvcLvl    *ServiceLevel8Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SvcLvl,omitempty" json:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 LclInstrm,omitempty" json:"LclInstrm,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CtgyPurp,omitempty" json:"CtgyPurp,omitempty"`
	Clssfctn  *MandateClassification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Clssfctn,omitempty" json:"Clssfctn,omitempty"`
}

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Max10Text May be no more than 10 items long
type Max10Text string

// Max128Text May be no more than 128 items long
type Max128Text string

// Max140Text May be no more than 140 items long
type Max140Text string

// Max15NumericText Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Max16Text May be no more than 16 items long
type Max16Text string

// Max2048Text May be no more than 2048 items long
type Max2048Text string

// Max34Text May be no more than 34 items long
type Max34Text string

// Max350Text May be no more than 350 items long
type Max350Text string

// Max35Text May be no more than 35 items long
type Max35Text string

// Max4Text May be no more than 4 items long
type Max4Text string

// Max70Text May be no more than 70 items long
type Max70Text string

type NameAndAddress16 struct {
	Nm  *Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Nm" json:"Nm"`
	Adr *PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Adr" json:"Adr"`
}

// NamePrefix2Code May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification29 struct {
	AnyBIC *AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 AnyBIC,omitempty" json:"AnyBIC,omitempty"`
	LEI    *LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 LEI,omitempty" json:"LEI,omitempty"`
	Othr   []*GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Othr,omitempty" json:"Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    *ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type OtherContact1 struct {
	ChanlTp *Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ChanlTp" json:"ChanlTp"`
	Id      *Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Id,omitempty" json:"Id,omitempty"`
}

type Party38Choice struct {
	OrgId  *OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 OrgId,omitempty" json:"OrgId,omitempty"`
	PrvtId *PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PrvtId,omitempty" json:"PrvtId,omitempty"`
}

type PartyIdentification135 struct {
	Nm        *Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Nm,omitempty" json:"Nm,omitempty"`
	PstlAdr   *PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PstlAdr,omitempty" json:"PstlAdr,omitempty"`
	Id        *Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Id,omitempty" json:"Id,omitempty"`
	CtryOfRes *CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CtryOfRes,omitempty" json:"CtryOfRes,omitempty"`
	CtctDtls  *Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CtctDtls,omitempty" json:"CtctDtls,omitempty"`
}

type PaymentIdentification13 struct {
	InstrId    *Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstrId,omitempty" json:"InstrId,omitempty"`
	EndToEndId *Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 EndToEndId" json:"EndToEndId"`
	TxId       *Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TxId,omitempty" json:"TxId,omitempty"`
	UETR       *UUIDv4Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 UETR,omitempty" json:"UETR,omitempty"`
	ClrSysRef  *Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ClrSysRef,omitempty" json:"ClrSysRef,omitempty"`
}

type PaymentTypeInformation28 struct {
	InstrPrty *Priority2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstrPrty,omitempty" json:"InstrPrty,omitempty"`
	ClrChanl  *ClearingChannel2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ClrChanl,omitempty" json:"ClrChanl,omitempty"`
	SvcLvl    []*ServiceLevel8Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SvcLvl,omitempty" json:"SvcLvl,omitempty"`
	LclInstrm *LocalInstrument2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 LclInstrm,omitempty" json:"LclInstrm,omitempty"`
	CtgyPurp  *CategoryPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CtgyPurp,omitempty" json:"CtgyPurp,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DtAndPlcOfBirth,omitempty" json:"DtAndPlcOfBirth,omitempty"`
	Othr            []*GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Othr,omitempty" json:"Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    *ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

// PhoneNumber Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 AdrTp,omitempty" json:"AdrTp,omitempty"`
	Dept        *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Dept,omitempty" json:"Dept,omitempty"`
	SubDept     *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SubDept,omitempty" json:"SubDept,omitempty"`
	StrtNm      *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 StrtNm,omitempty" json:"StrtNm,omitempty"`
	BldgNb      *Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 BldgNb,omitempty" json:"BldgNb,omitempty"`
	BldgNm      *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 BldgNm,omitempty" json:"BldgNm,omitempty"`
	Flr         *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Flr,omitempty" json:"Flr,omitempty"`
	PstBx       *Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PstBx,omitempty" json:"PstBx,omitempty"`
	Room        *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Room,omitempty" json:"Room,omitempty"`
	PstCd       *Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PstCd,omitempty" json:"PstCd,omitempty"`
	TwnNm       *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TwnNm,omitempty" json:"TwnNm,omitempty"`
	TwnLctnNm   *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TwnLctnNm,omitempty" json:"TwnLctnNm,omitempty"`
	DstrctNm    *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DstrctNm,omitempty" json:"DstrctNm,omitempty"`
	CtrySubDvsn *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CtrySubDvsn,omitempty" json:"CtrySubDvsn,omitempty"`
	Ctry        *CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Ctry,omitempty" json:"Ctry,omitempty"`
	AdrLine     []*Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 AdrLine,omitempty" json:"AdrLine,omitempty"`
}

// PreferredContactMethod1Code May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

// Priority2Code May be one of HIGH, NORM
type Priority2Code string

// Priority3Code May be one of URGT, HIGH, NORM
type Priority3Code string

type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	Id *Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Id" json:"Id"`
}

type ProxyAccountType1Choice struct {
	Cd    *ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type Purpose2Choice struct {
	Cd    *ExternalPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type ReferredDocumentInformation7 struct {
	Tp       *ReferredDocumentType4      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	Nb       *Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Nb,omitempty" json:"Nb,omitempty"`
	RltdDt   *ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RltdDt,omitempty" json:"RltdDt,omitempty"`
	LineDtls []*DocumentLineInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 LineDtls,omitempty" json:"LineDtls,omitempty"`
}

type ReferredDocumentType3Choice struct {
	Cd    *DocumentType6Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type ReferredDocumentType4 struct {
	CdOrPrtry *ReferredDocumentType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdOrPrtry" json:"CdOrPrtry"`
	Issr      *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Issr,omitempty" json:"Issr,omitempty"`
}

type RegulatoryAuthority2 struct {
	Nm   *Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Nm,omitempty" json:"Nm,omitempty"`
	Ctry *CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Ctry,omitempty" json:"Ctry,omitempty"`
}

type RegulatoryReporting3 struct {
	DbtCdtRptgInd *RegulatoryReportingType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DbtCdtRptgInd,omitempty" json:"DbtCdtRptgInd,omitempty"`
	Authrty       *RegulatoryAuthority2             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Authrty,omitempty" json:"Authrty,omitempty"`
	Dtls          []*StructuredRegulatoryReporting3 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Dtls,omitempty" json:"Dtls,omitempty"`
}

// RegulatoryReportingType1Code May be one of CRED, DEBT, BOTH
type RegulatoryReportingType1Code string

type RemittanceAmount2 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DuePyblAmt,omitempty" json:"DuePyblAmt,omitempty"`
	DscntApldAmt      []*DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DscntApldAmt,omitempty" json:"DscntApldAmt,omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdtNoteAmt,omitempty" json:"CdtNoteAmt,omitempty"`
	TaxAmt            []*TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TaxAmt,omitempty" json:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []*DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 AdjstmntAmtAndRsn,omitempty" json:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RmtdAmt,omitempty" json:"RmtdAmt,omitempty"`
}

type RemittanceAmount3 struct {
	DuePyblAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DuePyblAmt,omitempty" json:"DuePyblAmt,omitempty"`
	DscntApldAmt      []*DiscountAmountAndType1          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DscntApldAmt,omitempty" json:"DscntApldAmt,omitempty"`
	CdtNoteAmt        *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdtNoteAmt,omitempty" json:"CdtNoteAmt,omitempty"`
	TaxAmt            []*TaxAmountAndType1               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TaxAmt,omitempty" json:"TaxAmt,omitempty"`
	AdjstmntAmtAndRsn []*DocumentAdjustment1             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 AdjstmntAmtAndRsn,omitempty" json:"AdjstmntAmtAndRsn,omitempty"`
	RmtdAmt           *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RmtdAmt,omitempty" json:"RmtdAmt,omitempty"`
}

type RemittanceInformation16 struct {
	Ustrd []*Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Ustrd,omitempty" json:"Ustrd,omitempty"`
	Strd  []*StructuredRemittanceInformation16 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Strd,omitempty" json:"Strd,omitempty"`
}

type RemittanceLocation7 struct {
	RmtId       *Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RmtId,omitempty" json:"RmtId,omitempty"`
	RmtLctnDtls []*RemittanceLocationData1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RmtLctnDtls,omitempty" json:"RmtLctnDtls,omitempty"`
}

type RemittanceLocationData1 struct {
	Mtd        *RemittanceLocationMethod2Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Mtd" json:"Mtd"`
	ElctrncAdr *Max2048Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ElctrncAdr,omitempty" json:"ElctrncAdr,omitempty"`
	PstlAdr    *NameAndAddress16              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PstlAdr,omitempty" json:"PstlAdr,omitempty"`
}

// RemittanceLocationMethod2Code May be one of FAXI, EDIC, URID, EMAL, POST, SMSM
type RemittanceLocationMethod2Code string

type ServiceLevel8Choice struct {
	Cd    *ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type SettlementDateTimeIndication1 struct {
	DbtDtTm *ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DbtDtTm,omitempty" json:"DbtDtTm,omitempty"`
	CdtDtTm *ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdtDtTm,omitempty" json:"CdtDtTm,omitempty"`
}

type SettlementInstruction7 struct {
	SttlmMtd             *SettlementMethod1Code                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SttlmMtd" json:"SttlmMtd"`
	SttlmAcct            *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SttlmAcct,omitempty" json:"SttlmAcct,omitempty"`
	ClrSys               *ClearingSystemIdentification3Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ClrSys,omitempty" json:"ClrSys,omitempty"`
	InstgRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstgRmbrsmntAgt,omitempty" json:"InstgRmbrsmntAgt,omitempty"`
	InstgRmbrsmntAgtAcct *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstgRmbrsmntAgtAcct,omitempty" json:"InstgRmbrsmntAgtAcct,omitempty"`
	InstdRmbrsmntAgt     *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstdRmbrsmntAgt,omitempty" json:"InstdRmbrsmntAgt,omitempty"`
	InstdRmbrsmntAgtAcct *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 InstdRmbrsmntAgtAcct,omitempty" json:"InstdRmbrsmntAgtAcct,omitempty"`
	ThrdRmbrsmntAgt      *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ThrdRmbrsmntAgt,omitempty" json:"ThrdRmbrsmntAgt,omitempty"`
	ThrdRmbrsmntAgtAcct  *CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 ThrdRmbrsmntAgtAcct,omitempty" json:"thrd-rmbrsmnt-agt-acct,omitempty"`
}

// SettlementMethod1Code May be one of INDA, INGA, COVE, CLRG
type SettlementMethod1Code string

type SettlementTimeRequest2 struct {
	CLSTm  *ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CLSTm,omitempty" json:"CLSTm,omitempty"`
	TillTm *ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TillTm,omitempty" json:"TillTm,omitempty"`
	FrTm   *ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FrTm,omitempty" json:"fr-tm,omitempty"`
	RjctTm *ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RjctTm,omitempty" json:"RjctTm,omitempty"`
}

type StructuredRegulatoryReporting3 struct {
	Tp   *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	Dt   *ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Dt,omitempty" json:"Dt,omitempty"`
	Ctry *CountryCode                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Ctry,omitempty" json:"Ctry,omitempty"`
	Cd   *Max10Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Amt  *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Amt,omitempty" json:"Amt,omitempty"`
	Inf  []*Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Inf,omitempty" json:"Inf,omitempty"`
}

type StructuredRemittanceInformation16 struct {
	RfrdDocInf  []*ReferredDocumentInformation7 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RfrdDocInf,omitempty" json:"RfrdDocInf,omitempty"`
	RfrdDocAmt  *RemittanceAmount2              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RfrdDocAmt,omitempty" json:"RfrdDocAmt,omitempty"`
	CdtrRefInf  *CreditorReferenceInformation2  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CdtrRefInf,omitempty" json:"CdtrRefInf,omitempty"`
	Invcr       *PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Invcr,omitempty" json:"Invcr,omitempty"`
	Invcee      *PartyIdentification135         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Invcee,omitempty" json:"Invcee,omitempty"`
	TaxRmt      *TaxInformation7                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TaxRmt,omitempty" json:"TaxRmt,omitempty"`
	GrnshmtRmt  *Garnishment3                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 GrnshmtRmt,omitempty" json:"GrnshmtRmt,omitempty"`
	AddtlRmtInf []*Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 AddtlRmtInf,omitempty" json:"AddtlRmtInf,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm *Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 PlcAndNm,omitempty" json:"PlcAndNm,omitempty"`
	Envlp    *SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Envlp" json:"Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any" json:"Item"`
}

type TaxAmount2 struct {
	Rate         float64                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Rate,omitempty" json:"Rate,omitempty"`
	TaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TaxblBaseAmt,omitempty" json:"TaxblBaseAmt,omitempty"`
	TtlAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TtlAmt,omitempty" json:"TtlAmt,omitempty"`
	Dtls         []*TaxRecordDetails2               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Dtls,omitempty" json:"Dtls,omitempty"`
}

type TaxAmountAndType1 struct {
	Tp  *TaxAmountType1Choice              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Amt" json:"Amt"`
}

type TaxAmountType1Choice struct {
	Cd    *ExternalTaxAmountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cd,omitempty" json:"Cd,omitempty"`
	Prtry *Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prtry,omitempty" json:"Prtry,omitempty"`
}

type TaxAuthorisation1 struct {
	Titl *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Titl,omitempty" json:"Titl,omitempty"`
	Nm   *Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Nm,omitempty" json:"Nm,omitempty"`
}

type TaxInformation7 struct {
	Cdtr            *TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cdtr,omitempty" json:"Cdtr,omitempty"`
	Dbtr            *TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Dbtr,omitempty" json:"Dbtr,omitempty"`
	UltmtDbtr       *TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 UltmtDbtr,omitempty" json:"UltmtDbtr,omitempty"`
	AdmstnZone      *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 AdmstnZone,omitempty" json:"AdmstnZone,omitempty"`
	RefNb           *Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RefNb,omitempty" json:"RefNb,omitempty"`
	Mtd             *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Mtd,omitempty" json:"Mtd,omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TtlTaxblBaseAmt,omitempty" json:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TtlTaxAmt,omitempty" json:"TtlTaxAmt,omitempty"`
	Dt              *ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Dt,omitempty" json:"Dt,omitempty"`
	SeqNb           float64                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SeqNb,omitempty" json:"SeqNb,omitempty"`
	Rcrd            []*TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Rcrd,omitempty" json:"Rcrd,omitempty"`
}

type TaxInformation8 struct {
	Cdtr            *TaxParty1                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Cdtr,omitempty" json:"Cdtr,omitempty"`
	Dbtr            *TaxParty2                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Dbtr,omitempty" json:"Dbtr,omitempty"`
	AdmstnZone      *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 AdmstnZone,omitempty" json:"AdmstnZone,omitempty"`
	RefNb           *Max140Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RefNb,omitempty" json:"RefNb,omitempty"`
	Mtd             *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Mtd,omitempty" json:"Mtd,omitempty"`
	TtlTaxblBaseAmt *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TtlTaxblBaseAmt,omitempty" json:"TtlTaxblBaseAmt,omitempty"`
	TtlTaxAmt       *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TtlTaxAmt,omitempty" json:"TtlTaxAmt,omitempty"`
	Dt              *ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Dt,omitempty" json:"Dt,omitempty"`
	SeqNb           float64                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 SeqNb,omitempty" json:"SeqNb,omitempty"`
	Rcrd            []*TaxRecord2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Rcrd,omitempty" json:"Rcrd,omitempty"`
}

type TaxParty1 struct {
	TaxId  *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TaxId,omitempty" json:"TaxId,omitempty"`
	RegnId *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RegnId,omitempty" json:"RegnId,omitempty"`
	TaxTp  *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TaxTp,omitempty" json:"TaxTp,omitempty"`
}

type TaxParty2 struct {
	TaxId   *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TaxId,omitempty" json:"TaxId,omitempty"`
	RegnId  *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 RegnId,omitempty" json:"RegnId,omitempty"`
	TaxTp   *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TaxTp,omitempty" json:"TaxTp,omitempty"`
	Authstn *TaxAuthorisation1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Authstn,omitempty" json:"Authstn,omitempty"`
}

type TaxPeriod2 struct {
	Yr     *ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Yr,omitempty" json:"Yr,omitempty"`
	Tp     *TaxRecordPeriod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	FrToDt *DatePeriod2          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FrToDt,omitempty" json:"FrToDt,omitempty"`
}

type TaxRecord2 struct {
	Tp       *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Tp,omitempty" json:"Tp,omitempty"`
	Ctgy     *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Ctgy,omitempty" json:"Ctgy,omitempty"`
	CtgyDtls *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CtgyDtls,omitempty" json:"CtgyDtls,omitempty"`
	DbtrSts  *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 DbtrSts,omitempty" json:"DbtrSts,omitempty"`
	CertId   *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 CertId,omitempty" json:"CertId,omitempty"`
	FrmsCd   *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 FrmsCd,omitempty" json:"FrmsCd,omitempty"`
	Prd      *TaxPeriod2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prd,omitempty" json:"Prd,omitempty"`
	TaxAmt   *TaxAmount2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 TaxAmt,omitempty" json:"TaxAmt,omitempty"`
	AddtlInf *Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 AddtlInf,omitempty" json:"AddtlInf,omitempty"`
}

type TaxRecordDetails2 struct {
	Prd *TaxPeriod2                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Prd,omitempty" json:"Prd,omitempty"`
	Amt *ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.09 Amt" json:"Amt"`
}

// TaxRecordPeriod1Code May be one of MM01, MM02, MM03, MM04, MM05, MM06, MM07, MM08, MM09, MM10, MM11, MM12, QTR1, QTR2, QTR3, QTR4, HLF1, HLF2
type TaxRecordPeriod1Code string

// UUIDv4Identifier Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write(b)
	enc.Close()
	return buf.Bytes(), nil
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
