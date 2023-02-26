package main

import (
	"fmt"
	"math"
)

func sieveOfEratosthenes(n int) []int {
	prime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		prime[i] = true
	}

	for p := 2; p <= int(math.Sqrt(float64(n))); p++ {
		if prime[p] == true {
			for i := p * 2; i <= n; i += p {
				prime[i] = false
			}
		}
	}

	primeNumbers := []int{}
	for p := 2; p <= n; p++ {
		if prime[p] == true {
			primeNumbers = append(primeNumbers, p)
		}
	}

	return primeNumbers
}

func main() {
	fmt.Println(sieveOfEratosthenes(50))
}
