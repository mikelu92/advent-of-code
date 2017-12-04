package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortString(word string) string {
	array := strings.Split(word, "")
	sort.Strings(array)
	return strings.Join(array, "")
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	result := 0
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		set := make(map[string]bool)
		valid := true
		for _, word := range row {
			sorted := sortString(word) // Part 2
			_, exists := set[sorted]
			if exists {
				valid = false
				break
			} else {
				set[sorted] = true
			}
		}
		if valid {
			result++
		}
	}
	fmt.Println(result)
}
