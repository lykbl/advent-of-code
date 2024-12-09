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
  filesCount := 0
  for scanner.Scan() {
    line := scanner.Text()

    for i, char := range []rune(line) {
      blocksCount := int(char - '0')
      c := '.'
      if i % 2 == 0 {
        c = rune('0' + filesCount)
        filesCount++
      }

      for y := 0; y < blocksCount; y++ {
        fileString = append(fileString, c)
      }
    }
  }

  for fileI := filesCount; fileI > 0; fileI-- {
    fileRune := rune('0' + fileI)

    right := len(fileString) - 1
    for right > 0 {
      if fileString[right] != fileRune {
        right--
        continue
      }

      fileStart := right
      for fileString[right] == fileRune {
        right--
      }
      fileSize := fileStart - right
      left := 0

      for left < right {
        for left < len(fileString) - 1 && fileString[left] != '.' {
          left++
          continue
        }

        bufStart := left
        for fileString[left] == '.' && left <= right {
          left++
        }
        bufSize := left - bufStart

        if fileSize <= bufSize {
          for a := 0; a < fileSize; a++ {
            fileString[bufStart + a] = fileRune
            fileString[fileStart - a] = '.'
          }
          right = -1
          break
        }
      }
    }
  }
  log.Printf("Result: %s", string(fileString))

  checksum := 0
  for i, char := range fileString {
    if char == '.' {
      continue
    }
    checksum += i * int(char - '0')
  }

  log.Printf("%d", checksum)
}
