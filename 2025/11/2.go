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

func waysFromOneToAnother(a, b *node, memo map[*node]int) int {
	if a.name == b.name {
		return 1
	}

	if cached, ok := memo[a]; ok {
		return cached
	}

	result := 0
	for _, n := range a.outputs {
		result += waysFromOneToAnother(n, b, memo)
	}

	memo[a] = result
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

	svrNode := nodes[slices.IndexFunc(nodes, func(n *node) bool { return n.name == "svr" })]
	fftNode := nodes[slices.IndexFunc(nodes, func(n *node) bool { return n.name == "fft" })]
	dacNode := nodes[slices.IndexFunc(nodes, func(n *node) bool { return n.name == "dac" })]
	outNode := nodes[slices.IndexFunc(nodes, func(n *node) bool { return n.name == "out" })]

	fromSvrToFft := waysFromOneToAnother(svrNode, fftNode, make(map[*node]int))
	fromSvrToDac := waysFromOneToAnother(svrNode, dacNode, make(map[*node]int))
	fromFftToDac := waysFromOneToAnother(fftNode, dacNode, make(map[*node]int))
	fromDacToFft := waysFromOneToAnother(dacNode, fftNode, make(map[*node]int))
	fromDacToOut := waysFromOneToAnother(dacNode, outNode, make(map[*node]int))
	fromFftToOut := waysFromOneToAnother(fftNode, outNode, make(map[*node]int))
	fmt.Println(fromSvrToFft*fromFftToDac*fromDacToOut + fromSvrToDac*fromDacToFft*fromFftToOut)
}
