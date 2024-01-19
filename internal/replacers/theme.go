package replacers

import (
	"fmt"
	"reflect"

	"github.com/pedrobealves/pedrobealves/internal/models"
	"github.com/pedrobealves/pedrobealves/internal/readjson"
	"github.com/pedrobealves/pedrobealves/internal/utils"
)

// Colors armazena as cores disponíveis.
var Colors = map[string]string{}

func init() {

	fileData, err := utils.OpenReadFile("theme.json")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Call the Create function with the struct instance
	result, err := readjson.Read(fileData, &models.Theme{})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// The result will be the same struct with data populated
	resultModule, ok := result.(*models.Theme)

	if !ok {
		fmt.Println("Error: Result is not of the expected type")
		return
	}

	Colors["primary"] = resultModule.Colors.Primary
	Colors["secondary"] = resultModule.Colors.Secondary
	Colors["lightSecondary"] = resultModule.Colors.LightSecondary
	Colors["white"] = resultModule.Colors.White
	Colors["black"] = resultModule.Colors.Black

}

// Função para mapear as cores de forma genérica
func mapColors(data interface{}, keys ...string) {
	v := reflect.ValueOf(data).Elem()

	for _, key := range keys {
		field := v.FieldByName(key)
		if field.Kind() == reflect.String {
			colorName := field.Interface().(string)
			if newColor, ok := Colors[colorName]; ok {
				field.SetString(newColor)
			}
		}
	}
}
