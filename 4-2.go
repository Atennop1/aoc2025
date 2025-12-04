package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("4.txt")
	if err != nil {
		os.Exit(1)
	}

	totalResult := 0
	grid := strings.Split(string(f[:len(f)-1]), "\n")

	for {
		result := 0
		for i, line := range grid {
			for j, r := range line {
				if r != '@' {
					continue
				}

				rollsAmount := 0
				for iShift := -1; iShift <= 1; iShift++ {
					for jShift := -1; jShift <= 1; jShift++ {
						iNew := i + iShift
						jNew := j + jShift

						if (iShift != 0 || jShift != 0) && 0 <= iNew && iNew < len(grid) && 0 <= jNew && jNew < len(grid[0]) && grid[iNew][jNew] == '@' {
							rollsAmount++
						}
					}
				}

				if rollsAmount < 4 {
					grid[i] = grid[i][:j] + string('.') + grid[i][j+1:]
					result++
				}
			}
		}

		totalResult += result
		if result == 0 {
			break
		}
	}

	fmt.Println(totalResult)
}
