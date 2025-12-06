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
	operators := strings.Fields(lines[len(lines)-1])
	lines = lines[:len(lines)-1]

	rawOperands := make([][]string, 0)
	columnOperands := make([]string, 0)
	for i := 0; i < len(lines[0]); i++ {
		numberString := ""
		for j := 0; j < len(lines); j++ {
			numberString += string(lines[j][len(lines[0])-1-i])
		}

		numberString = strings.TrimSpace(numberString)
		if numberString == "" {
			rawOperands = append(rawOperands, columnOperands)
			columnOperands = make([]string, 0)
			continue
		}

		columnOperands = append(columnOperands, numberString)
	}

	rawOperands = append(rawOperands, columnOperands)
	operands := make([][]int, 0)

	for _, i := range rawOperands {
		lineOperands := make([]int, 0)
		for _, j := range i {
			number, _ := strconv.Atoi(j)
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

		for j := 0; j < len(operands[len(operands)-1-i]); j++ {
			if operator == "+" {
				columnResult += operands[len(operands)-1-i][j]
				continue
			}

			columnResult *= operands[len(operands)-1-i][j]
		}

		result += columnResult
	}

	fmt.Println(result)
}
