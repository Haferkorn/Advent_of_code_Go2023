package main

import (
	"GoProject/csv_reader"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileNameByScenario_ValidScenario(t *testing.T) {
	valid_scenario_input := "aoc_01"
	var fileNameRegistry csv_reader.FileNameRegistry = csv_reader.FileNameRegistry{}

	result, _ := fileNameRegistry.GetFileNamePerCase(valid_scenario_input)
	expected := "puzzle_data/aoc_01.csv"
	assert.Equal(t, result, expected)

}

func TestGetFileNameByScenario_InValidScenario(t *testing.T) {
	valid_scenario_input := "invalid"
	var fileNameRegistry csv_reader.FileNameRegistry = csv_reader.FileNameRegistry{}

	result, err := fileNameRegistry.GetFileNamePerCase(valid_scenario_input)
	expected := "puzzle_data/aoc_01.csv"
	assert.NotEqual(t, result, expected)
	assert.Equal(t, err, errors.New("This is not an implemented scenario"))

}
