// This solution was either stolen, adapted or translated from another language by me.
// I keep track of tasks which I failed and this is one of them.

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

func isIntersecting(edges [][2]cell, x1, y1, x2, y2 int) bool {
	for _, edge := range edges {
		ix1, ix2 := min(edge[0].x, edge[1].x), max(edge[0].x, edge[1].x)
		iy1, iy2 := min(edge[0].y, edge[1].y), max(edge[0].y, edge[1].y)

		if x1 < ix2 && x2 > ix1 && y1 < iy2 && y2 > iy1 {
			return true
		}
	}

	return false
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(-1)
	}

	dots := make([]cell, 0)
	for i := range strings.SplitSeq(strings.Trim(string(f), "\n"), "\n") {
		parts := strings.Split(i, ",")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		dots = append(dots, cell{first, second})
	}

	edges := make([][2]cell, 0)
	for i, dot1 := range dots {
		edges = append(edges, [2]cell{dot1, dots[(i+1)%len(dots)]})
	}

	result := 0
	for _, i := range dots {
		for _, j := range dots {
			if isIntersecting(edges, min(i.x, j.x), min(i.y, j.y), max(i.x, j.x), max(i.y, j.y)) {
				continue
			}

			result = max(result, (max(i.x, j.x)-min(i.x, j.x)+1)*(max(i.y, j.y)-min(i.y, j.y)+1))
		}
	}

	fmt.Println(result)
}
