package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Day13part1() int {
	var patterns = scanPatterns("input-day13.txt")
	total := 0

	for _, p := range patterns {

		found := false

		//check rows for point of possible reflection
		possibleRows := checkRows(p)

		//check if they are real reflections and calculate result
		for _, r := range possibleRows {
			if isHorizontalReflection(p, r) {
				total += (r + 1) * 100
				found = true
				break
			}
		}

		if found {
			continue
		}

		//check cols for point of possible reflection
		possibleCols := checkCols(p)

		//check if they are real reflections and calculate result
		for _, c := range possibleCols {
			if isVerticalReflection(p, c) {
				total += c + 1
				break
			}
		}
	}

	return total
}

func Day13part2() int {
	var patterns = scanPatterns("input-day13.txt")
	total := 0
	for _, p := range patterns {

		found := false

		//check rows for point of possible reflection
		possibleRows := checkRowsWithSmudge(p)

		//check if they are real reflections and calculate result
		for _, r := range possibleRows {
			if isHorizontalReflectionWithSmudge(p, r) {
				total += (r + 1) * 100
				found = true
				break
			}
		}

		if found {
			continue
		}

		//check cols for point of possible reflection
		possibleCols := checkColsWithSmudge(p)

		//check if they are real reflections and calculate result
		for _, c := range possibleCols {
			if isVerticalReflectionWithSmudge(p, c) {
				total += c + 1
				break
			}
		}
	}

	return total
}

func isHorizontalReflection(p []string, i int) bool {
	upIndex := i
	downIndex := i + 1
	for {
		upIndex--
		downIndex++
		//check boundaries
		if upIndex < 0 || downIndex >= len(p) {
			break
		}
		up := getRow(p, upIndex)
		down := getRow(p, downIndex)
		if up != down {
			return false
		}
	}
	return true
}

func isHorizontalReflectionWithSmudge(p []string, i int) bool {
	upIndex := i
	downIndex := i + 1
	smudgeCount := 0
	//check initial rows for smudge
	if len(strDiff(getRow(p, upIndex), getRow(p, downIndex))) == 1 {
		smudgeCount++
	}
	for {
		upIndex--
		downIndex++
		//check boundaries
		if upIndex < 0 || downIndex >= len(p) {
			break
		}
		up := getRow(p, upIndex)
		down := getRow(p, downIndex)

		if len(strDiff(up, down)) > 1 {
			return false
		}
		if len(strDiff(up, down)) == 1 {
			smudgeCount++
			if smudgeCount > 1 {
				break
			}
		}
	}
	//return only those with exactly 1 smudge
	return smudgeCount == 1
}

func isVerticalReflection(p []string, i int) bool {
	leftIndex := i
	rightIndex := i + 1
	for {
		leftIndex--
		rightIndex++
		//check boundaries
		if leftIndex < 0 || rightIndex >= len(p[0]) {
			break
		}
		left := getCol(p, leftIndex)
		right := getCol(p, rightIndex)
		if left != right {
			return false
		}
	}
	return true
}

func isVerticalReflectionWithSmudge(p []string, i int) bool {
	leftIndex := i
	rightIndex := i + 1
	smudgeCount := 0
	//check initial rows for smudge
	if len(strDiff(getCol(p, leftIndex), getCol(p, rightIndex))) == 1 {
		smudgeCount++
	}
	for {
		leftIndex--
		rightIndex++
		//check boundaries
		if leftIndex < 0 || rightIndex >= len(p[0]) {
			break
		}
		left := getCol(p, leftIndex)
		right := getCol(p, rightIndex)
		if len(strDiff(left, right)) > 1 {
			return false
		}
		if len(strDiff(left, right)) == 1 {
			smudgeCount++
			if smudgeCount > 1 {
				break
			}
		}

	}
	//return only those with exactly 1 smudge
	return smudgeCount == 1
}

func checkRows(p []string) []int {
	result := []int{}
	rows := len(p)
	for r := 0; r < rows-1; r++ {
		row := getRow(p, r)
		nextRow := getRow(p, r+1)
		if row == nextRow {
			result = append(result, r)
		}
	}
	return result
}

func checkRowsWithSmudge(p []string) []int {
	result := []int{}
	rows := len(p)
	for r := 0; r < rows-1; r++ {
		row := getRow(p, r)
		nextRow := getRow(p, r+1)
		if len(strDiff(row, nextRow)) <= 1 {
			result = append(result, r)
		}
	}
	return result
}

func checkCols(p []string) []int {
	result := []int{}
	cols := len(p[0])
	for c := 0; c < cols-1; c++ {
		col := getCol(p, c)
		nextCol := getCol(p, c+1)
		if col == nextCol {
			result = append(result, c)
		}
	}
	return result
}

func checkColsWithSmudge(p []string) []int {
	result := []int{}
	cols := len(p[0])
	for c := 0; c < cols-1; c++ {
		col := getCol(p, c)
		nextCol := getCol(p, c+1)
		if len(strDiff(col, nextCol)) <= 1 {
			result = append(result, c)
		}
	}
	return result
}

func strDiff(str1 string, str2 string) []int {
	result := []int{}
	//assume both the strings are the same length
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			result = append(result, i)
		}
	}
	return result
}

func getRow(p []string, i int) string {
	return p[i]
}

func getCol(p []string, i int) string {
	sb := strings.Builder{}
	for _, l := range p {
		sb.WriteByte(l[i])
	}
	return sb.String()
}

func scanPatterns(fileName string) [][]string {
	patterns := [][]string{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		} else {
			patterns = append(patterns, lines)
			lines = []string{}
		}
	}

	patterns = append(patterns, lines)

	return patterns
}
