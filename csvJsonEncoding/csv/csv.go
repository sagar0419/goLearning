package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Csv(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	readeCsv := csv.NewReader(file)

	records, err := readeCsv.ReadAll()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, record := range records {
		fmt.Println(record)
	}
}
