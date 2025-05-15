package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (c *Command) read() error {
	info, err := os.Stat(c.Config.Directory)
	if err != nil {
		return fmt.Errorf("failed to access directory %s: %w", c.Config.Directory, err)
	}

	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", c.Config.Directory)
	}

	// Split the pattern string by the pipe character to support multiple patterns
	patterns := strings.Split(c.Config.Pattern, "|")

	return filepath.Walk(c.Config.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %s: %w", path, err)
		}

		if info.IsDir() {
			return nil
		}

		// Check if the file matches any of the patterns
		fileName := filepath.Base(path)
		matched := false

		for _, pattern := range patterns {
			patternMatched, err := filepath.Match(pattern, fileName)
			if err != nil {
				return fmt.Errorf("error matching pattern %s for file %s: %w", pattern, path, err)
			}
			if patternMatched {
				matched = true
				break
			}
		}

		// Skip files that don't match any pattern
		if !matched {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file %s: %w", path, err)
		}

		c.Files = append(c.Files, string(content))

		return nil
	})
}
