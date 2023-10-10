package main

import (
	"fmt"
	"os"
)

func (classes *classMap) make(file string, warn bool) error {
	var css string
	var total int
	for _, class := range classes.classes {
		css += class
		total++
	}

	if warn {
		for _, err := range classes.errors {
			if err == nil {
				continue
			}

			fmt.Println(err)
		}
	}

	cssByte := []byte(css)

	fmt.Printf("styler generated %v classes in %s\n", total, file)
	return os.WriteFile(file, cssByte, 0644)
}
