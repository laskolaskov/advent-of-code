package aoc2024

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Match struct {
	Index int
	Type  string
	Val   string
}

func Day3part1() int {
	text := parseInputDay3()
	regex := regexp.MustCompile(`mul\(-?\d+,-?\d+\)`)
	total := 0

	matches := regex.FindAllString(text, -1)
	for _, match := range matches {
		total += parseAndMul(match)
	}

	return total
}

func Day3part2() int {
	text := parseInputDay3()
	total := 0

	rAll := regexp.MustCompile(`mul\(-?\d+,-?\d+\)|do\(\)|don't\(\)`)
	all := rAll.FindAllStringSubmatch(text, -1)

	enabled := true
	for _, match := range all {
		switch match[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default: //it is mul()
			if enabled {
				total += parseAndMul(match[0])
			}
		}
	}

	return total
}

func parseAndMul(s string) int {
	spl := strings.Split(s[4:len(s)-1], ",")
	a, _ := strconv.Atoi(spl[0])
	b, _ := strconv.Atoi(spl[1])
	return a * b
}

func parseInputDay3() string {
	file, err := os.Open("./aoc2024/input-day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		text += line
	}

	return text
}
