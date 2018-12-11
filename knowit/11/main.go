package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("rute.txt")
	if err != nil {
		panic(err)
	}

	coordinate := coordinate{
		x: 0,
		y: 0,
	}
	dataReader := bufio.NewReader(file)
	for {
		paceRune, _, err := dataReader.ReadRune()
		if checkReadError(err) {
			break
		}
		if paceRune == '\n' {
			break
		}
		pace, err := strconv.Atoi(string(paceRune))
		if err != nil {
			panic(err)
		}
		if checkReadError(err) {
			break
		}
		direction, _, err := dataReader.ReadRune()
		if checkReadError(err) {
			break
		}
		switch direction {
		case 'H':
			coordinate.moveX(pace)
		case 'V':
			coordinate.moveX(-pace)
		case 'F':
			coordinate.moveY(pace)
		case 'B':
			coordinate.moveY(-pace)
		}

	}
	fmt.Printf("[%d,%d]\n", coordinate.x, coordinate.y)
}

type coordinate struct {
	x, y int
}

func (coordinate *coordinate) moveX(pace int) {
	coordinate.x = coordinate.x + pace
}

func (coordinate *coordinate) moveY(pace int) {
	coordinate.y = coordinate.y + pace
}

func checkReadError(err error) bool {
	if err == io.EOF {
		return true
	} else if err != nil {
		panic(err)
	}
	return false
}
