package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCsv(filePath string) {

	// opening file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	// CSV Reader
	readeCsv := csv.NewReader(file)

	// Read CSV records
	records, err := readeCsv.ReadAll()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Record printing
	for _, record := range records {
		fmt.Println(record)
	}
}
