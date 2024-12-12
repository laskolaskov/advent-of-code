package aoc2024

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Day8part1() int { //bruteforce
	h, w, input := parseInputDay8()
	loc := map[string]bool{}

	for _, antennae := range input {
		for i := 0; i < len(antennae)-1; i++ {
			for j := i + 1; j < len(antennae); j++ {
				a := antennae[i]
				b := antennae[j]

				i1 := 2*a[0] - b[0]
				j1 := 2*a[1] - b[1]

				i2 := 2*b[0] - a[0]
				j2 := 2*b[1] - a[1]

				if inBounds(int(i1), int(j1), h, w) {
					k1 := fmt.Sprintf("%d-%d", int(i1), int(j1))
					loc[k1] = true
				}
				if inBounds(int(i2), int(j2), h, w) {
					k2 := fmt.Sprintf("%d-%d", int(i2), int(j2))
					loc[k2] = true
				}
			}
		}
	}
	return len(loc)
}

func Day8part2() int {
	//input := parseInputDay8()
	total := 0
	return total
}

func parseInputDay8() (int, int, map[rune][][]int) {
	file, err := os.Open("./aoc2024/input-day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := map[rune][][]int{}

	scanner := bufio.NewScanner(file)

	width := 0
	i := 0
	for scanner.Scan() {
		l := scanner.Text()
		if i == 0 {
			width = len(l)
		}
		for j, c := range l {
			if c != '.' {
				m[c] = append(m[c], []int{i, j})
			}
		}
		i++
	}

	return i, width, m
}
