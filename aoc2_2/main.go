package main

import (
	"bufio"
	"fmt"
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
	return numSlice
}

func getDivisibleValues(row []int) (num int, denom int) {
	fmt.Println(row)
	for i, num1 := range row {
		for _, num2 := range row[i:] {
			if num1 != num2 {
				if num1%num2 == 0 {
					return num1, num2
				} else if num2%num1 == 0 {
					return num2, num1
				}
			}
		}
	}
	return 0, 0
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	sum := 0
	for scanner.Scan() {
		row := toSlice(scanner.Text())

		num, denom := getDivisibleValues(row)
		fmt.Printf("%d,%d\n", num, denom)
		sum += num / denom
	}

	fmt.Println(sum)
}
