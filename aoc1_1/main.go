package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	input = input[:len(input)-1]

	prev := int(input[len(input)-1] - '0')
	sum := 0
	for _, char := range input {
		n := int(char - '0')
		if prev == n {
			sum += n
		}
		prev = n
	}
	fmt.Println(sum)
}
