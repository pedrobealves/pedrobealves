package main

import (
	"fmt"

	"github.com/pedrobealves/pedrobealves/internal/models"
	"github.com/pedrobealves/pedrobealves/internal/readjson"
	"github.com/pedrobealves/pedrobealves/internal/replacers"
	"github.com/pedrobealves/pedrobealves/internal/utils"
)

func main() {

	file := utils.NewFile()

	fileData, err := file.OpenReadFile("data.json")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Call the Create function with the struct instance
	result, err := readjson.Read(fileData, &models.Profile{})

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

	fileTemplate, err := file.OpenReadFile("README_TEMPLATE.md")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Criar um objeto Data com os valores de substituição
	data := models.Data{
		AboutMe:     replacers.ReplaceAbout(resultModule),
		GithubStats: replacers.ReplaceStats(resultModule),
	}

	// Criar um novo template com a sintaxe esperada
	tmpl, err := utils.CreateTemplate("readme", string(fileTemplate))
	if err != nil {
		panic(err)
	}

	// Criar um buffer para armazenar a saída do template
	updatedMd, err := utils.ApplyTemplate(tmpl, data)

	if err != nil {
		panic(err)
	}

	// Construa o caminho para o arquivo que você deseja escrever na raiz do projeto
	filePath := file.JoinFile("README.md")

	// Escrever o conteúdo atualizado no arquivo README.md
	err = file.WriteFile(filePath, updatedMd.String())
	if err != nil {
		panic(err)
	}

	// Concluir
	println("README update complete.")
}
