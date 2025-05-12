package options

import (
	"flag"
)

type Options struct {
	Create bool
}

func Run() Options {
	var create = flag.Bool("create", false, "Create a config file")
	flag.Parse()

	return Options{
		Create: *create,
	}
}
