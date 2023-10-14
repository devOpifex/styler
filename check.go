package main

import (
	"os"
)

func dirNotExists(dir string) bool {
	_, err := os.Stat(dir)
	return err != nil
}
