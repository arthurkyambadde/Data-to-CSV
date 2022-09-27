package createcsv

import (
	"encoding/csv"
	"log"
	"os"
)

func CreateCsv() {
	FarmerChallenges := [][]string{
		{"1", "Climate change"},
		{"2", "The ongoing trade war between the United States and China."},
		{"3", "Rapidly depleting reserves of freshwater around the world."},
		{"4", "The looming food crisis."},
		{"5", "Economic insecurity in the United States."},
		{"6", "Ongoing closures of food processing facilities and local businesses due to the COVID-19 pandemic."},
		{"7", "Depletion of natural resources due to widespread industrial agricultural practices."},
		{"8", "High rates of food waste, which threaten to intensify food insecurity around the globe."},
		{"9", "Disruptions in trade networks and fluctuations in global demand for agricultural products."},
		{"10", "Economic strife and crippling debt for individual farmers."},
	}

	csvFile, err := os.Create("Farmer_challenges.csv")

	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}

	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Comma = ':'

	for _, challenge := range FarmerChallenges {
		_ = csvwriter.Write(challenge)
	}

	csvwriter.Flush()
	csvFile.Close()

}
