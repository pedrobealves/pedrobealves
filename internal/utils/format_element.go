package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func encodeStr(str string) string {
	lowercaseStr := strings.ToLower(str)
	encodedStr := url.QueryEscape(lowercaseStr)
	return encodedStr
}

func escapeString(s string) string {
	re := regexp.MustCompile("[-_]")
	return re.ReplaceAllStringFunc(s, func(match string) string {
		return strings.Repeat(match, 2)
	})
}

// GenerateBadge generates a badge HTML element.
func GenerateBadge(name, message, color, href string, properties interface{}) string {
	escapedName := escapeString(encodeStr(name))
	escapedMessage := escapeString(encodeStr(message))
	escapedColor := escapeString(NormalizeColor(color))

	nonEmptyStrings := []string{escapedName, escapedMessage, escapedColor}
	filteredStrings := make([]string, 0)

	for _, s := range nonEmptyStrings {
		if s != "" {
			filteredStrings = append(filteredStrings, s)
		}
	}

	path := strings.Join(filteredStrings, "-")

	query := QueryFromObject(properties, false)

	badgeSrc := fmt.Sprintf("https://img.shields.io/badge/%s?%s", path, query)

	badgeProperties := map[string]string{"src": badgeSrc}
	badge := GenerateElement("img", badgeProperties, "")

	if href != "" {
		return GenerateElement("a", map[string]string{"href": href}, badge)
	}
	return badge
}

// GenerateViewsBadge generates a views badge HTML element.
func GenerateViewsBadge(badgeProperties interface{}) string {
	query := QueryFromObject(badgeProperties, false)
	badgeSrc := fmt.Sprintf("https://komarev.com/ghpvc/?%s", query)

	badgePropertiesMap := map[string]string{"src": badgeSrc}
	return GenerateElement("img", badgePropertiesMap, "")
}
