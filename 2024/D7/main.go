package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
  // file, err := os.Open("test.txt")
  file, err := os.Open("input.txt")
  if err != nil {
    panic("A")
  }

  scanner := bufio.NewScanner(file)
  p1 := 0
  for scanner.Scan() {
    line := scanner.Text()

    parts := strings.Split(line, ":")

    result, err := strconv.Atoi(parts[0])
    if err != nil {
      log.Fatal("ATOI")
    }

    target := result
    numbers := []int{}

    numbersParts := strings.Split(strings.TrimSpace(parts[1]), " ")
    for _, number := range numbersParts {
      a, err := strconv.Atoi(strings.TrimSpace(number))
      if err != nil {
        log.Fatal("ATOI")
      }
      numbers = append(numbers, a)
    }

    combinations := GenerateCombinations(len(numbers) - 1)
    for _, combo := range combinations {
      acc := numbers[0]
      for i, op := range(combo) {
        if op == '*' {
          acc = acc * numbers[i + 1]
        }
        if op == '+' {
          acc = acc + numbers[i + 1]
        }
      }
      if acc == target {
        p1 += target
        break
      }
    }
  }

  log.Printf("P1: %d", p1)
}

func GenerateCombinations(length int) [][]rune {
	operators := []rune{'+', '*'}
	totalCombinations := 1
	for i := 0; i < length; i++ {
		totalCombinations *= len(operators)
	}
	
	result := make([][]rune, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		combination := make([]rune, length)
		
		temp := i
		for j := 0; j < length; j++ {
			operatorIndex := temp % len(operators)
			combination[j] = operators[operatorIndex]
			temp /= len(operators)
		}
		
		result[i] = combination
	}

	return result
}
