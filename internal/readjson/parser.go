package readjson

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// openFile abre o arquivo com o nome fornecido.
func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return nil, err
	}
	return file, nil
}

// readFileContent lê o conteúdo do arquivo fornecido.
func readFileContent(file *os.File) ([]byte, error) {
	fileData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return nil, err
	}
	return fileData, nil
}

// decodeJSON decodifica o conteúdo JSON no objeto fornecido.
func decodeJSON(fileData []byte, data interface{}) (interface{}, error) {
	err := json.Unmarshal(fileData, data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, err
	}
	fmt.Println("Structure of the defined struct.")
	return data, nil
}

// Create abre um arquivo, lê seu conteúdo e decodifica o conteúdo JSON na estrutura fornecida.
func Create(fileName string, data interface{}) (interface{}, error) {
	file, err := openFile(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileData, err := readFileContent(file)
	if err != nil {
		return nil, err
	}

	return decodeJSON(fileData, data)
}
