package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func OpenReadFile(fileName string) ([]byte, error) {
	file, err := OpenFile(fileName)

	if err != nil {
		fmt.Println("Error opening the file:", err)
		return nil, err
	}

	fileData, err := ReadFileContent(file)

	if err != nil {
		fmt.Println("Error reading the file:", err)
		return nil, err
	}

	return fileData, nil
}

// openFile abre o arquivo com o nome fornecido.
func OpenFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening the file:", err)
		return nil, err
	}

	return file, nil
}

// readFileContent lê o conteúdo do arquivo fornecido.
func ReadFileContent(file *os.File) ([]byte, error) {
	fileData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return nil, err
	}
	return fileData, nil
}

// writeToFile escreve o conteúdo em um arquivo no caminho fornecido.
func WriteFile(filePath, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func JoinFile(fileName string) string {
	return filepath.Join(fileName)
}
