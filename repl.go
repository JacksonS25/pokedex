package main

import (
	"strings"
)

func cleanInput(text string) []string {
	cleanedText := []string{}

	for _, word := range strings.Fields(text) {
		cleanedText = append(cleanedText, strings.ToLower(word))
	}

	return cleanedText
}
