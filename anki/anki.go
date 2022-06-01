package anki

import (
	"encoding/csv"
	"go-shv-to-anki-convert/helper"
	"go-shv-to-anki-convert/shv"
	"log"
	"os"
)

func ConvertToAnki(questions []shv.Question, examCategory string, imageLocation string) ([]AnkiCard, error) {
	var ankiCards = []AnkiCard{}

	for _, question := range questions {
		front := question.Question

		if question.ImageID != "" {
			imageName, err := helper.HandleImage(question.ImageID, imageLocation)
			if err != nil {
				return nil, err
			}

			front += "<br><img src=\"" + imageName + "\">"
		}

		front += "<br><ol>"
		front += "<li>"
		front += question.Answer1
		front += "</li><li>"
		front += question.Answer2
		front += "</li><li>"
		front += question.Answer3
		front += "</li><li>"
		front += question.Answer4
		front += "</li></ol>"

		back := ""
		switch question.Answer {
		case "1":
			back = "1. " + question.Answer1
		case "2":
			back = "2. " + question.Answer2
		case "3":
			back = "3. " + question.Answer3
		case "4":
			back = "4. " + question.Answer4
		}

		ankiCard := AnkiCard{
			Front: front,
			Back:  back,
			Tags:  examCategory,
		}

		ankiCards = append(ankiCards, ankiCard)
	}

	return ankiCards, nil
}

func SaveAnkiAsCSV(fileName string, ankiCards []AnkiCard) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush() // Using Write

	// Using WriteAll
	var data [][]string
	for _, record := range ankiCards {
		row := []string{record.Front, record.Back, record.Tags}
		data = append(data, row)
	}
	w.WriteAll(data)
}

type AnkiCard struct {
	Front string
	Back  string
	Tags  string
}
