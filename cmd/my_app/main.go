package main

import (
	"GoProject/csv_reader"
	"GoProject/sum_digits"
	"fmt"
)

func main() {
	var fileNameRegistry csv_reader.FileNameRegistry = csv_reader.FileNameRegistry{}
	var input string = csv_reader.GetScenarioNameFromUser()

	var fileName string = csv_reader.GetFileNameByScenario(input, fileNameRegistry)
	var fileOpener csv_reader.FileOpener = csv_reader.GetFileOpenerByScenario(input, fileNameRegistry)

	// Read CSV file
	data, err := csv_reader.ReadCSV(fileName, fileOpener)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	sum := sum_digits.SumDigitsFromEntry(data)
	fmt.Println(sum)

}
