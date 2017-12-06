package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var memory []string
	for scanner.Scan() {
		memory = strings.Split(scanner.Text(), "\t")
	}

	memoryNums := make([]int, len(memory))
	seenMemory := make(map[string]bool)
	seenMemory[strings.Join(memory, ",")] = true
	fmt.Println(memory)

	max := 0
	maxIndex := 0
	steps := 0
	for i, word := range memory {
		num, _ := strconv.Atoi(word)
		memoryNums[i] = num
		if num > max {
			max = num
			maxIndex = i
		}
	}
	first := true
	exists := false
	for !exists {
		cursor := maxIndex
		currentMax := memoryNums[maxIndex]
		key := make([]string, len(memory))
		memoryNums[cursor] = 0
		max = 0
		for i := 0; i < currentMax; i++ {
			cursor = (cursor + 1) % len(memoryNums)
			memoryNums[cursor] = memoryNums[cursor] + 1
			if memoryNums[cursor] > max {
				max = memoryNums[cursor]
				maxIndex = cursor
			} else if memoryNums[cursor] == max && maxIndex > cursor {
				maxIndex = cursor
			}
		}

		for i, num := range memoryNums {
			key[i] = strconv.Itoa(num)
		}

		newKey := strings.Join(key, ",")
		fmt.Println(newKey)
		_, exists = seenMemory[newKey]

		steps++
		if exists && first {
			first = false
			seenMemory = make(map[string]bool)
			steps = 0
			exists = false
		}

		seenMemory[newKey] = true
	}

	fmt.Println(steps)

}
