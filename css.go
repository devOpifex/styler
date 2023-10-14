package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func (classes *classMap) css(file string) error {
	// we need to sort because media queries
	// should be at the end of the file
	sort.SliceStable(classes.classes, func(i, j int) bool {
		return strings.Contains(classes.classes[j], "@media")
	})

	var css string
	var total int
	for _, class := range classes.classes {
		css += class
		total++
	}

	cssByte := []byte(css)

	fmt.Printf("styler generated %v classes in %s\n", total, file)
	return os.WriteFile(file, cssByte, 0644)
}
