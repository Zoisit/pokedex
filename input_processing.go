package main

import (
	"strings"
)


func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	inputs := strings.Fields(text)

	return inputs
}