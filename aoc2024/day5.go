package aoc2024

import (
	"bufio"
	"lasko/advent-of-code/aoc2023"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day5part1() int {
	rules, updates := parseInputDay5()
	total := 0
	for _, update := range updates {
		//check if each entry has rule for each entry after it
		has := true
		for i := 0; i < len(update)-1; i++ {
			//fmt.Println("i: ", i)
			for j := i + 1; j < len(update); j++ {
				//fmt.Println("j: ", j)
				if !rules[update[i]][update[j]] {
					has = false
					break
				}
			}
			if !has {
				break
			}
		}
		if has {
			mid := (len(update) - 1) / 2
			total += update[mid]
		}
	}
	return total
}

func parseInputDay5() (map[int]map[int]bool, [][]int) {
	file, err := os.Open("./aoc2024/input-day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	first := true

	rules := map[int]map[int]bool{}
	updates := [][]int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		//reach the empty line - start parsing the second part of the input
		if l == "" {
			first = false
			continue
		}
		if first {
			rule := aoc2023.MapSlice(strings.Split(l, "|"), func(el string) int {
				i, _ := strconv.Atoi(el)
				return i
			})
			if rules[rule[0]] == nil {
				rules[rule[0]] = map[int]bool{}
			}
			rules[rule[0]][rule[1]] = true
		} else {
			update := aoc2023.MapSlice(strings.Split(l, ","), func(el string) int {
				i, _ := strconv.Atoi(el)
				return i
			})
			updates = append(updates, update)
		}
	}
	return rules, updates
}
