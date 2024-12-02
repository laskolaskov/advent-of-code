package aoc2024

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1part1() int {
	first, second := parseInputDay1part1()

	if len(first) != len(second) {
		log.Fatal("input slices should not be with different size, check your input")
	}

	total := 0

	for i := 0; i < len(first); i++ {
		diff := int(math.Abs(float64(first[i]) - float64(second[i])))
		total += diff
	}

	return total
}

func Day1part2() int {
	first, secondMap := parseInputDay1part2()

	total := 0

	for _, v := range first {
		scale := v * secondMap[v]
		total += scale
	}

	return total
}

func parseInputDay1part1() (sort.IntSlice, sort.IntSlice) {
	file, err := os.Open("./aoc2024/input-day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	first := sort.IntSlice{}
	second := sort.IntSlice{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		spl := strings.Split(line, "   ")
		one, _ := strconv.Atoi(spl[0])
		two, _ := strconv.Atoi(spl[1])
		first = append(first, one)
		second = append(second, two)
	}

	first.Sort()
	second.Sort()

	return first, second
}

func parseInputDay1part2() (sort.IntSlice, map[int]int) {
	file, err := os.Open("./aoc2024/input-day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	first := sort.IntSlice{}
	second := map[int]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		spl := strings.Split(line, "   ")
		one, _ := strconv.Atoi(spl[0])
		two, _ := strconv.Atoi(spl[1])
		//append to first
		first = append(first, one)
		//increase the count of the appearance map for the second column
		second[two]++
	}

	return first, second
}
