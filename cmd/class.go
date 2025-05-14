package cmd

import (
	"fmt"
	"regexp"
	"strings"
)

var mediaRegex = regexp.MustCompile("^.*@")
var prefixRegex = regexp.MustCompile("^.*:")

func (c *Command) Class() {
	for _, str := range c.Strings {
		t := classType(str)

		if t == "media" {
			c.makeMediaClass(str)
			continue
		}

		p := c.makeProperty(str)

		if !c.checkProperty(p) {
			continue
		}

		c.ClassMap[c.makeClassName(str)] = p
	}
}

func (c *Command) checkProperty(str string) bool {
	if len(c.Properties) == 0 {
		return true
	}

	property := strings.Split(str, ":")
	if len(property) == 0 {
		return false
	}

	for _, p := range c.Properties {
		if p == property[0] {
			return true
		}
	}
	return false
}

func (c *Command) makeProperty(str string) string {
	// remove media and prefix from attribute
	str = mediaRegex.ReplaceAllString(str, "")
	str = prefixRegex.ReplaceAllString(str, "")

	last := strings.LastIndex(str, "-")

	if last == -1 {
		return str
	}

	value := str[last+1:]

	var intValue int
	_, err := fmt.Sscanf(value, "%d", &intValue)

	if err == nil {
		// it ends in a number it may be a color, e.g.: color-red-400
		str, ok := c.makeColor(str)
		if ok {
			return str
		}
		val := float32(intValue) / float32(c.Config.Divider)
		return str[:last] + ":" + fmt.Sprintf("%v", val) + c.Config.Unit
	}

	return str[:last] + ":" + value
}

func (c *Command) makeColor(str string) (string, bool) {
	tokens := strings.Split(str, "-")

	if len(tokens) < 3 {
		return str, false
	}

	m, ok := c.Config.Colors[tokens[len(tokens)-2]]

	if !ok {
		return str, false
	}

	if color, exists := m[tokens[len(tokens)-1]]; exists {
		attr := strings.Join(tokens[:len(tokens)-2], "-")
		return attr + ":" + color, true
	}

	return str, false
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
	if len(strs) == 0 {
		return
	}

	_, ok := c.MediaMaps[strs[0]]

	if !ok {
		fmt.Printf("%v media not found in .styler", strs[0])
		return
	}

	p := c.makeProperty(str)

	if !c.checkProperty(p) {
		return
	}

	c.MediaMaps[strs[0]][c.makeClassName(str)] = p
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
	if len(strs) < 2 {
		return str
	}
	return strs[0] + "\\:" + strs[1] + ":" + strs[0]
}
