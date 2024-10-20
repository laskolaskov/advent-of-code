package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode"

	"golang.org/x/exp/slices"
)

var reNumbers = regexp.MustCompile(`\d+`)
var reAllNumbers = regexp.MustCompile(`-?\d+`)
var reSymbols = regexp.MustCompile("[^0-9.]")
var reGears = regexp.MustCompile(`\*`)

func Day3part1() int {
	lines := scanFile("input-day3.txt")

	result := 0

	for lineIndex, line := range lines {
		//match tokens
		matchedNumbersPos := reNumbers.FindAllStringSubmatchIndex(line, -1)
		matchedNumbersVals := reNumbers.FindAllStringSubmatch(line, -1)

		notTouching := []int{}

		//for each number position in the line
		for numberIndex, numPos := range matchedNumbersPos {
			isTouching := false
			//calc left/right of substring to check
			left := numPos[0] - 1
			right := numPos[len(numPos)-1]
			start := left
			if start < 0 {
				start = 0
			}
			end := right + 1 //+1 to actually include the 'right' index in the subslice from the line
			if end > len(line)-1 {
				end = len(line) - 1
			}

			//check left
			if !isTouching && (left > -1) {
				if c := line[left]; c != '.' {
					isTouching = true
				}
			}
			//check right
			if !isTouching && (right < len(line)) {
				if c := line[right]; c != '.' {
					isTouching = true
				}
			}
			//check above
			if !isTouching && (lineIndex > 0) {
				//the substring from the above row we need to check
				substr := lines[lineIndex-1][start:end]
				matchedSymbolsAbove := reSymbols.FindAllStringSubmatchIndex(substr, -1)
				if len(matchedSymbolsAbove) > 0 {
					isTouching = true
				}
			}
			//check below
			if !isTouching && (lineIndex < len(lines)-1) {
				//the substring from the below row we need to check
				substr := lines[lineIndex+1][start:end]
				matchedSymbolsBelow := reSymbols.FindAllStringSubmatchIndex(substr, -1)
				if len(matchedSymbolsBelow) > 0 {
					isTouching = true
				}
			}

			//if number is not touching any symbol, mark it for removal
			if !isTouching {
				notTouching = append(notTouching, numberIndex)
			}
		}

		//calculate line result
		lineResult := 0
		for i, val := range matchedNumbersVals {
			if c := slices.Contains(notTouching, i); !c {
				intval, err := strconv.Atoi(val[0])
				if err != nil {
					log.Fatal(err)
				}
				lineResult += intval
			}
		}
		//calculate result
		result += lineResult
	}

	return result // 533775
}

func Day3part2() int64 {
	lines := scanFile("input-day3.txt")

	var result int64

	for lineIndex, line := range lines {
		//match tokens
		matchedGears := reGears.FindAllStringSubmatchIndex(line, -1)

		var lineRatioSum int64

		//for each number position in the line
		for _, gearPos := range matchedGears {
			touches := []string{}

			left := gearPos[0] - 1
			right := gearPos[1]
			mid := gearPos[0]

			//check left
			if left > -1 && unicode.IsDigit(rune(line[left])) {
				touches = append(touches, findInStringPositions(line, left))
			}
			//check right
			if right < len(line) && unicode.IsDigit(rune(line[right])) {
				touches = append(touches, findInStringPositions(line, right))
			}
			//check above
			if lineIndex > 0 {
				l := false
				r := false
				//check left
				if left > -1 && unicode.IsDigit(rune(lines[lineIndex-1][left])) {
					touches = append(touches, findInStringPositions(lines[lineIndex-1], left))
					l = true
				}
				//check right
				if right < len(line) && unicode.IsDigit(rune(lines[lineIndex-1][right])) {
					touches = append(touches, findInStringPositions(lines[lineIndex-1], right))
					r = true
				}
				//check middle as well
				m := unicode.IsDigit(rune(lines[lineIndex-1][mid]))
				if m && l && r {
					//if all 3 are touches, it is only one touch, remove last one
					touches = touches[:len(touches)-1]
				}
				if m && !l && !r {
					//if only mid, then it is a touch with single-digit number
					touches = append(touches, findInStringPositions(lines[lineIndex-1], mid))
				}
			}
			//check below
			if lineIndex < len(lines)-1 {
				l := false
				r := false
				//check left
				if left > -1 && unicode.IsDigit(rune(lines[lineIndex+1][left])) {
					touches = append(touches, findInStringPositions(lines[lineIndex+1], left))
					l = true
				}
				//check right
				if right < len(line) && unicode.IsDigit(rune(lines[lineIndex+1][right])) {
					touches = append(touches, findInStringPositions(lines[lineIndex+1], right))
					r = true
				}
				//check middle as well
				m := unicode.IsDigit(rune(lines[lineIndex+1][mid]))
				if m && l && r {
					//if all 3 are touches, it is only one touch, sub 1
					touches = touches[:len(touches)-1]
				}
				if m && !l && !r {
					//if only mid, then it is a touch with single-digit number
					touches = append(touches, findInStringPositions(lines[lineIndex+1], mid))
				}
			}

			//if touches are exactly 2, it is real gear
			if len(touches) == 2 {
				fisrt, _ := strconv.Atoi(touches[0])
				second, _ := strconv.Atoi(touches[1])
				ratio := int64(fisrt) * int64(second)
				lineRatioSum += int64(ratio)
				fmt.Println(gearPos, touches, ratio)
			} else {
				fmt.Println(gearPos, touches)
			}
		}
		result += lineRatioSum
	}

	return result // 78236071
}

func findInStringPositions(line string, pos int) string {
	//match tokens
	matchedNumbersPos := reAllNumbers.FindAllStringSubmatchIndex(line, -1)
	matchedNumbersVals := reAllNumbers.FindAllStringSubmatch(line, -1)

	for i, numPos := range matchedNumbersPos {
		if numPos[0] <= pos && pos <= numPos[1] {
			return matchedNumbersVals[i][0]
		}
	}
	return ""
}
