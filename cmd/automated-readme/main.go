package main

import (
	"fmt"

	"github.com/pedrobealves/pedrobealves/internal/readjson"
)

type Styles struct {
	Align       string `json:"align"`
	Style       string `json:"style"`
	Color       string `json:"color"`
	LogoColor   string `json:"logoColor"`
	HighlightColor string `json:"highlightColor"`
	TitleColor  string `json:"title_color"`
	TextColor   string `json:"text_color"`
	BgColor     string `json:"bg_color"`
	BorderColor string `json:"border_color"`
	IconColor   string `json:"icon_color"`
	LabelColor  string `json:"labelColor"`
	ShowIcons   bool   `json:"show_icons"`
	RankIcon    string `json:"rank_icon"`
}

type Image struct {
	Src   string `json:"src"`
	Width int    `json:"width"`
}

type Badge struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Href string `json:"href"`
	Logo string `json:"logo"`
}

type AboutMe struct {
	List []string `json:"list"`
}

type Skill struct {
	Name         string `json:"name"`
	IsHighlighted bool   `json:"isHighlighted"`
	Logo         string `json:"logo"`
}

type Module struct {
	User             string    `json:"user"`
	StartedProgramming string    `json:"startedProgramming"`
	RepoQuantity      int       `json:"repoQuantity"`
	Header            struct {
		Styles      Styles   `json:"styles"`
		Image       Image    `json:"image"`
		Description string   `json:"description"`
		Badges      []Badge  `json:"badges"`
	} `json:"header"`
	AboutMe    AboutMe           `json:"aboutme"`
	SkillsWall struct {
		Styles       Styles  `json:"styles"`
		RandomOrder  bool    `json:"randomOrder"`
		Skills       []Skill `json:"skills"`
	} `json:"skillswall"`
	RecentWorks   struct {
		Styles Styles `json:"styles"`
	} `json:"recentworks"`
	SocialMedias  struct {
		Styles Styles `json:"styles"`
		Links  []struct {
			Name       string `json:"name"`
			Href       string `json:"href"`
			Logo       string `json:"logo"`
			LabelColor string `json:"labelColor"`
		} `json:"links"`
	} `json:"socialMedias"`
	GithubStats   struct {
		Styles Styles `json:"styles"`
	} `json:"githubStats"`
}

func main() {
	// Create an instance of your struct
	module := Module{}

	// Call the Create function with the struct instance
	result, err := readjson.Create(&module)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// The result will be the same struct with data populated
	resultModule, ok := result.(*Module)
	if !ok {
		fmt.Println("Error: Result is not of the expected type")
		return
	}

	// Print the module
	// Example of accessing the data
	fmt.Println("User:", resultModule.User)
	fmt.Println("Header Description:", resultModule.Header.Description)
	// ... continue accessing the data as needed
}
