package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	var numbers []int
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	n1 := numbers[:len(numbers)/2]
	n2 := numbers[len(numbers)/2:]

	fmt.Println(len(n1), len(n2))

}

func vekkSort(numbers []int) []int {
	var sortedNumbers []int
	lastElement := 0

	for _, element := range numbers {
		if element >= lastElement {
			sortedNumbers = append(sortedNumbers, element)
			lastElement = element
		}
	}

	return sortedNumbers
}
