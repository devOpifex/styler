package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var dir = flag.String("dir", "R", "Directory containing files to process")
	var file = flag.String("file", "", "File to process")
	var out = flag.String("output", "style.min.css", "Path to output CSS file")
	var warn = flag.Bool("warn", false, "Print warnings")
	var verbose = flag.Bool("verbose", false, "Verbose output")
	flag.Parse()

	var classes classMap
	var err error
	if *file == "" {
		if dirNotExists(*dir) {
			fmt.Println("directory `" + *dir + "` does not exist")
			return
		}
		classes, err = readFiles(*dir)
	}

	if *file != "" {
		classes, err = readFile(*file)
	}

	if err != nil {
		log.Fatal(err)
	}

	err = classes.css(*out)
	if err != nil {
		log.Fatal(err)
	}

	classes.verbose(*verbose, *warn)
}
