package cmd

import (
	"log"
	"os"
)

func (c *Command) Css() {
	var css string

	for media, m := range c.MediaMaps {
		if len(m) == 0 {
			continue
		}

		css += c.mediaProperty(media)
		for k, v := range m {
			css += "." + k + "{" + v + "}"
		}
		css += "}"
	}

	for k, v := range c.ClassMap {
		css += "." + k + "{" + v + "}"
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
			return "@media only screen and (min-width: " + m.MinWidth + "){"
		}

		if m.MaxWidth != "" {
			return "@media only screen and (max-width: " + m.MaxWidth + "){"
		}
	}

	log.Panic(name + " media not found")

	return ""
}

