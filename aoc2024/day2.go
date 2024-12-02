package aoc2024

import (
	"bufio"
	"fmt"
	"lasko/advent-of-code/aoc2023"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2part1() int {
	reports := parseInputDay2()
	total := 0

	for _, r := range reports {
		order := ""

		if r[0] < r[1] {
			order = "ASC"
		} else if r[0] > r[1] {
			order = "DESC"
		} else {
			//if elements are equal the report is unsafe, skip and check next
			continue
		}

		//assume it is safe
		total++

		for i := 1; i < len(r); i++ {
			a := r[i-1]
			b := r[i]

			//"unsafe" conditions
			if !isSafe(a, b, order) {
				total--
				break
			}
		}
	}
	return total
}

func Day2part2() int {
	reports := parseInputDay2()
	total := 0

	for _, r := range reports {
		fmt.Println(r)
	}
	return total
}

func isSafe(a, b int, order string) bool {
	diff := a - b
	//"unsafe" conditions
	return !((a == b) || (order == "ASC" && a > b) || (order == "DESC" && a < b) || (diff > 3 || diff < -3))
}

func parseInputDay2() [][]int {
	file, err := os.Open("./aoc2024/input-day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		spl := strings.Split(line, " ")
		arr := aoc2023.MapSlice(spl, func(el string) int {
			c, _ := strconv.Atoi(el)
			return c
		})
		r = append(r, arr)
	}

	return r
}
