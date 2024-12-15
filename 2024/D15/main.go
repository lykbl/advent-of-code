
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
  // file, err := os.Open("input.txt")
  file, err := os.Open("test.txt")
  if err != nil {
    log.Fatal("FAIL")
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
  }
}

