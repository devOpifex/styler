package main

import (
	"flag"
)

func main() {
	var dir = flag.String("dir", "R", "Directory containing files to process")
	flag.Parse()
	readFiles(*dir)
}
