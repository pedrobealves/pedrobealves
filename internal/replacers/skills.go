package replacers

import (
	"math/rand"
	"strings"
	"time"

	"github.com/pedrobealves/pedrobealves/internal/models"
	"github.com/pedrobealves/pedrobealves/internal/utils"
)

func getWallColor(options map[string]bool, wallColors []string, highlightColor string) string {
	isHighlighted, exists := options["isHighlighted"]
	if !exists {
		isHighlighted = false
	}

	if isHighlighted {
		return highlightColor
	}

	rand.Seed(time.Now().UnixNano())
	return wallColors[rand.Intn(len(wallColors))]
}

func ReplaceSkills(data *models.Profile) string {

	align := data.SkillsWall.Styles.Align

	imgSkills := []string{}

	for _, skill := range data.SkillsWall.Skills {
		color := getWallColor(map[string]bool{"isHighlighted": skill.IsHighlighted}, data.SkillsWall.Styles.WallColors, data.SkillsWall.Styles.HighlightColor)

		logo := skill.Logo
		if logo == "" {
			logo = skill.Name
		}

		type ElementData struct {
			Logo      string
			LogoColor string
			Style     string
		}

		properties := ElementData{
			Logo:      logo,
			LogoColor: Colors[data.SkillsWall.Styles.LogoColor],
			Style:     data.SkillsWall.Styles.Style,
		}

		imgSkills = append(imgSkills, utils.GenerateBadge(skill.Name, "", Colors[color], "", properties))
	}

	if data.SkillsWall.RandomOrder {
		rand.Shuffle(len(imgSkills), func(i, j int) {
			imgSkills[i], imgSkills[j] = imgSkills[j], imgSkills[i]
		})
	}

	properties := map[string]string{
		"align": align,
	}

	return utils.GenerateElement("p", properties, strings.Join(imgSkills, "\n"))
}
