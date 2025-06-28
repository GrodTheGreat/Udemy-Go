package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	defer checkClose(file, &err)

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return nil, errors.New("failed to read line in file")
	}

	return lines, err
}

func (fm *FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}
	defer checkClose(file, &err)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	return err
}

func checkClose(f *os.File, err *error) {
	if closeErr := f.Close(); closeErr != nil {
		if *err != nil {
			fmt.Println(*err)
		} else {
			*err = closeErr
		}
	}
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
