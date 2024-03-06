package main

import "fmt"

func main() {
	var fileNameRegistry FileNameRegistry = FileNameRegistry{}
	var input string = getScenarioNameFromUser()

	var fileName string = getFileNameByScenario(input, fileNameRegistry)
	var fileOpener FileOpener = getFileOpenerByScenario(input, fileNameRegistry)

	// Read CSV file
	data, err := readCSV(fileName, fileOpener)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	sum := sumDigitsFromEntry(data)
	fmt.Println(sum)

}
