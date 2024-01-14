package main

import (
	"fmt"

	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pedrobealves/pedrobealves/internal/models"
	"github.com/pedrobealves/pedrobealves/internal/readjson"
	"github.com/pedrobealves/pedrobealves/internal/replacers"
)

type Data struct {
	Name string
	// Adicione mais campos conforme necessário
}

func main() {

	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	rootDir := filepath.Join(dir, "..", "..")

	// Create an instance of your struct
	profile := models.Profile{}

	// Call the Create function with the struct instance
	result, err := readjson.Create(filepath.Join(rootDir, "data.json"), &profile)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// The result will be the same struct with data populated
	resultModule, ok := result.(*models.Profile)
	if !ok {
		fmt.Println("Error: Result is not of the expected type")
		return
	}

	// Print the module
	// Example of accessing the data
	//fmt.Println("User:", resultModule.User)
	//fmt.Println("Header Description:", resultModule.Header.Description)
	// ... continue accessing the data as needed
	//fmt.Println("About:", replacers.ReplaceAbout(resultModule))
	//fmt.Println("Stats:", replacers.ReplaceStats(resultModule))
	// Ler o conteúdo do arquivo README_TEMPLATE.md
	// Abrir o arquivo README_TEMPLATE.md

	file, err := os.Open(filepath.Join(rootDir, "README_TEMPLATE.md"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Ler todo o conteúdo do arquivo
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Criar um objeto Data com os valores de substituição
	data := Data{
		Name: replacers.ReplaceStats(resultModule),
		// Preencha outros campos conforme necessário
	}

	// Criar um novo template com a sintaxe esperada
	tmpl, err := template.New("readme").Parse(string(content))
	if err != nil {
		panic(err)
	}

	// Criar um buffer para armazenar a saída do template
	var updatedMd strings.Builder

	// Aplicar o template aos dados e armazenar o resultado no buffer
	err = tmpl.Execute(&updatedMd, data)
	if err != nil {
		panic(err)
	}

	// Construa o caminho para o arquivo que você deseja escrever na raiz do projeto
	filePath := filepath.Join(rootDir, "README.md")

	// Escrever o conteúdo atualizado no arquivo README.md
	err = os.WriteFile(filePath, []byte(updatedMd.String()), 0644)
	if err != nil {
		panic(err)
	}

	// Concluir
	println("README update complete.")
}
