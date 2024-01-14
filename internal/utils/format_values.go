package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
)

func NormalizeColor(color string) string {
	return strings.TrimPrefix(color, "#")
}

func QueryFromObject(properties interface{}) string {
	var result []string

	val := reflect.ValueOf(properties)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := strcase.ToSnake(typ.Field(i).Name)
		fieldValue := fmt.Sprintf("%v", field.Interface())

		// Se o nome do campo contém "color", normalize o valor
		if strings.Contains(strings.ToLower(fieldName), "color") {
			fieldValue = NormalizeColor(fieldValue)
		}

		result = append(result, fmt.Sprintf("%s=%s", fieldName, fieldValue))
	}

	return strings.Join(result, "&")
}
