package main

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/devOpifex/styler/cmd"
	"github.com/devOpifex/styler/options"
)

func main() {
	fmt.Println("styler initialized!")
	js.Global().Set("goGetCSS", js.FuncOf(getCSS))

	<-make(chan bool)
}

// Helper function to determine class type without calling unexported function
func classType(str string) string {
	if strings.Contains(str, ":") {
		return "prefix"
	}
	if strings.Contains(str, "@") {
		return "media"
	}
	return "normal"
}

func getCSS(this js.Value, args []js.Value) interface{} {
	// Use a global defer/recover to catch any panics
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic in getCSS:", r)
		}
	}()

	// Handle edge case if no arguments are passed
	if len(args) == 0 {
		return js.ValueOf("")
	}

	str := args[0].String()
	if str == "" {
		return js.ValueOf("")
	}

	// Create config
	conf := options.New()

	// Make sure the Colors map is properly initialized
	if conf.Colors == nil {
		conf.Colors = make(map[string]map[string]string)
	}

	mediaMaps := make(map[string]cmd.MediaMap)
	for _, m := range conf.Media {
		mediaMaps[m.Name] = make(map[string]string)
	}

	command := cmd.Command{
		Config:    conf,
		ClassMap:  make(map[string]string),
		MediaMaps: mediaMaps,
		Files:     []string{str}, // Pass the HTML string as a "file" for parsing
	}

	// Load properties first with error handling
	err := command.LoadProperties()
	if err != nil {
		fmt.Println("Error loading properties:", err)
		return js.ValueOf("")
	}

	// Now try to use the standard Build() method with regular Go WASM
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error during Build():", r)
			}
		}()

		command.Parse()

		// Try using the standard Class() method
		command.Class()

		command.Css(true)
	}()

	if command.CSS == "" {
		fmt.Println("No CSS generated")
		return js.ValueOf("")
	}

	return js.ValueOf(command.CSS)
}
