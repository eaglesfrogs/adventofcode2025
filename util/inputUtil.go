package util

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const DefaultDirectory = "./.input"
const InputUrlTemplate = "https://adventofcode.com/2025/day/%d/input"

func GetInputPath(day int, session string) (string, error) {
	_, err := os.Stat(DefaultDirectory)
	if err != nil {
		log.Printf("Input directory %s not found, creating...\n", DefaultDirectory)
		err = os.Mkdir(DefaultDirectory, 0755)

		if err != nil {
			return "", err
		}
	}

	inputPath := DefaultDirectory + "/day" + strconv.Itoa(day) + ".txt"
	_, err = os.Stat(inputPath)

	if os.IsNotExist(err) {
		log.Printf("Input file for day %d not found, downloading...\n", day)
		err = DownloadInputFile(day, session, inputPath)
		if err != nil {
			return "", err
		}
	} else {
		log.Printf("Input file for day %d found locally.\n", day)
	}

	return inputPath, nil
}

func DownloadInputFile(day int, session string, inputPath string) error {
	url := fmt.Sprintf(InputUrlTemplate, day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download input file: status code %d", resp.StatusCode)
	}

	outFile, err := os.Create(inputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = outFile.ReadFrom(resp.Body)
	if err != nil {
		return err
	}

	return nil
}
