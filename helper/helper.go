package helper

import (
	"go-shv-to-anki-convert/shv"
	"os"
)

func HandleImage(imageID string, imageLocation string) (string, error) {

	if _, err := os.Stat(imageLocation); os.IsNotExist(err) {
		err := os.Mkdir(imageLocation, 0755)
		if err != nil {
			return "", err
		}
	}

	imageName := imageID + ".jpg"

	if _, err := os.Stat(imageLocation + imageName); os.IsNotExist(err) {
		err := shv.DownloadExamImage(imageLocation, imageName)
		if err != nil {
			return "", err
		}
	}

	return imageName, nil
}
