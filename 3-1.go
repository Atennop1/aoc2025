package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.OpenFile("3.txt", os.O_RDONLY, 644)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		maxJoltage := 0
		line := scanner.Text()

		for i := 0; i < len(line)-1; i++ {
			for j := i + 1; j < len(line); j++ {
				joltage, _ := strconv.Atoi(string(line[i]) + string(line[j]))
				maxJoltage = max(maxJoltage, joltage)
			}
		}

		total += maxJoltage
	}

	fmt.Println(total)
}

