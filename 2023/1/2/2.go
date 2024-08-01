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

	lines := strings.Split(strings.Trim(input, " \n"), "\n")

	nums := make([]int, len(lines))

	for i, line := range lines {
		nums[i] = getNumber(line)
	}

	sum := utils.Sum(nums)

	fmt.Println(sum)
}

func getNumber(line string) int {
	twoDigits := [2]rune{0, 0}

	for i := 0; i < len(line); i++ {
		fromCurrent := line[i:]

		digitRune := getDigit(fromCurrent)

		if digitRune != rune(0) {
			if twoDigits[0] == 0 {
				twoDigits[0] = digitRune
			}
			twoDigits[1] = digitRune
		}
	}

	numStr := string(twoDigits[:])
	num, err := strconv.Atoi(numStr)
	utils.Check(err)

	return num
}

func getDigit(buf string) rune {
	digitRune := rune(0)

	firstChar := rune(buf[0])

	if unicode.IsDigit(firstChar) {
		digitRune = firstChar
	} else if unicode.IsLetter(firstChar) {
		for j := 3; j <= min(5, len(buf)); j++ {
			found := getDigitFromString(buf[:j])
			if found != 0 {
				digitRune = found
			}
		}
	}

	return digitRune
}

func getDigitFromString(str string) rune {
	digitsWord := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	return digitsWord[str]
}
