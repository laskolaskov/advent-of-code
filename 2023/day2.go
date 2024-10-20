package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2part1() int {
	config := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	file, err := os.Open("input-day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sp1 := strings.Split(line, ":")
		id, err := strconv.Atoi(strings.TrimPrefix(sp1[0], "Game "))
		if err != nil {
			log.Fatal(err)
		}

		gameInput := strings.Split(sp1[1], ";")
		impossible := false
		for _, game := range gameInput {
			input := strings.Split(strings.Trim(game, " "), ",")
			gameMap := make(map[string]int)
			for _, cubeInfo := range input {
				cube := strings.Split(strings.Trim(cubeInfo, " "), " ")
				cubeCol := cube[1]
				cubeVal, err := strconv.Atoi(cube[0])
				if err != nil {
					log.Fatal(err)
				}
				gameMap[cubeCol] = cubeVal
			}
			for col, val := range gameMap {
				if val > config[col] {
					impossible = true
					break
				}
			}
			if impossible {
				break
			}
		}
		if !impossible {
			result += id
		}
	}
	return result // 2439
}

func Day2part2() int {
	file, err := os.Open("input-day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sp1 := strings.Split(line, ":")

		gameInput := strings.Split(sp1[1], ";")
		fewestPossible := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, game := range gameInput {
			input := strings.Split(strings.Trim(game, " "), ",")
			gameMap := make(map[string]int)
			for _, cubeInfo := range input {
				cube := strings.Split(strings.Trim(cubeInfo, " "), " ")
				cubeCol := cube[1]
				cubeVal, err := strconv.Atoi(cube[0])
				if err != nil {
					log.Fatal(err)
				}
				gameMap[cubeCol] = cubeVal
			}
			for col, val := range gameMap {
				if val > fewestPossible[col] {
					fewestPossible[col] = val
				}
			}
		}
		currentResult := 1
		for _, val := range fewestPossible {
			currentResult = currentResult * val
		}
		result += currentResult
	}
	return result // 63711
}
