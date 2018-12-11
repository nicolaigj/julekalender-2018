package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var sum int
	for i := 1; i <= 18163106; i++ {
		if isZeroHeavy(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func isZeroHeavy(numb int) bool {
	n := strconv.Itoa(numb)
	zeroes := strings.Count(n, "0")
	return zeroes > len(n)-zeroes
}
