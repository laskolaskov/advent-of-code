package aoc2023

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"slices"
)

func Day4part1() int {
	file, err := os.Open("./aoc2023/input-day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		matches := []string{}
		cardResult := 0

		//get the winning and mine numbers for the game card
		spl := strings.Split(strings.Split(line, ":")[1], "|")
		//match and flatten the winning and mine numbers
		winning := []string{}
		for _, a := range reNumbers.FindAllStringSubmatch(spl[0], -1) {
			winning = append(winning, a...)
		}
		mine := []string{}
		for _, a := range reNumbers.FindAllStringSubmatch(spl[1], -1) {
			mine = append(mine, a...)
		}

		//check if mine is among the winning numbers
		for _, number := range mine {
			if c := slices.Contains(winning, number); c {
				matches = append(matches, number)
			}
		}

		//calculate card result
		m := len(matches)
		if m < 3 {
			cardResult = m
		} else {
			cardResult = int(math.Pow(float64(2), float64(m-1)))
		}

		result += cardResult
	}

	return result // 26914
}

func Day4part2() int {
	file, err := os.Open("./aoc2023/input-day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	cardMap := map[int]int{}

	var result int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		matches := []string{}
		//cardResult := 0

		//split the line
		s := strings.Split(line, ":")
		//get card ID
		id, err := strconv.Atoi(reNumbers.FindAllString(s[0], -1)[0])
		if err != nil {
			log.Fatal(err)
		}
		//get the winning and mine numbers for the game card
		spl := strings.Split(s[1], "|")

		//add the original card
		cardMap[id] += 1

		//match and flatten the winning and mine numbers
		winning := []string{}
		for _, a := range reNumbers.FindAllStringSubmatch(spl[0], -1) {
			winning = append(winning, a...)
		}
		mine := []string{}
		for _, a := range reNumbers.FindAllStringSubmatch(spl[1], -1) {
			mine = append(mine, a...)
		}

		//check if mine is among the winning numbers
		for _, number := range mine {
			if c := slices.Contains(winning, number); c {
				matches = append(matches, number)
			}
		}
		//add next cards based on the number of matches and current copies of the card
		l := len(matches)
		for i := id + 1; i <= id+l; i++ {
			cardMap[i] += cardMap[id]
		}
	}

	//calculate result
	for _, n := range cardMap {
		result += n
	}

	return result // 13080971
}
