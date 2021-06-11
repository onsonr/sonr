package util

import "strings"

// ^ Extracts HTTP Rest function from URL
func ExtractHttpFunction(url string) string {
	start := strings.Index(url, "/") + 1
	length := len(url)
	asRunes := []rune(url)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

// ^ Gets Substring from Input
func Substring(input string, start int, length int) string {
	asRunes := []rune(input)
	if start >= len(asRunes) {
		return ""
	}
	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}
	return string(asRunes[start : start+length])
}
