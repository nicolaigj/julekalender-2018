package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	numbers := downloadData("https://s3-eu-west-1.amazonaws.com/knowit-julekalender-2018/input-vekksort.txt")
	// numbers := []int{5, 4, 3, 6, 7, 5, 2, 7, 5, 1, 1, 10}

	numbers = vekkSort(numbers)

	sum := sumArray(numbers)

	fmt.Println(sum)

}

func downloadData(url string) []int {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	var numbers []int
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	return numbers
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

func sumArray(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
