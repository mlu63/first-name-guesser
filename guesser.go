package guesser

import (
	"strings"
)

func GuessFirstName(name string) string {
	parts := strings.Split(name, " ")
	if len(parts) == 1 { // mononym (ie Madonna, or many Brazilian names)
		return name
	}
	if isPrefix(parts[0]) {
		parts = parts[1:]
	}
	return parts[0]
}

func isPrefix(namePart string) bool {
	namePart = strings.ToLower(namePart)
	prefixes := map[string]bool{
		"dr":   true,
		"mr":   true,
		"mrs":  true,
		"ms":   true,
		"miss": true,
	}
	return prefixes[namePart]
}
