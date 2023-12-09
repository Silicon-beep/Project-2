package builtins
import (
  "fmt"
  "os"
)
func TouchFile(args ...string) error {
  for _, file := range args {
    f, err := os.Create(file)
    if err != nil {
      return fmt.Errorf("Error creating file %s: %v", file, err)
    }
    f.Close()
  }
  return nil
}

