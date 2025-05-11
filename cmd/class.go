package cmd

import (
	"fmt"
	"regexp"
	"strings"
)

var mediaRegex = regexp.MustCompile("^.*@")
var prefixRegex = regexp.MustCompile("^.*:")

func (c *Command) class() {
	for _, str := range c.Strings {
		t := classType(str)

		if t == "media" {
			c.makeMediaClass(str)
			continue
		}

		c.ClassMap[c.makeClassName(str)] = makeProperty(str)
	}
}

func makeProperty(str string) string {
	str = mediaRegex.ReplaceAllString(str, "")
	str = prefixRegex.ReplaceAllString(str, "")

	last := strings.LastIndex(str, "-")

	if last == -1 {
		return str
	}

	return str[:last] + ":" + str[last+1:]
}

func classType(str string) string {
	if strings.Contains(str, ":") {
		return "prefix"
	}

	if strings.Contains(str, "@") {
		return "media"
	}

	return "normal"
}

func (c *Command) makeMediaClass(str string) {
	strs := strings.Split(str, "@")
	_, ok := c.MediaMaps[strs[0]]

	if !ok {
		fmt.Printf("%v media not found in .styler", strs[0])
		return
	}

	c.MediaMaps[strs[0]][c.makeClassName(str)] = makeProperty(str)
}

func (c *Command) makeClassName(str string) string {
	t := classType(str)

	switch t {
	case "prefix":
		return c.makeClassNamePrefix(str)
	case "media":
		return c.makeClassNameMedia(str)
	default:
		return str
	}
}

func (c *Command) makeClassNameMedia(str string) string {
	return strings.ReplaceAll(str, "@", "\\@")
}

func (c *Command) makeClassNamePrefix(str string) string {
	strs := strings.Split(str, ":")
	return strs[0] + "\\:" + strs[1] + ":" + strs[0]
}
