package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	data, err := os.Open("input-rain.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	parallelLines := make(map[float64][]line)

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := stringToLine(scanner.Text())
		gradient := gradient(line.p1.x, line.p1.y, line.p2.x, line.p2.y)
		savedLines := parallelLines[gradient]
		savedLines = append(savedLines, line)
		parallelLines[gradient] = savedLines
	}
	largestSet := 0
	for _, value := range parallelLines {
		if count := len(value); count > largestSet {
			largestSet = count
		}
	}
	fmt.Println(largestSet)
}

type coordinate struct {
	x, y float64
}

type line struct {
	p1, p2 coordinate
}

func gradient(x1, y1, x2, y2 float64) float64 {
	return (y2 - y1) / (x2 - x1)
}

func stringToLine(s string) line {
	filter := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	coordinates := strings.FieldsFunc(s, filter)
	x1, _ := strconv.ParseFloat(coordinates[0], 64)
	y1, _ := strconv.ParseFloat(coordinates[1], 64)
	x2, _ := strconv.ParseFloat(coordinates[2], 64)
	y2, _ := strconv.ParseFloat(coordinates[3], 64)
	line := line{
		p1: coordinate{
			x: x1,
			y: y1,
		},
		p2: coordinate{
			x: x2,
			y: y2,
		},
	}
	return line
}
