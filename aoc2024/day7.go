package aoc2024

import (
	"bufio"
	"fmt"
	"lasko/advent-of-code/aoc2023"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day7part1() int { //bruteforce
	input := parseInputDay7()
	total := 0
	for value, numbers := range input {
		operations := []string{}
		p := float64(len(numbers) - 1)
		max := int(math.Pow(2, p))
		found := map[int]bool{}
		//find all possible combinations of 2 operations (+ and *) in N spaces (length of numbers - 1)
		//represented as binary strings with length N
		//ex.: in 3 spaces we have 8 possible combinations of 2 operations
		//numbers from 0 to 7 represented as binary strings give all of these combinations
		//where + and * correspond to 0 and 1
		for i := 0; i < max; i++ {
			//represent as binary string with specific length (p)
			//create dynamic format string of type: %0xb where x = p
			s := fmt.Sprintf("%%0%db\n", int(p)) //ex.: for p=3 => %03b
			//format the number as binary string using the format from above
			operations = append(operations, fmt.Sprintf(s, i))
		}
		//calculate the numbers using each operation string
		//compare the result to the needed value
		for _, o := range operations {
			result := numbers[0]
			for i := 1; i < len(numbers); i++ {
				if o[i-1] == '0' {
					result += numbers[i]
				}
				if o[i-1] == '1' {
					result *= numbers[i]
				}
				if result > value {
					break
				}
			}
			if result == value {
				if !found[value] {
					total += value
					found[value] = true
				}
			}
		}
	}
	return total
}

func Day7part2() int {
	//input := parseInputDay7()
	total := 0
	return total
}

func parseInputDay7() map[int][]int {
	file, err := os.Open("./aoc2024/input-day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := map[int][]int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		spl := strings.Split(l, ": ")
		val, _ := strconv.Atoi(spl[0])
		m[val] = aoc2023.MapSlice(strings.Split(spl[1], " "), func(s string) int {
			n, _ := strconv.Atoi(s)
			return n
		})
	}

	return m
}
