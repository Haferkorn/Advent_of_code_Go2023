package sum_digits

import (
	"unicode"
)

func _getNumberPerEntry(entry string) int {
	digitsSlice := make([]int, 0, len(entry)) // Preallocate space for efficiency
	for _, char := range entry {
		if unicode.IsDigit(char) {
			digitsSlice = append(digitsSlice, int(char-'0'))
		}
	}

	// Check if there are digits before accessing the slice
	if len(digitsSlice) > 0 {
		return digitsSlice[0]*10 + digitsSlice[len(digitsSlice)-1]
	}
	return 0 // No digits found
}

func SumDigitsFromEntry(formattedDataList []string) int {
	total_sum := 0

	for _, entry := range formattedDataList {
		total_sum += _getNumberPerEntry(entry)
	}

	return total_sum
}
