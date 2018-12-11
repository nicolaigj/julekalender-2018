package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numbers := []int{1, 2, 3}
	operands := []string{}
	operandList := generate(len(numbers)-1, operands)
	var i int
	for _, element := range operandList {
		wg.Add(1)
		go func(el []string) {
			defer wg.Done()
			i = test(numbers, el)
			fmt.Println(i, el)
		}(element)
	}
	wg.Wait()
}

func generate(length int, operands []string) [][]string {
	var operandList [][]string
	if length == 0 {
		return append(operandList, operands)
	} else {
		o1 := append(operands, "c")
		o2 := append(operands, "+")
		o3 := append(operands, "-")
		return append(operandList,
			append(generate(length-1, o1),
				append(generate(length-1, o2),
					generate(length-1, o3)...)...)...)
	}
}

func test(numbers []int, operands []string) int {
	sum := 0
	var op string
	var concat string
	for index, number := range numbers {
		if index == 0 {
			sum = number
		} else if index < len(numbers)-1 && operands[index] == "c" {
			concat += string(number)
			if operands[index-1] != "c" {
				op = operands[index-1]
			}
		} else {
			if concat != "" {
				number, _ = strconv.Atoi(concat)
				operands[index-1] = op
			}
			switch operands[index-1] {
			case "+":
				sum += number
			case "-":
				sum -= number
			default:
				sum += number
			}
		}
	}
	return sum
}

func sum(sum, number int, operand string) int {
	return 1
}
