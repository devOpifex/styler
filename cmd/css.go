package cmd

import "os"

func (c *Command) css() {
	var css string

	for k, v := range c.ClassMap {
		css += "." + k + "{" + v + "}"
	}

	os.WriteFile(c.Config.Output, []byte(css), 0644)
}
