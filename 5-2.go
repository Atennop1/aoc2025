package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

type diapasone struct {
	start int
	end   int
}

func main() {
	f, err := os.ReadFile("5.txt")
	if err != nil {
		os.Exit(-1)
	}

	result := 0
	diapasones := make([]diapasone, 0)

	for _, i := range strings.Split(strings.Split(string(f[:len(f)-1]), "\n\n")[0], "\n") {
		splitted := strings.Split(i, "-")
		left, _ := strconv.Atoi(splitted[0])
		right, _ := strconv.Atoi(splitted[1])

		added := false
		newDiapasone := diapasone{start: left, end: right}

		for j, _ := range diapasones {
			switch {
			case diapasones[j].start == newDiapasone.start:
				diapasones[j].end = max(diapasones[j].end, newDiapasone.end)
				added = true
			case diapasones[j].end == newDiapasone.end:
				diapasones[j].start = min(diapasones[j].start, newDiapasone.start)
				added = true
			case diapasones[j].start == newDiapasone.end:
				diapasones[j].start = newDiapasone.start
				added = true
			case diapasones[j].end == newDiapasone.start:
				diapasones[j].end = newDiapasone.end
				added = true
			}
		}

		if !added {
			diapasones = append(diapasones, newDiapasone)
		}
	}

	startsCount := 0
	lastStart := -1
	idsMap := make(map[int]int, 0)

	for _, i := range diapasones {
		idsMap[i.start] = 1
		idsMap[i.end] = 2

		if i.start == i.end {
			idsMap[i.start] = 3
		}
	}

	for _, i := range slices.Sorted(maps.Keys(idsMap)) {
		switch {
		case idsMap[i] == 1:
			startsCount++
			if lastStart == -1 {
				lastStart = i
			}
		case idsMap[i] == 2:
			startsCount--
			if startsCount == 0 {
				result += i - lastStart + 1
				lastStart = -1
			}
		case idsMap[i] == 3 && startsCount == 0:
			result++
		}
	}

	fmt.Println(result)
}
