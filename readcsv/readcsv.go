package readcsv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type challengefaced struct {
	number             string
	challengeStatement string
}

func ReadCsv() {
	csvfile, err := os.Open("Farmer_challenges.csv")

	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	defer csvfile.Close()

	csvReader := csv.NewReader(csvfile)
	csvReader.Comment = '#'
	csvReader.Comma = ':'

	for {
		challenges, err := csvReader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			if parseError, ok := err.(*csv.ParseError); ok {
				fmt.Println("Bad Column: ", parseError.Column)
				fmt.Println("Bad Line: ", parseError.Line)
				fmt.Println("Error reported: ", parseError.Err)
				if parseError == csv.ErrFieldCount {
					continue
				}
			}
		}

		var workersChallenges []challengefaced
		for _, challenge := range challenges {

			data := challengefaced{

				challengeStatement: challenge,
			}

			workersChallenges = append(workersChallenges, data)
		}
		fmt.Println(workersChallenges)

	}
}
