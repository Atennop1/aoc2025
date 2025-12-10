package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type node struct {
	name    string
	outputs []*node
}

func waysFromOneToAnother(a, b *node) int {
	if a.name == b.name {
		return 1
	}

	result := 0
	for _, n := range a.outputs {
		result += waysFromOneToAnother(n, b)
	}

	return result
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(-1)
	}

	nodes := make([]*node, 0)
	for line := range strings.SplitSeq(strings.Trim(string(f), "\n"), "\n") {
		parts := strings.Split(line, " ")
		inputName := parts[0][:len(parts[0])-1]
		outputNames := parts[1:]

		inputIndex := slices.IndexFunc(nodes, func(n *node) bool { return inputName == n.name })
		if inputIndex == -1 {
			nodes = append(nodes, &node{name: inputName, outputs: make([]*node, 0)})
			inputIndex = len(nodes) - 1
		}

		for _, outputName := range outputNames {
			outputIndex := slices.IndexFunc(nodes, func(n *node) bool { return outputName == n.name })
			if outputIndex == -1 {
				nodes = append(nodes, &node{name: outputName, outputs: make([]*node, 0)})
				outputIndex = len(nodes) - 1
			}

			nodes[inputIndex].outputs = append(nodes[inputIndex].outputs, nodes[outputIndex])
		}
	}

	startNode := nodes[slices.IndexFunc(nodes, func(n *node) bool { return n.name == "you" })]
	endNode := nodes[slices.IndexFunc(nodes, func(n *node) bool { return n.name == "out" })]
	fmt.Println(waysFromOneToAnother(startNode, endNode))
}
