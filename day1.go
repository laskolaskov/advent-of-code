package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

func day1part1() int {
	file, err := os.Open("input-day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		var firstDigit, lastDigit int
		var lineResult int

		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				firstDigit = int(runes[i] - '0')
				break
			}
		}
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				lastDigit = int(runes[i] - '0')
				break
			}
		}
		lineResult = firstDigit*10 + lastDigit
		result += lineResult
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result //55130
}

func day1part2() int {
	tokens := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	file, err := os.Open("input-day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		var firstDigit, lastDigit int
		firstPosition := len(line)
		lastPosition := 0
		var lineResult int

		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				firstDigit = int(runes[i] - '0')
				firstPosition = i
				break
			}
		}
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				lastDigit = int(runes[i] - '0')
				lastPosition = i
				break
			}
		}
		for t, v := range tokens {
			pos := strings.Index(line, t)
			if pos > -1 && pos < firstPosition {
				firstDigit = v
				firstPosition = pos
			}
			if pos > -1 && pos > lastPosition {
				lastDigit = v
				lastPosition = pos
			}

			pos = strings.LastIndex(line, t)
			if pos > -1 && pos < firstPosition {
				firstDigit = v
				firstPosition = pos
			}
			if pos > -1 && pos > lastPosition {
				lastDigit = v
				lastPosition = pos
			}
		}
		lineResult = firstDigit*10 + lastDigit
		result += lineResult
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result //54985
}
