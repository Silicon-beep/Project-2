package builtins

import (
  "fmt"
  "os"
)

// RemoveFile implements the rm command.
func RemoveFile(args ...string) error {
  if len(args) == 0 {
    return fmt.Errorf("Usage: rm <file1> [<file2> ...]")
  }

  for _, file := range args {
    err := os.Remove(file)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error removing %s: %v\n", file, err)
    }
  }

  return nil
}