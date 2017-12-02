package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func toSlice(row string) []int {
	rowSlice := strings.Split(row, "\t")
	numSlice := make([]int, len(rowSlice))
	for i, num := range rowSlice {
		data, _ := strconv.Atoi(num)
		numSlice[i] = data
	}
	fmt.Println(numSlice)
	return numSlice
}

func getExtremes(row []int) (int, int) {
	min := math.MaxInt64
	max := math.MinInt64
	for _, num := range row {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}
	return min, max
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	sum := 0
	for scanner.Scan() {
		row := toSlice(scanner.Text())

		min, max := getExtremes(row)
		sum += max - min
	}

	fmt.Println(sum)
}
