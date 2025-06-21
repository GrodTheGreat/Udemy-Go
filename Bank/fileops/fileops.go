package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const defaultValue float64 = 1_000

func WriteFloatToFile(fileName string, value float64) {
	text := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(text), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}

	text := string(data)
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return defaultValue, errors.New("failed to parse stored value")
	}

	return value, nil
}
