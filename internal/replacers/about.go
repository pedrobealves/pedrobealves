package replacers

import (
	"strings"

	"github.com/pedrobealves/pedrobealves/internal/models"
)

func ReplaceAbout(data *models.Profile) string {
	aboutMe := data.AboutMe.List
	return strings.Join(aboutMe, "\n")
}
