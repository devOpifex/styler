package options

import (
	"flag"
	"fmt"
)

type Options struct {
	Build  bool
	Create bool
}

func Run() (Options, bool) {
	opts := parse()
	if !opts.check() {
		return opts, false
	}

	return opts, true
}

func parse() Options {
	var create = flag.Bool("create", false, "Create a config file")
	var build = flag.Bool("build", false, "Build the CSS")
	flag.Parse()

	return Options{
		Build:  *build,
		Create: *create,
	}
}

func (o Options) check() bool {
	if o.Build && o.Create {
		fmt.Println("You can't use both -build and -create")
		return false
	}

	if !o.Build && !o.Create {
		o.Build = true
	}

	return true
}
