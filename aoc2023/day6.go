package aoc2023

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day6part1() int {
	result := 1

	races := ParseRaces("./aoc2023/input-day6.txt")

	for _, r := range races {
		isEven := false
		first := r.time / 2
		second := r.time - first
		matches := [][]int{}

		if first == second {
			isEven = true
		}

		for {
			currentResult := first * second
			if currentResult > r.record {
				matches = append(matches, []int{first, second, currentResult, r.record})
				first--
				second++
			} else {
				break
			}
		}

		if isEven {
			result *= len(matches)*2 - 1
		} else {
			result *= len(matches) * 2
		}
	}

	return result // 1312850
}

func Day6part2() int {
	result := 1

	r := ParseRacesPart2("./aoc2023/input-day6.txt")
	fmt.Println(r)

	fmt.Println(r)
	isEven := false
	first := r.time / 2
	second := r.time - first
	if first == second {
		isEven = true
	}
	matches := [][]int{}
	for {
		currentResult := first * second
		if currentResult > r.record {
			matches = append(matches, []int{first, second, currentResult, r.record})
			first--
			second++
		} else {
			fmt.Println("end:: ", first, second, currentResult, r.record)
			break
		}
	}
	if isEven {
		result *= len(matches)*2 - 1
	} else {
		result *= len(matches) * 2
	}
	fmt.Println(first, second, isEven, matches, len(matches))
	fmt.Println(result)

	fmt.Println(result)
	return result // 36749103 again, bruteforcing is too slow // TODO
}

func ParseRaces(fileName string) []RaceRecord {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	times := []int{}
	dists := []int{}
	races := []RaceRecord{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//extract times
		if s, found := strings.CutPrefix(line, "Time: "); found {
			times = MapSlice(reNumbers.FindAllString(s, -1), func(el string) int {
				n, err := strconv.Atoi(el)
				if err != nil {
					log.Fatal(err)
				}
				return n
			})
			continue
		}
		//extract dists
		if s, found := strings.CutPrefix(line, "Distance: "); found {
			dists = MapSlice(reNumbers.FindAllString(s, -1), func(el string) int {
				n, err := strconv.Atoi(el)
				if err != nil {
					log.Fatal(err)
				}
				return n
			})
			continue
		}
	}

	for i := range times {
		races = append(races, RaceRecord{
			time:   times[i],
			record: dists[i],
		})
	}

	return races
}

func ParseRacesPart2(fileName string) RaceRecord {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	times := ""
	dists := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//extract times
		if s, found := strings.CutPrefix(line, "Time: "); found {
			times += strings.ReplaceAll(s, " ", "")
			continue
		}
		//extract dists
		if s, found := strings.CutPrefix(line, "Distance: "); found {
			dists += strings.ReplaceAll(s, " ", "")
			continue
		}
	}

	t, err := strconv.Atoi(times)
	if err != nil {
		log.Fatal(err)
	}
	r, err := strconv.Atoi(dists)
	if err != nil {
		log.Fatal(err)
	}

	return RaceRecord{
		time:   t,
		record: r,
	}
}

type RaceRecord struct {
	time   int
	record int
}
