package replacers

import (
	"github.com/pedrobealves/pedrobealves/internal/models"
	"github.com/pedrobealves/pedrobealves/internal/utils"

	"fmt"
)

func ReplaceStats(data *models.Profile) string {

	result := utils.QueryFromObject(data.GithubStats.Styles)

	query := fmt.Sprintf("https://github-readme-stats.vercel.app/api/?username=pedrobealves&%s", result)

	properties := map[string]string{
		"src": query,
	}

	return utils.GenerateElement("img", properties, "")
}
