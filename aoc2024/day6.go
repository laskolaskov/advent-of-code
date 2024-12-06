package aoc2024

import (
	"bufio"
	"log"
	"os"
)

func Day6part1() int {
	parseInputDay6()
	return 0
}

func parseInputDay6() {
	file, err := os.Open("./aoc2024/input-day5test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//l := scanner.Text()
	}
}
