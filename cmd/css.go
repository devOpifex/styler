package cmd

import (
	"log"
	"os"
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

	for media, m := range c.MediaMaps {
		if len(m) == 0 {
			continue
		}

		css += c.mediaProperty(media)
		for k, v := range m {
			css += "." + k + space + "{" + linebreak + tab + v + linebreak + "}" + linebreak
		}
		css += "}"
	}

	for k, v := range c.ClassMap {
		css += "." + k + space + "{" + linebreak + tab + v + linebreak + "}" + linebreak
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
			return "@media (min-width: " + m.MinWidth + "){"
		}

		if m.MaxWidth != "" {
			return "@media (max-width: " + m.MaxWidth + "){"
		}
	}

	log.Panic(name + " media not found")

	return ""
}

