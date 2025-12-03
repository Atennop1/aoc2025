package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getLargestNumber(str string, requiredDigits int) string {
	if requiredDigits == 0 {
		return ""
	}

	for i := 9; i >= 1; i-- {
		for j, r := range str {
			if string(r) == strconv.Itoa(i) && len(str)-j >= requiredDigits {
				return string(r) + getLargestNumber(str[j+1:], requiredDigits-1)
			}
		}
	}

	return ""
}

func main() {
	f, err := os.OpenFile("3.txt", os.O_RDONLY, 644)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		largestNumber, _ := strconv.Atoi(getLargestNumber(scanner.Text(), 12))
		total += largestNumber
	}

	fmt.Println(total)
}
