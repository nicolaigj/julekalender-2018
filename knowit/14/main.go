package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("route.txt")
	if err != nil {
		panic(err)
	}
	// file := "2H2F2H1B3V"

	point := coordinate{
		x: 0,
		y: 0,
	}
	box := box{
		0, 0, 0, 0,
	}
	path := make(map[coordinate]int)
	path[point]++
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
			path = savePath(path, point, pace, true, "x")
			point.moveX(pace)
		case 'V':
			path = savePath(path, point, pace, false, "x")
			point.moveX(-pace)
		case 'F':
			path = savePath(path, point, pace, true, "y")
			point.moveY(pace)
		case 'B':
			path = savePath(path, point, pace, false, "y")
			point.moveY(-pace)
		}
		box.setBounds(point)
	}
	fmt.Printf(
		"[%d,%d], minX=%d, maxX=%d, minY=%d, maxY=%d\n",
		point.x, point.y, box.minX, box.maxX,
		box.minY, box.maxY,
	)
	fmt.Println("Box size:", float64((box.maxX-box.minX+1)*(box.maxY-box.minY+1)))
	fmt.Println("Visited points:", float64(len(path)))
	fmt.Println(path)
	fmt.Println(float64(len(path)) / float64((box.maxX-box.minX+1)*(box.maxY-box.minY+1)-len(path)))
}

type coordinate struct {
	x, y int
}

type box struct {
	minX, maxX, minY, maxY int
}

func (coordinate *coordinate) moveX(pace int) {
	coordinate.x = coordinate.x + pace
}

func (coordinate *coordinate) moveY(pace int) {
	coordinate.y = coordinate.y + pace
}
func (box *box) setBounds(coordinate coordinate) {
	if coordinate.x < box.minX {
		box.minX = coordinate.x
	}
	if coordinate.x > box.maxX {
		box.maxX = coordinate.x
	}
	if coordinate.y < box.minY {
		box.minY = coordinate.y
	}
	if coordinate.y > box.maxY {
		box.maxY = coordinate.y
	}
}

func checkReadError(err error) bool {
	if err == io.EOF {
		return true
	} else if err != nil {
		panic(err)
	}
	return false
}

func savePath(path map[coordinate]int, point coordinate, steps int, forward bool, axis string) map[coordinate]int {
	for i := 0; i < steps; i++ {
		if axis == "x" {
			if forward {
				point.moveX(1)
			} else {
				point.moveX(-1)
			}
		} else if axis == "y" {
			if forward {
				point.moveY(1)
			} else {
				point.moveY(-1)
			}
		}
		path[point]++
	}
	return path
}
