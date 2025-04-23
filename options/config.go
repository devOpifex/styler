package options

import (
	"encoding/json"
	"fmt"
	"os"
)

var configPath string = ".styler"

type Config struct {
	Pattern   string
	Directory string
	Output    string
}

func new() Config {
	return Config{
		Pattern:   "*.r|*.js|*.html",
		Directory: ".",
		Output:    "style.min.css",
	}
}

func Create() {
	conf := new()
	confJson, _ := json.MarshalIndent(conf, "", "    ")
	err := os.WriteFile(configPath, confJson, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Config file created at", configPath)
}

func Read() (Config, error) {
	var conf Config
	file, err := os.ReadFile(configPath)

	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(file, &conf)

	if err != nil {
		return conf, err
	}

	return conf, nil
}
