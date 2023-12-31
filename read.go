package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type classMap struct {
	m        map[string]int
	classes  []string
	suffixes []string
	errors   []error
}

type inputFile struct {
	path     string
	contents []string
}

type inputFiles []inputFile

func read(s string) (inputFile, error) {
	inFile := inputFile{
		path: s,
	}

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inFile.contents = append(inFile.contents, scanner.Text())
	}

	return inFile, err
}

func readFile(file string) (classMap, error) {
	var classes classMap
	classes.m = make(map[string]int)

	var fls inputFiles
	fl, err := read(file)
	fls = append(fls, fl)

	classes.parseClasses(fls)

	return classes, err
}

func readFiles(dir string) (classMap, error) {
	var classes classMap
	classes.m = make(map[string]int)

	entries, err := os.ReadDir(dir)

	if err != nil {
		return classes, err
	}

	var fls inputFiles
	for _, entry := range entries {
		file, err := read(dir + "/" + entry.Name())

		if err != nil {
			continue
		}

		fls = append(fls, file)
	}

	classes.parseClasses(fls)

	return classes, err
}

func (classes *classMap) parseClasses(files inputFiles) {
	for _, file := range files {
		classLines := file.findClassLines()
		classes.getClasses(classLines)
	}
}

func (file inputFile) findClassLines() []string {
	var classes []string
	r, _ := regexp.Compile(`class\s*=\s*(['"])([^'"]+)`)
	for _, line := range file.contents {
		withClass := r.FindAllString(line, -1)

		if len(withClass) == 0 {
			continue
		}

		classes = append(classes, withClass...)
	}

	return classes
}

func (classes *classMap) getClasses(lines []string) {
	r, _ := regexp.Compile(`^.*"`)

	for _, line := range lines {
		split := strings.Split(line, "class")
		for _, s := range split {
			trimmed := r.ReplaceAllString(s, "")
			trimmed = strings.TrimSpace(trimmed)

			class := strings.Split(trimmed, " ")
			classes.parse(class)
		}
	}
}

func (classes *classMap) parse(classString []string) {
	for _, class := range classString {
		// it's an empty class
		if class == "" {
			continue
		}

		// it's already mapped
		_, ok := classes.m[class]
		if ok {
			continue
		}

		classes.m[class] = 0

		m, suffix, err := match(class)

		classes.classes = append(classes.classes, m)
		classes.suffixes = append(classes.suffixes, suffix)
		classes.errors = append(classes.errors, err)
	}

}
