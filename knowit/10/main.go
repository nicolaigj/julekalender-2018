package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func main() {
	stack := stack.New()
	file, err := os.Open("input.spp")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reader := strings.NewReader(scanner.Text())
	lineLoop:
		for {
			currentRune, _, err := reader.ReadRune()
			if err == io.EOF {
				break
			}
			switch currentRune {
			case ' ':
				stack.Push(31)
			case ':':
				stack.Push(sumNumbersInStack(stack))
			case '|':
				stack.Push(3)
			case '\'':
				a := stack.Pop().(int)
				b := stack.Pop().(int)
				stack.Push(a + b)
			case '.':
				a := stack.Pop().(int)
				b := stack.Pop().(int)
				stack.Push(a - b)
				stack.Push(b - a)
			case '_':
				a := stack.Pop().(int)
				b := stack.Pop().(int)
				stack.Push(a * b)
				stack.Push(a)
			case '/':
				stack.Pop()
			case 'i':
				a := stack.Peek().(int)
				stack.Push(a)
			case '\\':
				a := stack.Pop().(int)
				stack.Push(a + 1)
			case '*':
				a := stack.Pop().(int)
				b := stack.Pop().(int)
				if b == 0 {
					stack.Push(0)
				} else {
					stack.Push(a / b)
				}
			case ']':
				a := stack.Pop().(int)
				if a%2 == 0 {
					stack.Push(1)
				}
			case '[':
				a := stack.Pop().(int)
				if a%2 != 0 {
					stack.Push(a)
				}
			case '~':
				a := returnBiggestOfLastThree(stack)
				stack.Push(a)
			case 'K':
				break lineLoop
			}
		}
	}

	fmt.Println(largestNumberInStack(stack))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func largestNumberInStack(stack *stack.Stack) int {
	var largestNumber int
	for i := stack.Len(); i > 0; i-- {
		currentNumber := stack.Pop().(int)
		if currentNumber > largestNumber {
			largestNumber = currentNumber
		}
	}

	return largestNumber
}

func sumNumbersInStack(stack *stack.Stack) int {
	var sum int
	for i := stack.Len(); i > 0; i-- {
		sum += stack.Pop().(int)
	}
	return sum
}

func returnBiggestOfLastThree(stack *stack.Stack) int {
	var largestNumber int
	for i := 0; i < 3; i++ {
		currentNumber := stack.Pop().(int)
		if currentNumber > largestNumber {
			largestNumber = currentNumber
		}
	}
	return largestNumber
}
