package createcsv

import (
	"encoding/csv"
	"log"
	"os"
)

func CreateCsv() {
	Employees := [][]string{
		{"Name", "City", "Skills"},
		{"Smith", "Newyork", "Java"},
		{"William", "PAris", "Golang"},
		{"Rose", "London", "PHP"},
	}

	csvFile, err := os.Create("employee.csv")

	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}

	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	// csvwriter.Comma = ";"

	for _, employee := range Employees {
		_ = csvwriter.Write(employee)
	}

	csvwriter.Flush()
	csvFile.Close()

}
