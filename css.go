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

	// "standard" classes
	for i, class := range classes.classes {
		if classes.suffixes[i] != "" {
			continue
		}

		css += class
		total++
	}

	// group media queries
	mq := make(map[string][]string)
	for i, class := range classes.classes {
		if classes.suffixes[i] == "" {
			continue
		}

		mq[classes.suffixes[i]] = append(mq[classes.suffixes[i]], class)
	}

	for key, media := range mq {
		css += "@media (min-width: " + key + "px){"

		for _, class := range media {
			css += class
			total++
		}

		css += "}"
	}

	cssByte := []byte(css)

	fmt.Printf("styler generated %v classes in %s\n", total, file)
	return os.WriteFile(file, cssByte, 0644)
}
