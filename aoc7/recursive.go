package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name     string
	weight   int
	children []string
}

func getTowerLevel(text string) (root node) {
	rootName := strings.Split(text, " ")[0]
	weightStr := strings.Split(text, " ")[1]
	weight, _ := strconv.Atoi(weightStr[1 : len(weightStr)-1])
	var children []string
	if strings.Contains(text, "->") {
		children = strings.Split(text, " -> ")
		children = strings.Split(children[1], ", ")
	}
	root = node{name: rootName, weight: weight, children: children}
	return root
}

func (n node) checkChildrenWeight(allNodes map[string]node) (sum int, offset int) {
	var _ = allNodes
	if len(n.children) == 0 {
		return n.weight, 0
	}
	weight := 0
	var prev node
	for i, child := range n.children {
		curr := allNodes[child]
		w, off := curr.checkChildrenWeight(allNodes)
		if off != 0 {
			return w, off
		}
		sum += w
		if weight == 0 {
			weight = w
		} else if w != weight {
			next := allNodes[n.children[(i+1%len(n.children))]]
			nextW, _ := next.checkChildrenWeight(allNodes)
			if nextW == w {
				return 0, prev.weight + w - weight
			} else {
				return 0, curr.weight + weight - w
			}
		}
		prev = allNodes[child]
	}
	return n.weight + sum, 0
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	scanner := bufio.NewScanner(input)

	all := make(map[string]node)
	seen := make(map[string]node)
	var root node
	for scanner.Scan() {
		base := getTowerLevel(scanner.Text())

		all[base.name] = base
		if base.children != nil {
			for _, child := range base.children {
				seen[child] = base
			}
			root = seen[base.name]
		}
	}

	for {
		next, exists := seen[root.name]
		if !exists {
			break
		}
		root = next
	}
	_, diff := root.checkChildrenWeight(all)
	fmt.Println(diff)
}
