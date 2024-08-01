package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input := utils.GetInput()

	lines := strings.Split(input, "\n")

	sum := 0

	for _, line := range lines {
		digits := getDigits(line)
		fmt.Println(string(digits[:]))

		num, _ := strconv.Atoi(string(digits[:]))
		sum += num
	}
	fmt.Println(sum)
}

func getDigits(line string) [2]rune {
	digits := [2]rune{}

	for i := 0; i < len(line); i++ {
		r := rune(line[i])
		if unicode.IsDigit(r) {
			digits[0] = r
			break
		}
	}

	for i := 0; i < len(line); i++ {
		r := rune(line[len(line)-1-i])
		if unicode.IsDigit(r) {
			digits[1] = r
			break
		}
	}

	return digits
}
