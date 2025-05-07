package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func (c *Command) read() error {
	info, err := os.Stat(c.Config.Directory)
	if err != nil {
		return fmt.Errorf("failed to access directory %s: %w", c.Config.Directory, err)
	}

	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", c.Config.Directory)
	}

	return filepath.Walk(c.Config.Directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %s: %w", path, err)
		}

		if info.IsDir() {
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
