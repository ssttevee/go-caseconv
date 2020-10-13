package caseconv

import (
	"strings"
	"unicode"
)

func isAllFunc(s string, f func(rune) bool) bool {
	for _, r := range s {
		if !f(r) {
			return false
		}
	}

	return true
}

func titleWord(s string) string {
	if isAllFunc(s, unicode.IsUpper) {
		return s
	}

	s = strings.ToLower(s)
	if s == "id" {
		return "ID"
	}

	if s == "ids" {
		return "IDs"
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func PascalCase(s string) string {
	var result string
	for _, word := range SplitWords(s) {
		result += titleWord(word)
	}

	return result
}

func CamelCase(s string) string {
	words := SplitWords(s)
	if len(words) == 0 {
		return ""
	}

	result := strings.ToLower(words[0])
	for _, word := range words[1:] {
		result += titleWord(word)
	}

	return result
}

func SnakeCase(s string) string {
	return Custom(s, "_")
}

func KebabCase(s string) string {
	return Custom(s, "-")
}

func Custom(s, delimiter string) string {
	var words []string
	for _, word := range SplitWords(s) {
		words = append(words, strings.ToLower(word))
	}

	return strings.Join(words, delimiter)
}
