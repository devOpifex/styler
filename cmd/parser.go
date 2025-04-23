package cmd

import (
	"regexp"
)

var pat = regexp.MustCompile("['\"`](.*?)['\"`]")

func (c Command) parse() {
	for _, file := range c.Files {
		matches := pat.FindAllStringSubmatch(file, -1)
		for _, match := range matches {
			c.Strings = append(c.Strings, match[1])
		}
	}
}
