package builtins

import (
  "fmt"
  "os"
  "testing"
)

func TouchFile(fileName string) error {
  file, err := os.Create(fileName)
  if err != nil {
    return err
  }
  defer file.Close()
  return nil
}

func TestTouchFile(t *testing.T) {
  fileName := "testFile.txt"
  if err := TouchFile(fileName); err != nil {
    t.Errorf("TouchFile() error = %v, want nil", err)
  }

  // Add assertions to check if the file "testFile.txt" was successfully created
}
