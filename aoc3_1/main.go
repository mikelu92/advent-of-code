package main

import (
	"fmt"
	"math"
)

const (
	start = iota
	bottom
	left
	top
	right
)

func getSide(target int, length int) (side int, data int) {
	cursor := length * length
	side = 0
	for ; cursor > target; side++ {
		cursor -= length - 1
	}
	return side, cursor
}

func getCoords(target int, length int, cursor int, side int) (x int, y int) {
	var w, h int
	switch side {
	case start:
	case bottom:
		h = length - 1
	case top:
		w = length - 1
	case right:
		h, w = length-1, length-1
	}
	fmt.Printf("starting coord: %d,%d\n", w, h)

	if side == right && cursor == target {
		w--
		h--
	}

	for cursor < target {
		switch side {
		case start:
		case bottom:
			w++
		case left:
			h++
		case top:
			w--
		case right:
			h--
		}
		cursor++
	}
	return w, h
}

func main() {
	input := 347991
	// Calculate the square root
	// find the closes odd number
	// traverse the grid back or forward until you find the number
	// calculate distance from center
	sqrt := math.Sqrt(float64(input))
	length := int(sqrt) + 1
	if length%2 == 0 {
		length++
	}
	side, cursor := getSide(input, length)
	fmt.Printf("%d,%d\n", side, cursor)
	x, y := getCoords(input, length, cursor, side)
	fmt.Printf("%d,%d\n", x, y)
	distance := math.Abs(float64(x-length/2)) + math.Abs(float64(y-length/2))
	fmt.Println(distance)

}
