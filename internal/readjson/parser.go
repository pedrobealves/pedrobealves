package readjson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Create(fileName string, data interface{}) (interface{}, error) {
	// Open the file
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening the file:", err)
		return nil, err
	}
	defer file.Close()

	// Read the file content
	fileData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return nil, err
	}

	// Decode the file content into the provided structure
	err = json.Unmarshal(fileData, data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}

	fmt.Println("Structure of the defined struct.")

	return data, nil
}
