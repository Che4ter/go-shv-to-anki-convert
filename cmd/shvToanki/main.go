package main

import (
	"flag"
	"fmt"
	"go-shv-to-anki-convert/anki"
	"go-shv-to-anki-convert/shv"
	"log"
	"strconv"
)

var Version = "dev build"

func main() {
	fmt.Println("Welcome to the Anki SHV Exam Question converter - Version", Version)

	var examPath string
	var examCategory string
	var imagePath string

	flag.StringVar(&examPath, "xml", "eLearning.asmx", "path to exam xml")
	flag.StringVar(&examCategory, "category", "", "exam category")
	flag.StringVar(&imagePath, "imgPath", "images/", "path to store exam images")
	flag.Parse()

	questions, err := shv.ParseQuestions(examPath)
	if err != nil {
		log.Fatalln(questions)
	}
	ankiCards, err := anki.ConvertToAnki(questions, examCategory, imagePath)
	if err != nil {
		log.Fatalln(questions)
	}
	anki.SaveAnkiAsCSV(examCategory+".csv", ankiCards)

	fmt.Println(strconv.Itoa(len(questions)) + " Questions have been converted to " + strconv.Itoa(len(ankiCards)) + " Anki Cards, stored in the file " + examCategory + ".csv")
	fmt.Println("Copy the images from the folder " + imagePath + " to \"~/.local/share/Anki2/User 1/collection.media\" before importing the questions into anki")

}
