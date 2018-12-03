package main

import (
	"fmt"
	"sync"
)

const maxUint = ^uint32(0)

func main() {
	var wg sync.WaitGroup
	primes := sieveOfEratosthenes(512)
	fmt.Println(primes)
	count := 0
	// fmt.Println(isChristmasNumber(25165824, primes))
	for i := 16777216; i <= int(maxUint); i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			isChristmasNumber, primeFactors := isChristmasNumber(i, primes)
			if isChristmasNumber {
				count++
				fmt.Println(i, primeFactors)
			}
		}(i, &wg)
	}

	wg.Wait()

	// fmt.Println(isChristmasNumber(55023912, primes))
	fmt.Println("Number of christmas numbers:", count)
}

func isChristmasNumber(n int, primes []int) (bool, []int) {
	var primeFactors []int
	for _, prime := range primes {
		for {
			if n%prime == 0 {
				n /= prime
				primeFactors = append(primeFactors, prime)
			} else {
				break
			}
		}
	}

	isChristmasNumber := len(primeFactors) == 24

	return isChristmasNumber, primeFactors
}

func sieveOfEratosthenes(n int) []int {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If integers[p] is not changed, then it is a prime
		if integers[p] == true {
			// Update all multiples of p
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	// return all prime numbers <= n
	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}
	return primes
}
