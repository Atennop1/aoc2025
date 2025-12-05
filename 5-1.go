package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("5.txt")
	if err != nil {
		os.Exit(-1)
	}

	parts := strings.Split(string(f[:len(f)-1]), "\n\n")
	diapasoneStrings := strings.Split(parts[0], "\n")
	idStrings := strings.Split(parts[1], "\n")

	diapasones := make([][2]int, 0)
	ids := make([]int, 0)
	result := 0

	for _, i := range diapasoneStrings {
		splitted := strings.Split(i, "-")
		left, _ := strconv.Atoi(splitted[0])
		right, _ := strconv.Atoi(splitted[1])
		diapasones = append(diapasones, [2]int{left, right})
	}

	for _, i := range idStrings {
		number, _ := strconv.Atoi(i)
		ids = append(ids, number)
	}

	for _, i := range ids {
		for _, j := range diapasones {
			if j[0] <= i && i <= j[1] {
				result++
				break
			}
		}
	}

	fmt.Println(result)
}
