package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
)

// getFieldName converte o nome do campo para snake_case.
func getFieldName(fieldName string, isSnakeCase bool) string {
	if isSnakeCase {
		return strcase.ToSnake(fieldName)
	}
	return strcase.ToLowerCamel(fieldName)
}

// normalizeFieldValue normaliza o valor do campo, especialmente se o nome do campo contiver "color".
func normalizeFieldValue(fieldName, fieldValue string) string {
	if strings.Contains(strings.ToLower(fieldName), "color") {
		return NormalizeColor(fieldValue)
	}
	return fieldValue
}

// queryFromObject realiza a iteração sobre os campos da estrutura e formata os dados para uma string de consulta.
func queryFromObject(val reflect.Value, isSnakeCase bool) string {
	var result []string

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := getFieldName(typ.Field(i).Name, isSnakeCase)
		fieldValue := fmt.Sprintf("%v", field.Interface())
		fieldValue = normalizeFieldValue(fieldName, fieldValue)

		result = append(result, fmt.Sprintf("%s=%s", fieldName, fieldValue))
	}

	return strings.Join(result, "&")
}

// QueryFromObject gera e retorna uma string de consulta a partir dos campos de uma estrutura.
func QueryFromObject(properties interface{}, isSnakeCase bool) string {
	val := reflect.ValueOf(properties)
	return queryFromObject(val, isSnakeCase)
}

func NormalizeColor(color string) string {
	return strings.TrimPrefix(color, "#")
}
