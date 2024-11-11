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

	// open existing csv to write, if csv is not there then create it
	file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// CSV writer
	writer := csv.NewWriter(file)

	// as I have multiple row, so this loop will iterate over multiple row and add it to CSV
	for _, dataTimeZone := range newTimeZone {
		if err := writer.Write(dataTimeZone); err != nil {
			log.Fatalf("failed to write to csv file: %s", err)
			return
		}
	}

	// Flushing writer so that all data is written in CSV
	writer.Flush()

	// CHecking if any error happened during flushing
	if err := writer.Error(); err != nil {
		log.Fatal(err)
		return
	}

	// Printing the success msg
	log.Println("Data is scucessfully written")
}
