

package builtins

import (
  "fmt"
  "os"
  "os/exec"
)

// RunSH implements the sh command.
func RunSH(args ...string) error {
  
  cmd := exec.Command("sh", args...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Stdin = os.Stdin

  err := cmd.Run()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error running sh: %v\n", err)
  }

  return nil
}
