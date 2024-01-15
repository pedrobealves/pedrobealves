package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type File struct {
	rootDir string
}

func NewFile() *File {

	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	return &File{rootDir: filepath.Join(dir, "..", "..")}
}

func (f *File) OpenReadFile(fileName string) ([]byte, error) {
	file, err := f.OpenFile(fileName)

	if err != nil {
		fmt.Println("Error opening the file:", err)
		return nil, err
	}

	fileData, err := f.ReadFileContent(file)

	if err != nil {
		fmt.Println("Error reading the file:", err)
		return nil, err
	}

	return fileData, nil
}

// openFile abre o arquivo com o nome fornecido.
func (f *File) OpenFile(fileName string) (*os.File, error) {
	file, err := os.Open(filepath.Join(f.rootDir, fileName))

	if err != nil {
		fmt.Println("Error opening the file:", err)
		return nil, err
	}

	return file, nil
}

// readFileContent lê o conteúdo do arquivo fornecido.
func (*File) ReadFileContent(file *os.File) ([]byte, error) {
	fileData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return nil, err
	}
	return fileData, nil
}

// writeToFile escreve o conteúdo em um arquivo no caminho fornecido.
func (f *File) WriteFile(filePath, content string) error {
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) JoinFile(fileName string) string {
	return filepath.Join(f.rootDir, fileName)
}
