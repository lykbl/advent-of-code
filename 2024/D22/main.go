package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
  // f, _ := os.Open("test.txt")
  f, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(f)

  secrets := make([]int, 0)
  for scanner.Scan() {
    secret, _ := strconv.Atoi(scanner.Text())
    secrets = append(secrets, secret)
  }

  p1 := 0
  genTimes := 2000
  for _, secret := range secrets {
    for i := 0; i < genTimes; i++ {
      secret = calcSecret(secret)
    }

    p1 += secret
  }
  fmt.Printf("P1: %d\n", p1)

  buyerSecrets := [][]int{}
  for buyerI, secret := range secrets {
    buyerSecrets = append(buyerSecrets, []int{secret})
    for i := 1; i < genTimes; i++ {
      buyerSecrets[buyerI] = append(buyerSecrets[buyerI], calcSecret(buyerSecrets[buyerI][i - 1]))
    }
  }
  
  prices := make([][]int, 0)
  for buyerI, secrets := range buyerSecrets {
    prices = append(prices, []int{})
    for _, secret := range secrets {
      prices[buyerI] = append(prices[buyerI], secret % 10)
    }
  }

  priceChanges := make([][]int, 0)
  for buyerI, buyerPrices := range prices {
    priceChanges = append(priceChanges, []int{})
    lastPrice := 0
    for _, price := range buyerPrices {
      priceChanges[buyerI] = append(priceChanges[buyerI], price - lastPrice)
      lastPrice = price
    }
    priceChanges[buyerI][0] = 0
  }

  sellingPrices := make(map[string]int, 0)
  maxProfit := 0
  for buyerI, buyerPriceChanges := range priceChanges {
    checkedRanges := map[string]struct{}{}
    for i := 1; i < len(buyerPriceChanges) - 3; i++ {
      curRange := buyerPriceChanges[i:i + 4]
      rangeKey := ""
      for _, v := range curRange {
        rangeKey += strconv.Itoa(v)
      }
      if _, alreadyChecked := checkedRanges[rangeKey]; alreadyChecked {
        continue
      }
      checkedRanges[rangeKey] = struct{}{}
      sellingPrices[rangeKey] += prices[buyerI][i + 3]
      maxProfit = int(math.Max(float64(maxProfit), float64( sellingPrices[rangeKey] )))
    }
  }

  fmt.Printf("Max profit: %d\n", maxProfit)
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
