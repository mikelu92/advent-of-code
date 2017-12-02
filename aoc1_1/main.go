package main

import (
	"fmt"
	"io/ioutil"
)

func solution(input []byte) {
	sum := 0
	off := len(input) / 2 // for part 1 use 1 for off
	for i, c := range input {
		n := int(c - '0')
		n2 := int(input[(i+off)%len(input)] - '0')
		if n == n2 {
			sum += n
		}
	}
	fmt.Println(sum)
}

func main() {
	//input := []byte{'9', '1', '2', '9'}
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	input = input[:len(input)-1]
	fmt.Print(input)
	//solution(input)

	prev := int(input[len(input)-1] - '0')
	sum := 0
	for _, char := range input {
		n := int(char - '0')
		fmt.Print(prev)
		fmt.Print(" ")
		fmt.Print(n)
		if prev == n {
			sum += n
			fmt.Print(" ")
			fmt.Print(sum)
		}
		fmt.Println()
		prev = n
	}
	fmt.Println(sum)
}
