package cmd

import (
	"fmt"

	"github.com/devOpifex/styler/options"
)

type MediaMap map[string]string

type Command struct {
	Options    options.Options
	Config     options.Config
	Files      []string
	Strings    []string
	Properties []string
	ClassMap   map[string]string
	MediaMaps  map[string]MediaMap
	CSS        string
}

func Run() {
	opts := options.Run()

	if opts.Version {
		fmt.Printf("Styler version %v\n", options.Version)
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

	mediaMaps := make(map[string]MediaMap)

	for _, m := range conf.Media {
		mediaMaps[m.Name] = make(map[string]string)
	}

	command := Command{
		Options:   opts,
		Config:    conf,
		ClassMap:  make(map[string]string),
		MediaMaps: mediaMaps,
	}

	err = command.LoadProperties()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = command.read()

	if err != nil {
		fmt.Println(err)
		return
	}

	command.Parse()
	command.Class()
	command.Css(false)

	command.write()
	command.verbose()
}
