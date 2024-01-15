package readjson

import (
	"encoding/json"
	"fmt"
)

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
func Read(file []byte, data interface{}) (interface{}, error) {
	return decodeJSON(file, data)
}
