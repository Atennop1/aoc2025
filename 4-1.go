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

	result := 0
	grid := strings.Split(string(f[:len(f)-1]), "\n")

	for i, line := range grid {
		for j, r := range line {
			if r != '@' {
				continue
			}

			rollsAmount := 0
			for iShift := -1; iShift <= 1; iShift++ {
				for jShift := -1; jShift <= 1; jShift++ {
					if (iShift != 0 || jShift != 0) && 0 <= i+iShift && i+iShift < len(grid) && 0 <= j+jShift && j+jShift < len(grid[0]) && grid[i+iShift][j+jShift] == '@' {
						rollsAmount++
					}
				}
			}

			if rollsAmount < 4 {
				result++
			}
		}
	}

	fmt.Println(result)
}
