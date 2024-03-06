package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileNameByScenario_ValidScenario(t *testing.T) {
	valid_scenario_input := "aoc_01"
	var fileNameRegistry FileNameRegistry = FileNameRegistry{}

	result, _ := fileNameRegistry.getFileNamePerCase(valid_scenario_input)
	expected := "puzzle_data/aoc_01.csv"
	assert.Equal(t, result, expected)

}

func TestGetFileNameByScenario_InValidScenario(t *testing.T) {
	valid_scenario_input := "invalid"
	var fileNameRegistry FileNameRegistry = FileNameRegistry{}

	result, err := fileNameRegistry.getFileNamePerCase(valid_scenario_input)
	expected := "puzzle_data/aoc_01.csv"
	assert.NotEqual(t, result, expected)
	assert.Equal(t, err, errors.New("This is not an implemented scenario"))

}
