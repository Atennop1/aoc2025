package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("6.txt")
	if err != nil {
		os.Exit(-1)
	}

	lines := strings.Split(strings.Trim(string(f), "\n"), "\n")
	operands := make([][]int, 0)

	operators := strings.Fields(lines[len(lines)-1])
	lines = lines[:len(lines)-1]

	for _, line := range lines {
		lineOperands := make([]int, 0)
		numberStrings := strings.Fields(line)

		for _, numberString := range numberStrings {
			number, _ := strconv.Atoi(numberString)
			lineOperands = append(lineOperands, number)
		}

		operands = append(operands, lineOperands)
	}

	result := 0
	for i, operator := range operators {
		columnResult := 0
		if operator == "*" {
			columnResult = 1
		}

		for j := 0; j < len(operands); j++ {
			if operator == "+" {
				columnResult += operands[j][i]
				continue
			}

			columnResult *= operands[j][i]
		}

		result += columnResult
	}

	fmt.Println(result)
}
