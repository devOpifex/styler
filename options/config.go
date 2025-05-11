package options

import (
	"encoding/json"
	"fmt"
	"os"
)

var configPath string = ".styler"

type Media struct {
	MaxWidth string `json:"maxWidth"`
	MinWidth string `json:"minWidth"`
	Name     string `json:"name"`
}

type Config struct {
	Pattern   string  `json:"pattern"`
	Directory string  `json:"directory"`
	Output    string  `json:"output"`
	Media     []Media `json:"media"`
}

func new() Config {
	return Config{
		Pattern:   "*.r|*.js|*.html",
		Directory: ".",
		Output:    "style.min.css",
		Media: []Media{
			{
				MinWidth: "640px",
				Name:     "sm",
			},
			{
				MinWidth: "768px",
				Name:     "md",
			},
			{
				MinWidth: "1024px",
				Name:     "lg",
			},
			{
				MinWidth: "1280px",
				Name:     "xl",
			},
		},
	}
}

func Create() {
	if _, err := os.Stat(configPath); err == nil {
		fmt.Println(".styler already exists")
		return
	}

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
