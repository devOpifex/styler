package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var dir = flag.String("dir", "R", "Directory containing files to process")
	var out = flag.String("out", "style.min.css", "Path to output CSS file")
	var warn = flag.Bool("warn", false, "Print warnings")
	var verbose = flag.Bool("verbose", false, "Verbose output")
	flag.Parse()

	if dirNotExists(*dir) {
		fmt.Println("directory `" + *dir + "` does not exist")
		return
	}

	classes, err := readFiles(*dir)
	if err != nil {
		log.Fatal(err)
	}

	err = classes.css(*out)
	if err != nil {
		log.Fatal(err)
	}

	classes.verbose(*verbose, *warn)
}
