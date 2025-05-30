package cmd

import (
	"fmt"
	"regexp"
	"strconv"
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

	// First check for color patterns like "color-red-400"
	colorStr, isColor := c.makeColor(str)
	if isColor {
		return colorStr
	}

	// Split by - to get parts
	parts := strings.Split(str, "-")
	if len(parts) <= 1 {
		return str
	}

	// Find the first part that contains a number
	propertyEndIdx := -1
	for i, part := range parts {
		// Check if this part contains a number (either directly or after ~)
		if regexp.MustCompile(`\d`).MatchString(part) || (strings.Contains(part, "~") && regexp.MustCompile(`\d`).MatchString(strings.Split(part, "~")[1])) {
			propertyEndIdx = i
			break
		}
	}

	// If no number found, use traditional approach with last part as value
	if propertyEndIdx == -1 {
		last := strings.LastIndex(str, "-")
		if last == -1 {
			return str
		}
		return str[:last] + ":" + str[last+1:]
	}

	// Join parts before the property end to form the property name
	property := strings.Join(parts[:propertyEndIdx], "-")
	
	// Process each value part individually
	valueParts := parts[propertyEndIdx:]
	processedValues := make([]string, len(valueParts))
	
	for i, part := range valueParts {
		// Check if this part contains a ~ for strict value
		if strings.Contains(part, "~") {
			// Split by ~ and take the strict value
			strictParts := strings.Split(part, "~")
			
			// First part could be empty if ~ is at the beginning
			if strictParts[0] != "" {
				// First part is not strict, process normally
				numValue, err := strconv.Atoi(strictParts[0])
				if err == nil {
					// It's a number, apply divider and unit
					val := float32(numValue) / float32(c.Config.Divider)
					processedValues[i] = fmt.Sprintf("%v%s", val, c.Config.Unit)
				} else {
					// Not a number, keep as is
					processedValues[i] = strictParts[0]
				}
			}
			
			// Second part is strict, don't apply divider or unit
			if len(strictParts) > 1 {
				if processedValues[i] != "" {
					processedValues[i] += " " + strictParts[1]
				} else {
					processedValues[i] = strictParts[1]
				}
			}
		} else {
			// No ~ in this part, process normally
			numValue, err := strconv.Atoi(part)
			if err == nil {
				// It's a number, apply divider and unit
				val := float32(numValue) / float32(c.Config.Divider)
				processedValues[i] = fmt.Sprintf("%v%s", val, c.Config.Unit)
			} else {
				// Not a number, keep as is
				processedValues[i] = part
			}
		}
	}
	
	// Join the processed values with spaces
	return property + ":" + strings.Join(processedValues, " ")
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
		fmt.Printf("%v media not found in .styler\n", strs[0])
		return
	}

	p := c.makeProperty(str)

	if !c.checkProperty(p) {
		return
	}

	c.MediaMaps[strs[0]][c.makeClassName(str)] = p
}

func (c *Command) makeClassName(str string) string {
	str = strings.ReplaceAll(str, "@", "\\@")
	str = strings.ReplaceAll(str, "%", "\\%")
	str = strings.ReplaceAll(str, ":", "\\:")
	str = strings.ReplaceAll(str, "~", "\\~")
	return str
}
