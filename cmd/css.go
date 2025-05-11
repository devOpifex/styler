package cmd

import (
	"log"
	"os"
)

func (c *Command) css() {
	var css string

	for media, m := range c.MediaMaps {
		css += c.mediaProperty(media)
		for k, v := range m {
			css += "." + k + "{" + v + "}"
		}
		css += "}"
	}

	for k, v := range c.ClassMap {
		css += "." + k + "{" + v + "}"
	}

	os.WriteFile(c.Config.Output, []byte(css), 0644)
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
			return "@media only screen and (max-width: " + m.MinWidth + "){"
		}
	}

	log.Panic(name + " media not found")

	return ""
}
