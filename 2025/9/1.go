package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	x int
	y int
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(-1)
	}

	cells := make([]cell, 0)
	for i := range strings.SplitSeq(strings.Trim(string(f), "\n"), "\n") {
		parts := strings.Split(i, ",")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		cells = append(cells, cell{first, second})
	}

	result := 0
	for _, i := range cells {
		for _, j := range cells {
			result = max(result, (max(i.x, j.x)-min(i.x, j.x)+1)*(max(i.y, j.y)-min(i.y, j.y)+1))
		}
	}

	fmt.Println(result)
}

