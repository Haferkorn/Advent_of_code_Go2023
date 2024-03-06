package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

type FileNameRegistry struct{}

func (FileNameRegistry FileNameRegistry) getFileNamePerCase(scenario string) (string, error) {
	switch scenario {
	case "aoc_01":
		return "puzzle_data/aoc_01.csv", nil
	default:
		return "", errors.New("This is not an implemented scenario")
	}
}

func (FileNameRegistry FileNameRegistry) getCSVOpenerCase(scenario string) (FileOpener, error) {
	switch scenario {
	case "aoc_01":
		return realFileOpener{}, nil
	default:
		return nil, errors.New("This is not an implemented scenario")
	}
}

func _flatten(slice []string) string {
	return fmt.Sprintf("%v", slice[0])
}

// FileOpener interface abstracts the os.Open function for testing
type FileOpener interface {
	Open(name string) (*os.File, error)
}

// realFileOpener is the concrete implementation of the FileOpener interface
type realFileOpener struct{}

func (r realFileOpener) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func getScenarioNameFromUser() string {
	fmt.Print("Enter the name of the scenario you are working on: ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic("Error when entering input")
	}

	return input

}

func getFileNameByScenario(input string, fileRegistry FileNameRegistry) string {
	fileName, err := fileRegistry.getFileNamePerCase(input)

	if err != nil {
		getScenarioNameFromUser()
	}

	return fileName
}

func getFileOpenerByScenario(input string, fileRegistry FileNameRegistry) FileOpener {
	fileOpener, err := fileRegistry.getCSVOpenerCase(input)

	if err != nil {
		getScenarioNameFromUser()
	}

	return fileOpener
}

func _format_file_data(dataFromFile [][]string) []string {
	var result []string
	for _, listed_item := range dataFromFile {
		result = append(result, _flatten(listed_item))
	}
	return result
}

func readCSV(fileName string, fileOpener FileOpener) ([]string, error) {

	// Open the CSV file
	file, err := fileOpener.Open(fileName)

	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all lines from the CSV
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return _format_file_data(lines), nil
}
