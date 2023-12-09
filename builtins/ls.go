package builtins

import (
  "fmt"
  "os"
)

// List implements the ls command.
func List(args ...string) error {
  dir := "."
  if len(args) > 0 {
    dir = args[0]
  }

  files, err := os.ReadDir(dir)
  if err != nil {
    return fmt.Errorf("Error reading directory %s: %v", dir, err)
  }

  for _, file := range files {
    fmt.Println(file.Name())
  }

  return nil
}
