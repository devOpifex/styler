package cmd

import (
	"log"
	"os"
	"strings"
)

func (c *Command) Css(pretty bool) {
	var css string

	linebreak := ""
	tab := ""
	space := ""
	if pretty {
		linebreak = "\n"
		tab = "\t"
		space = " "
	}

	for k, v := range c.ClassMap {
		selector := "." + k
		if strings.Contains(k, "hover\\:") {
			selector += ":hover"
		} else if strings.Contains(k, "active\\:") {
			selector += ":active"
		} else if strings.Contains(k, "focus\\:") {
			selector += ":focus"
		}
		css += selector + space + "{" + linebreak + tab + v + linebreak + "}" + linebreak
	}

	for media, m := range c.MediaMaps {
		if len(m) == 0 {
			continue
		}

		css += c.mediaProperty(media)
		for k, v := range m {
			selector := "." + k
			if strings.Contains(k, "hover\\:") {
				selector += ":hover"
			} else if strings.Contains(k, "active\\:") {
				selector += ":active"
			} else if strings.Contains(k, "focus\\:") {
				selector += ":focus"
			}
			css += selector + space + "{" + linebreak + tab + v + linebreak + "}" + linebreak
		}
		css += "}"
	}

	c.CSS = css
}

func (c *Command) write() {
	os.WriteFile(c.Config.Output, []byte(c.CSS), 0644)
}

func (c *Command) mediaProperty(name string) string {
	for _, m := range c.Config.Media {
		if m.Name != name {
			continue
		}

		if m.MinWidth != "" {
			return "@media(min-width: " + m.MinWidth + "){"
		}

		if m.MaxWidth != "" {
			return "@media(max-width: " + m.MaxWidth + "){"
		}
	}

	log.Panic(name + " media not found")

	return ""
}
