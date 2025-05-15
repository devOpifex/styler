package options

import (
	"flag"
)

type Options struct {
	Create  bool
	Version bool
}

func Run() Options {
	var create = flag.Bool("create", false, "Create a config file")
	var version = flag.Bool("version", false, "Version of styler")
	flag.Parse()

	return Options{
		Create:  *create,
		Version: *version,
	}
}
