package models

type Skills struct {
	Name          string `json:"name"`
	IsHighlighted bool   `json:"isHighlighted,omitempty"`
	Logo          string `json:"logo,omitempty"`
}

type Header struct {
	Styles struct {
		Align     string `json:"align"`
		Style     string `json:"style"`
		Color     string `json:"color"`
		LogoColor string `json:"logoColor"`
	} `json:"styles"`
	Image struct {
		Src   string `json:"src"`
		Width int    `json:"width"`
	} `json:"image"`
	Description string `json:"description"`
	Badges      []struct {
		Type string `json:"type"`
		Name string `json:"name"`
		Href string `json:"href,omitempty"`
		Logo string `json:"logo,omitempty"`
	} `json:"badges"`
}

type AboutMe struct {
	List []string `json:"list"`
}

type SkillsWall struct {
	Styles struct {
		Style          string   `json:"style"`
		Align          string   `json:"align"`
		HighlightColor string   `json:"highlightColor"`
		WallColors     []string `json:"wallColors"`
	} `json:"styles"`
	RandomOrder bool     `json:"randomOrder"`
	Skills      []Skills `json:"skills"`
}

type RecentWorks struct {
	Styles struct {
		TitleColor  string `json:"title_color"`
		TextColor   string `json:"text_color"`
		BgColor     string `json:"bg_color"`
		BorderColor string `json:"border_color"`
		IconColor   string `json:"icon_color"`
	} `json:"styles"`
}

type SocialMedias struct {
	Styles struct {
		Align     string `json:"align"`
		Style     string `json:"style"`
		Color     string `json:"color"`
		LogoColor string `json:"logoColor"`
	} `json:"styles"`
	Links []struct {
		Name       string `json:"name"`
		Href       string `json:"href,omitempty"`
		Logo       string `json:"logo,omitempty"`
		LabelColor string `json:"labelColor,omitempty"`
	} `json:"links"`
}

type GithubStats struct {
	Styles struct {
		Style       string `json:"style"`
		Align       string `json:"align"`
		TitleColor  string `json:"title_color"`
		TextColor   string `json:"text_color"`
		BgColor     string `json:"bg_color"`
		BorderColor string `json:"border_color"`
		ShowIcons   bool   `json:"show_icons"`
		IconColor   string `json:"icon_color"`
		RankIcon    string `json:"rank_icon"`
	} `json:"styles"`
}

type Profile struct {
	User               string       `json:"user"`
	StartedProgramming string       `json:"startedProgramming"`
	RepoQuantity       int          `json:"repoQuantity"`
	Header             Header       `json:"header"`
	AboutMe            AboutMe      `json:"aboutme"`
	SkillsWall         SkillsWall   `json:"skillswall"`
	RecentWorks        RecentWorks  `json:"recentworks"`
	SocialMedias       SocialMedias `json:"socialMedias"`
	GithubStats        GithubStats  `json:"githubStats"`
}

type Data struct {
	GithubStats string
	AboutMe     string
}
