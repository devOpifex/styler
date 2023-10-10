package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type classes map[string]string

type inputFile struct {
	path     string
	contents []string
}

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

func readFiles(dir string) ([]inputFile, error) {
	var fls []inputFile
	entries, err := os.ReadDir(dir)

	if err != nil {
		return fls, err
	}

	for _, entry := range entries {
		file, err := read(dir + "/" + entry.Name())

		if err != nil {
			continue
		}

		file.findClasses()
		fls = append(fls, file)
	}

	return fls, err
}

func (file inputFile) findClasses() {
	r, _ := regexp.Compile("class[es]?\\s?\\=\\s?['|\"].*['|\"]")
	for _, line := range file.contents {
		classes := r.FindAllString(line, -1)

		if len(classes) == 0 {
			continue
		}
	}
}

func (cls *classes) get(s string) {
	r, _ := regexp.Compile(`[^"]+`)
	matches := r.FindAllString(s, -1)
	fmt.Println(matches)
}
