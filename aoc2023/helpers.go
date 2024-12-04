package aoc2023

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// directions map for {i, j}, where "i" is row and "j" is column index
var Directions = map[string][]int{
	"UP":         {-1, 0},
	"DOWN":       {1, 0},
	"RIGHT":      {0, 1},
	"LEFT":       {0, -1},
	"UP-LEFT":    {-1, -1},
	"DOWN-LEFT":  {1, -1},
	"UP-RIGHT":   {-1, 1},
	"DOWN-RIGHT": {1, 1},
}

func scanFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func scanBytes(fileName string) [][]byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes := [][]byte{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//TODO: check more about this bug
		//reading bytes from scanner and then appending them, results in strange corrupted data (./aoc2023/input-day14test2-bug-with-bytes-slice-append.txt)
		//despite that the line we read is identical(!) to the input, something happens when appending (???)
		/* b := scanner.Bytes()
		fmt.Println(b)
		fmt.Println(string(b))
		bytes = append(bytes, b) (???? seems to mess the data completely - ./aoc2023/input-day14test2-bug-with-bytes-slice-append.txt) */

		l := scanner.Text()
		bytes = append(bytes, []byte(l))
	}

	return bytes
}

func printBytes(state [][]byte) {
	fmt.Println("----------")
	for _, l := range state {
		fmt.Println(string(l))
	}
}

func deepCopy(sl [][]byte) [][]byte {
	r := [][]byte{}
	for _, l := range sl {
		new := append([]byte{}, l...)
		r = append(r, new)
	}
	return r
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// https://en.wikipedia.org/wiki/Shoelace_formula
/*
* Each point is described as array of integers.
* x and y prams show the index of the x and y coords in the point array.
* ex.: x = 1 and y = 0 means the point array looks like [y, x]
 */
func Shoelace(points [][]int, x int, y int) float64 {
	var detSum int
	for i := 0; i < len(points); i++ {
		var p1, p2 []int
		//last index
		if i == len(points)-1 {
			p1 = points[i]
			p2 = points[0] //closing with the first point
		} else {
			p1 = points[i]
			p2 = points[i+1]
		}
		//x1*y2 - x2*y1
		det := p1[x]*p2[y] - p2[x]*p1[y]
		detSum += det
	}
	area := math.Abs(float64(detSum / 2))
	return area
}
func ShoelaceInt64(points [][]int64, x int, y int) float64 {
	var detSum int64
	for i := 0; i < len(points); i++ {
		var p1, p2 []int64
		//last index
		if i == len(points)-1 {
			p1 = points[i]
			p2 = points[0] //closing with the first point
		} else {
			p1 = points[i]
			p2 = points[i+1]
		}
		//x1*y2 - x2*y1
		det := p1[x]*p2[y] - p2[x]*p1[y]
		detSum += det
	}
	area := math.Abs(float64(detSum / 2))
	return area
}

// https://en.wikipedia.org/wiki/Pick%27s_theorem
func InsidePointsPickTheorem(a float64, b int) float64 {
	i := float64(b/2-1) - a
	return math.Abs(i)
}
func InsidePointsPickTheoremInt64(a float64, b int64) float64 {
	i := float64(b/2-1) - a
	return math.Abs(i)
}
