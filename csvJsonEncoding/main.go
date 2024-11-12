package main

import (
	json "github.com/sagar0419/csvJsonEncoding/json"
)

func main() {
	// csv.ReadCsv("timezone.csv")
	// csv.WriteCsv("timezone.csv")
	json.EncodeJson()
	json.DecodeJson()
	json.ReadJson("sample3.json")
	json.WriteJson("sample3.json")
}
