package utils

import (
	"net/url"
	"regexp"
	"strings"
)

// BuildURL builds a url given an URL template (that is a URL with some placeholders
// in the form `{foo}`) replacing each of the placeholders with the corresponding
// value in the map `variables`
func BuildURL(urlTemplate string, variables []map[string]string) string {
	for _, mapOfVariables := range variables {
		for key, value := range mapOfVariables {
			value = url.QueryEscape(value)
			urlTemplate = strings.ReplaceAll(urlTemplate, "{"+key+"}", value)
		}
	}

	// Cleaning up missing placeholders
	re := regexp.MustCompile(`{\w+}`)
	urlTemplate = re.ReplaceAllLiteralString(urlTemplate, "")

	return urlTemplate
}
