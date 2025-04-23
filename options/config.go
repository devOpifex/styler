package options

import (
	"encoding/json"
	"os"
)

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
	confJson, _ := json.Marshal(conf)
	err := os.WriteFile("config.json", confJson, 0644)

	if err != nil {
		panic(err)
	}
}

func Read() (Config, error) {
	var conf Config
	file, err := os.ReadFile("config.json")

	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(file, &conf)

	if err != nil {
		return conf, err
	}

	return conf, nil
}
