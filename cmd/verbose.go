package cmd

import (
	"fmt"
)

func (c *Command) verbose() {
	c.count()
	fmt.Printf("✓ File %v generated \n", c.Config.Output)
}

func (c *Command) count() {
	total := len(c.ClassMap)

	for _, m := range c.MediaMaps {
		total += len(m)
	}

	fmt.Printf("Found %v classes \n", total)
}
