package utils

import (
	"strings"
	"text/template"
)

// createTemplate cria um novo template com a sintaxe esperada.
func CreateTemplate(templateName, content string) (*template.Template, error) {
	tmpl, err := template.New(templateName).Parse(content)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

// applyTemplate aplica o template aos dados e retorna a string resultante.
func ApplyTemplate(tmpl *template.Template, data interface{}) (strings.Builder, error) {
	var result strings.Builder
	err := tmpl.Execute(&result, data)
	if err != nil {
		return strings.Builder{}, err
	}
	return result, nil
}
