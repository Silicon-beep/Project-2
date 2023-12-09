package builtins

import (
	"fmt"
	"os"
)

// MakeDirectory implements the mkdir command.
func MakeDirectory(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Usage: mkdir <directory1> [<directory2> ...]")
	}

	for _, dir := range args {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Error creating directory %s: %v", dir, err)
		}
	}

	return nil
}
