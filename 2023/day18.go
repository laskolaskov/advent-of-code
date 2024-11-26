package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type DigRow struct {
	Dir   string
	Count int64
}

func Day18part1() int64 {
	plan := scanDigPlan("input-day18.txt")

	polygon := [][]int64{}

	start := []int64{0, 0}
	polygon = append(polygon, start)
	boundaryPoints := int64(0)

	for _, l := range plan {
		i := start[0]
		j := start[1]

		switch l.Dir {
		case "L":
			j -= l.Count
		case "R":
			j += l.Count
		case "U":
			i -= l.Count
		case "D":
			i += l.Count
		}
		e := []int64{i, j}
		if i != polygon[0][0] || j != polygon[0][1] {
			polygon = append(polygon, e)
		}
		boundaryPoints += l.Count
		start = e
	}

	area := ShoelaceInt64(polygon, 1, 0)
	insidePoints := InsidePointsPickTheoremInt64(area, boundaryPoints)
	total := boundaryPoints + int64(insidePoints)

	return total
}

func Day18part2() int64 {
	plan := scanDigPlanPart2("input-day18.txt")

	polygon := [][]int64{}

	start := []int64{0, 0}
	polygon = append(polygon, start)
	boundaryPoints := int64(0)

	for _, l := range plan {
		i := start[0]
		j := start[1]

		switch l.Dir {
		case "2": //left
			j -= l.Count
		case "0": //right
			j += l.Count
		case "3": //up
			i -= l.Count
		case "1": //down
			i += l.Count
		}
		e := []int64{i, j}
		if i != polygon[0][0] || j != polygon[0][1] {
			polygon = append(polygon, e)
		}
		boundaryPoints += l.Count
		start = e
	}

	area := ShoelaceInt64(polygon, 1, 0)
	insidePoints := InsidePointsPickTheoremInt64(area, boundaryPoints)
	total := boundaryPoints + int64(insidePoints)

	return total
}

func scanDigPlan(fileName string) []DigRow {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	plan := []DigRow{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()

		spl := strings.Split(l, " ")
		val, _ := strconv.Atoi(spl[1])
		row := DigRow{Dir: spl[0], Count: int64(val)}
		plan = append(plan, row)
	}

	return plan
}

func scanDigPlanPart2(fileName string) []DigRow {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	plan := []DigRow{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()

		spl := strings.Split(l, " ")
		text := spl[2]
		dir := text[7]
		count := text[2:7]

		hex, _ := strconv.ParseInt(count, 16, 64)
		row := DigRow{Dir: string(dir), Count: hex}
		plan = append(plan, row)
	}

	return plan
}

func printPlan(plan []DigRow) {
	fmt.Println("----------")
	for _, l := range plan {
		fmt.Println(l)
	}
}
