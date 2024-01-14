package utils

import (
	"bytes"
	"html/template"
)

// ElementData é a estrutura para armazenar dados do elemento HTML.
type ElementData struct {
	Tag        string
	Properties map[string]string
	Children   string
}

// generateElementString gera a string do elemento HTML com base nos parâmetros fornecidos.
func generateElementString(data ElementData) string {
	tmpl, err := template.New("element").Parse(`<{{.Tag}} {{range $key, $value := .Properties}} {{$key}}="{{$value}}"{{end}}{{if .Children}}>{{.Children}}</{{.Tag}}>{{else}}/>{{end}}`)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		panic(err)
	}

	return tpl.String()
}

// GenerateElement gera e retorna a string do elemento HTML.
func GenerateElement(tag string, properties map[string]string, children string) string {
	data := ElementData{
		Tag:        tag,
		Properties: properties,
		Children:   children,
	}

	return generateElementString(data)
}
