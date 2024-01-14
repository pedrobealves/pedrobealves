package main

import (
	"fmt"

	"github.com/pedrobealves/pedrobealves/internal/models"
	"github.com/pedrobealves/pedrobealves/internal/readjson"
	"github.com/pedrobealves/pedrobealves/internal/replacers"
)

func main() {
	// Create an instance of your struct
	module := models.Profile{}

	// Call the Create function with the struct instance
	result, err := readjson.Create(&module)
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
	fmt.Println("About:", replacers.ReplaceAbout(resultModule))
	fmt.Println("Stats:", replacers.ReplaceStats(resultModule))
}
