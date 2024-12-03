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
  lastCommand := true
  for scanner.Scan() {
    line := scanner.Text()
	  re := regexp.MustCompile(`mul\(\d{1,},\d{1,}\)|(do(?:n't)?)`)
    matches := re.FindAllStringSubmatch(line, -1)

    for _, match := range matches {
      if strings.Contains(match[0], "don't") {
        lastCommand = false
      } else if (strings.Contains(match[0], "do")) {
        lastCommand = true
      } else {
        if !lastCommand {
          continue
        }

        numbersParts := strings.Split(match[0][4:len(match[0]) - 1], ",")
        a, _ := strconv.Atoi(strings.TrimSpace(numbersParts[0]))
        b, _ := strconv.Atoi(strings.TrimSpace(numbersParts[1]))
        result += a * b
      }
    }
  }

  log.Printf("Result: %d", result)
}
