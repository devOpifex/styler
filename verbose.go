package main

import "fmt"

func (classes *classMap) verbose(v, w bool) {
	if !v && !w {
		return
	}

	if v {
		for _, c := range classes.classes {
			fmt.Println(c)
		}
	}

	if !w {
		return
	}

	for _, err := range classes.errors {
		if err == nil {
			continue
		}

		fmt.Println(err)
	}
}
