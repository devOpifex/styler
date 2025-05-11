package cmd

import (
	"fmt"

	"github.com/devOpifex/styler/options"
)

type mediaMap map[string]string

type Command struct {
	Options    options.Options
	Config     options.Config
	Files      []string
	Strings    []string
	Properties []string
	ClassMap   map[string]string
	MediaMaps  map[string]mediaMap
}

func Run() {
	opts, ok := options.Run()
	if !ok {
		return
	}

	if opts.Create {
		options.Create()
		return
	}

	conf, err := options.Read()

	if err != nil {
		fmt.Println(err)
		return
	}

	mediaMaps := make(map[string]mediaMap)

	for _, m := range conf.Media {
		mediaMaps[m.Name] = make(map[string]string)
	}

	command := Command{
		Options:   opts,
		Config:    conf,
		ClassMap:  make(map[string]string),
		MediaMaps: mediaMaps,
	}

	command.build()
}

func (c *Command) build() {
	err := c.properties()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.read()

	if err != nil {
		fmt.Println(err)
		return
	}

	c.parse()
	c.class()
	c.css()
}
