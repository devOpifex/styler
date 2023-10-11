package main

import (
	"fmt"
	"os"
)

func (classes *classMap) css(file string) error {
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
