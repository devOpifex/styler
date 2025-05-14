package cmd

import (
	"regexp"
	"strings"
)

var pat = regexp.MustCompile("['\"`](.*?)['\"`]")

func (c *Command) Parse() {
	for _, file := range c.Files {
		matches := pat.FindAllStringSubmatch(file, -1)
		for _, match := range matches {
			strs := strings.Split(match[1], " ")
			// only keep strings with - in them
			strs = filter(strs)
			c.Strings = append(c.Strings, strs...)
		}
	}
}

func filter(strs []string) []string {
	var result []string
	for _, str := range strs {
		if !strings.Contains(str, "-") {
			continue
		}
		result = append(result, str)
	}
	return result
}
