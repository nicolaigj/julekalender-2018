package main

import "fmt"

func main() {
	primes := sieveOfEratosthenes(10000000)
	sum := 0
	for i, prime := range primes {
		if i < len(primes)-1 {
			if primes[i+1]-prime == 2 {
				divisors := getDivisors(prime + 1)
				sumDivisors := sumSlice(divisors)
				if prime+1 < sumDivisors {
					sum += (prime + 1)
				}
			}
		}
	}
	fmt.Println(sum)
}

func getDivisors(number int) []int {
	divisors := []int{}
	for i := 1; i <= number/2; i++ {
		if number%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func sumSlice(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
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
