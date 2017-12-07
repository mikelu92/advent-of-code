package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getTowerLevel(text string) (root string, parents []string) {
	root = strings.Split(text, " ")[0]
	if strings.Contains(text, "->") {
		parents = strings.Split(text, " -> ")
		parents = strings.Split(parents[1], ", ")
	}
	return root, parents
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	seen := make(map[string]string)
	var root string
	for scanner.Scan() {
		base, parents := getTowerLevel(scanner.Text())
		if parents != nil {
			for _, parent := range parents {
				seen[parent] = base
			}
			root = seen[base]
		}
	}

	for {
		next, exists := seen[root]
		if !exists {
			fmt.Println(root)
			break
		}
		root = next
	}
}
