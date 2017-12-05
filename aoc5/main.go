package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	array := make([]int, 0)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		array = append(array, num)
	}

	cursor := 0
	steps := 0
	for cursor < len(array) && cursor >= 0 {
		temp := array[cursor]
		if temp < 3 {
			array[cursor] = temp + 1
		} else {
			array[cursor] = temp - 1 // Part 2
		}
		cursor += temp
		steps++
	}
	fmt.Println(steps)
}
