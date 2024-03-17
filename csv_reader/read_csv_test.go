package csv_reader

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func return_valid_test_data() [][]string {
	var valid_test_data [][]string = [][]string{{"gfdgi"}}
	return valid_test_data
}

// MockFileOpener is a mock implementation of the FileOpener interface
type MockFileOpener struct{}

func (m MockFileOpener) Open(name string) (*os.File, error) {
	return nil, nil
}

func TestCSVReader_valid_data(t *testing.T) {
	test_data := return_valid_test_data()
	result := _format_file_data(test_data)
	expected := []string{"gfdgi"}
	if result[0] != expected[0] {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
	assert.Equal(t, expected, result, "The result should be equal to the expected value")
}
