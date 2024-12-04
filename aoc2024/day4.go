package aoc2024

import (
	"bufio"
	"log"
	"os"

	"lasko/advent-of-code/aoc2023"
)

func Day4part1() int {
	text := parseInputDay4()
	word := "XMAS"
	total := 0

	//search for the first letter of the word and check in each direction
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[0]); j++ {
			if text[i][j] == word[0] {
				total += checkWord(i, j, text, word)
			}
		}
	}

	return total
}

func Day4part2() int {
	text := parseInputDay4()
	total := 0

	//search for "A" (65 as byte) and check if it is center of the shape
	//we can skip the first and last row and column here
	//and not check if the shape pieces are in bounds later
	for i := 1; i < len(text)-1; i++ {
		for j := 1; j < len(text[0])-1; j++ {
			if text[i][j] == 65 {
				total += checkMAS(i, j, text)
			}
		}
	}

	return total
}

func checkWord(i, j int, text []string, word string) int {
	c := 0
	for _, dir := range aoc2023.Directions {
		c += checkWordInDir(i, j, dir, text, word)
	}
	return c
}

func checkMAS(i, j int, text []string) int {
	upLeft := string(text[i-1][j-1])
	upRight := string(text[i-1][j+1])
	downLeft := string(text[i+1][j-1])
	downRight := string(text[i+1][j+1])

	if (upLeft == "M" && upRight == "M" && downRight == "S" && downLeft == "S") ||
		(upLeft == "S" && upRight == "M" && downRight == "M" && downLeft == "S") ||
		(upLeft == "S" && upRight == "S" && downRight == "M" && downLeft == "M") ||
		(upLeft == "M" && upRight == "S" && downRight == "S" && downLeft == "M") {
		return 1
	}
	return 0
}

func checkWordInDir(i, j int, dir []int, text []string, search string) int {
	//checking for the other letters after first
	word := search[1:]
	for _, letter := range word {
		i += dir[0]
		j += dir[1]
		if !inBounds(i, j, len(text), len(text[0])) {
			return 0
		}
		if text[i][j] != byte(letter) {
			return 0
		}
	}
	return 1
}

func inBounds(i, j, iMax, jMax int) bool {
	if i >= iMax || i < 0 || j < 0 || j >= jMax {
		return false
	}
	return true
}

func parseInputDay4() []string {
	file, err := os.Open("./aoc2024/input-day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	text := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		text = append(text, line)
	}

	return text
}
