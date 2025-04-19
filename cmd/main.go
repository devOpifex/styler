package cmd

import "github.com/devOpifex/styler/options"

type Command struct {
	Options options.Options
	Config  options.Config
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

	command := Command{
		Options: opts,
		Config:  options.Read(),
	}
}
