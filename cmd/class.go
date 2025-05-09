package cmd

import "strings"

func (c *Command) class() {
	for _, str := range c.Strings {
		c.ClassMap[str] = makeProperty(str)
	}
}

func makeProperty(str string) string {
	last := strings.LastIndex(str, "-")

	if last == -1 {
		return str
	}

	return str[:last] + ":" + str[last+1:]
}
