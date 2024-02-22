package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
  cmd := exec.Command("git", "shortlog", "-snu")
  output, err := cmd.Output()
  if err != nil {
    fmt.Println("Error: ", err)
    return
  }

  lines := strings.Split(string(output), "\n")

  for _, line := range lines {
    fields := strings.Fields(line)

    if len(fields) < 2 {
      continue
    }

    commits := fields[0]
    author := fields[1]

    fmt.Println("Commit: ")
  }

}
