package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type jbox struct {
	x float64
	y float64
	z float64
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(-1)
	}

	jboxes := make([]jbox, 0)
	circuits := make([][]jbox, 0)

	for line := range strings.SplitSeq(strings.Trim(string(f), "\n"), "\n") {
		parts := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(parts[0], 64)
		y, _ := strconv.ParseFloat(parts[1], 64)
		z, _ := strconv.ParseFloat(parts[2], 64)
		circuits = append(circuits, []jbox{{x, y, z}})
		jboxes = append(jboxes, jbox{x, y, z})
	}

	first := jbox{}
	second := jbox{}
	latestMinD := 0.0

	for len(circuits) != 1 {
		minD := 10e9
		first = jbox{}
		second = jbox{}

		// for some reason this O(n^2) code is awfully slow even when n=1000, as i checked it's because golang itself, not an algorithm :(
		for _, i := range jboxes {
			for _, j := range jboxes {
				if d := math.Sqrt((i.x-j.x)*(i.x-j.x) + (i.y-j.y)*(i.y-j.y) + (i.z-j.z)*(i.z-j.z)); d < minD && d > latestMinD {
					first = i
					second = j
					minD = d
				}
			}
		}

		latestMinD = minD
		firstCircuitIndex := -1
		secondCircuitIndex := -1

		for i, circuit := range circuits {
			for _, box := range circuit {
				switch box {
				case first:
					firstCircuitIndex = i
				case second:
					secondCircuitIndex = i
				}
			}
		}

		if firstCircuitIndex == secondCircuitIndex {
			continue
		}

		circuits[firstCircuitIndex] = append(circuits[firstCircuitIndex], circuits[secondCircuitIndex]...)
		circuits = append(circuits[:secondCircuitIndex], circuits[secondCircuitIndex+1:]...)
	}

	fmt.Println(first.x * second.x)
}
