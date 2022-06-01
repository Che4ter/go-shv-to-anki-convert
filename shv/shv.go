package shv

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const ImageURL = "https://elearning.shv-fsvl.ch/pictures/"

func ParseQuestions(fileName string) ([]Question, error) {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var root ShvExamSoapResponse
	err = xml.Unmarshal(byteValue, &root)
	if err != nil {
		return nil, err
	}

	return root.Body.GetQuestionsResponse.GetQuestionsResult.Result.Questions, nil
}

func DownloadExamImage(filepath string, imageName string) error {

	// Get the data
	resp, err := http.Get(ImageURL + imageName)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath + imageName)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

type ShvExamSoapResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Text                 string `xml:",chardata"`
		GetQuestionsResponse struct {
			GetQuestionsResult struct {
				Result struct {
					Questions []Question `xml:"__Question"`
				} `xml:"Result"`
			} `xml:"GetQuestionsResult"`
		} `xml:"GetQuestionsResponse"`
	} `xml:"Body"`
}

type Question struct {
	Text     string `xml:",chardata"`
	ID       string `xml:"ID"`
	ImageID  string `xml:"ImageID"`
	Answer   string `xml:"Answer"`
	Question string `xml:"Question"`
	Answer1  string `xml:"Answer1"`
	Answer2  string `xml:"Answer2"`
	Answer3  string `xml:"Answer3"`
	Answer4  string `xml:"Answer4"`
}
