package main

import (
	csv "github.com/sagar0419/csvJsonEncoding/csv"
)

func main() {
	csv.ReadCsv("timezone.csv")
	csv.WriteCsv("timezone.csv")
}
