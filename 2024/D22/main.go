package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
  f, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(f)

  secrets := make([]int, 0)
  for scanner.Scan() {
    secret, _ := strconv.Atoi(scanner.Text())
    secrets = append(secrets, secret)
  }

  p1 := 0
  genTimes := 2000
  for i, secret := range secrets {
    fmt.Printf("Buyer: %d, secret: %d\n", i, secret)
    for i := 0; i < genTimes; i++ {
      secret = calcSecret(secret)
    }
    fmt.Printf("Result: %d\n", secret)

    p1 += secret
  }

  fmt.Printf("P1: %d\n", p1)
}

func calcSecret(secret int) int {
	secret = mix(secret, secret * 64)
  secret = prune(secret)

	secret = mix(secret, secret / 32)
	secret = prune(secret)

	secret = mix(secret, secret * 2048)
	secret = prune(secret)

  return secret
}

func mix(secret int, result int) int {
  newSecret := secret ^ result

  return newSecret
}

func prune(secret int) int {
  newSecret := secret % 16777216

  return newSecret
}
