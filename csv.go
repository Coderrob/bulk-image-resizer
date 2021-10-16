package main

import (
	"encoding/csv"
	"os"
)

func CreateCSV(outputDir string, images []*Image) {
	file, err := os.Create(outputDir + "results.csv")
	if err != nil {
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, image := range images {
		var result []string
		result = append(result, image.From)
		result = append(result, image.To)
		writer.Write(result)
	}
}
