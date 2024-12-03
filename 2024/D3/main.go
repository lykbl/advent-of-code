package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
  // file, _ := os.Open("test.txt")
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)

  result := 0
  for scanner.Scan() {
    line := scanner.Text()
	  re := regexp.MustCompile(`mul\((\d{1,}),(\d{1,})\)`)
    matches := re.FindAllStringSubmatch(line, -1)

    for _, groups := range matches {
      a, _ := strconv.Atoi(strings.TrimSpace(groups[1]))
      b, _ := strconv.Atoi(strings.TrimSpace(groups[2]))
      result += a * b
    }
  }

  log.Printf("Result: %d", result)
}
