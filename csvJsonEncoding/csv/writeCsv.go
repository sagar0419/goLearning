package csv

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteCsv(FilePath string) {

	// New Data
	newTimeZone := [][]string{
		{"Asia/India", "Kolkata", "Asia"},
		{"Asia/Nepal", "Kathmandu", "Asia"},
		{"Asia/Srilanka", "colombo", "Asia"},
	}

	file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	for _, dataTimeZone := range newTimeZone {
		if err := writer.Write(dataTimeZone); err != nil {
			log.Fatalf("failed to write to csv file: %s", err)
			return
		}
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Data is scucessfully written")
}
