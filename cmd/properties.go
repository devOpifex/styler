package cmd

import (
	_ "embed"
	"encoding/json"
)

//go:embed properties.json
var properties []byte

func (c *Command) properties() error {
	var props []string
	err := json.Unmarshal(properties, &props)
	if err != nil {
		return err
	}

	c.Properties = props

	return nil
}
