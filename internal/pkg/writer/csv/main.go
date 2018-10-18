package csv

import (
	"encoding/csv"
	"log"
	"os"
)

// Write CSV in the filenameOut, overwriting its contents
func Write(filenameOut string, records [][]string) {
	file, _ := os.Create(filenameOut)
	defer file.Close()

	w := csv.NewWriter(file)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("Error writing record to csv: ", err)
		}
	}

	defer w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
