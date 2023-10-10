package main

import (
	"fmt"
	"os"
)

func (classes *classMap) make(file string) error {
	var css string
	for _, class := range classes.classes {
		css += class
	}

	for _, err := range classes.errors {
		fmt.Println(err)
	}

	cssByte := []byte(css)
	return os.WriteFile(file, cssByte, 0644)
}
