package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//input := []byte{'9', '1', '2', '9'}
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	input = input[:len(input)-1]

	half := len(input) / 2
	sum := 0
	for i, char := range input {
		n := int(char - '0')
		next := int(input[(i+half)%len(input)] - '0')
		if next == n {
			sum += n
		}
	}
	fmt.Println(sum)
}
