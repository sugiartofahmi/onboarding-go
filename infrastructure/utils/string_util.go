package utils

import (
	"slices"
	"strconv"
	"strings"
)

func StringToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}

func GenerateSlug(text string) string {
	slug := strings.ToLower(text)
	slug = strings.TrimSpace(slug)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}


func Contains(slice []string, str string) bool {
	return slices.Contains(slice, str)
}