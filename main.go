package main

import (
	"flag"
	"log"
)

func main() {
	var dir = flag.String("dir", "R", "Directory containing files to process")
	var out = flag.String("out", "styles.css", "Path to output CSS")
	flag.Parse()

	classes, err := readFiles(*dir)
	if err != nil {
		log.Fatal(err)
	}

	err = classes.make(*out)
	log.Fatal(err)
}
