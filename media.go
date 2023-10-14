package main

import (
	"regexp"
	"strings"
)

var rMedia = regexp.MustCompile(`sm\:|md\:|lg\:|xl\:|2xl\:`)

var mediaQueries = map[string]string{
	"sm":  "640",
	"md":  "768",
	"lg":  "1024",
	"xl":  "1280",
	"2xl": "1536",
}

func isMedia(s string) bool {
	if strings.Contains(s, "sm\\:") {
		return true
	}

	if strings.Contains(s, "md\\:") {
		return true
	}

	if strings.Contains(s, "lg\\:") {
		return true
	}

	if strings.Contains(s, "xl\\:") {
		return true
	}

	if strings.Contains(s, "2xl\\:") {
		return true
	}

	return false
}

func getMedia(s string) string {
	if strings.Contains(s, "sm:") {
		return "sm"
	}

	if strings.Contains(s, "md:") {
		return "md"
	}

	if strings.Contains(s, "lg:") {
		return "lg"
	}

	if strings.Contains(s, "xl:") {
		return "xl"
	}

	return "2xl"
}

func makeMedia(prefix, class, attr, suffix string) (string, error) {
	media := getMedia(prefix)
	m, _ := mediaQueries[media]
	return "@media (max-width: " + m + "px){" + "." + suffix + prefix + "{" + attr + ";}" + "}", nil
}
