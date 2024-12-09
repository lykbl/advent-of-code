package main

import (
	"bufio"
	_ "fmt"
	"log"
	"os"
	// "strconv"
)

func main() {
  // file, err := os.Open("test.txt")
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal("FILE")
  }

  scanner := bufio.NewScanner(file)

  fileString := make([]rune, 0)
  for scanner.Scan() {
    line := scanner.Text()

    fileI := 0
    for i, char := range []rune(line) {
      blocksCount := int(char - '0')
      c := '.'
      if i % 2 == 0 {
        c = rune('0' + fileI)
        fileI++
      }

      for y := 0; y < blocksCount; y++ {
        fileString = append(fileString, c)
      }
    }
  }

  left, right := 0, len(fileString) - 1
  for left < right {
    if fileString[left] != '.' {
      left++
      continue
    }

    fileString[left] = fileString[right]
    fileString[right] = '.'
    right--
  }

  checksum := 0
  for i, char := range fileString {
    if char == '.' {
      continue
    }
    checksum += i * int(char - '0')
  }

  log.Printf("%d", checksum)
}
