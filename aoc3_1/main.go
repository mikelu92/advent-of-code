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
	input := 1024
	// Calculate the square root
	// find the closes odd number
	// traverse the grid back or forward until you find the number
	// calculate distance from center
	result := betterAnswer(input)
	sqrt := math.Sqrt(float64(input))
	length := int(sqrt) + 1
	if length%2 == 0 {
		length++
	}
	side, cursor := getSide(input, length)
	x, y := getCoords(input, length, cursor, side)
	distance := math.Abs(float64(x-length/2)) + math.Abs(float64(y-length/2))
	fmt.Printf("%d,%f\n", result, distance)

}

func betterAnswer(input int) int {
	cursor := 0
	length := 0
	for i := 3; cursor < input; i += 2 {
		cursor = i * i
		length = i
	}
	result := length - 1
	for i := 0; cursor > input; i++ {
		cursor -= length - 1
	}
	for i := 0; cursor < input; i++ {
		if i < length/2 {
			result--
		} else {
			result++
		}
		cursor++
	}
	return result
}
