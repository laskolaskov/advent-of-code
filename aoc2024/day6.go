package aoc2024

import (
	"bufio"
	"log"
	"os"
)

func Day6part1() int {
	m, pos := parseInputDay6()
	path, _ := findPath(m, pos)
	return len(path)
}

func Day6part2() int {
	m, pos := parseInputDay6()
	path, isLoop := findPath(m, pos)
	if isLoop {
		log.Fatal("the first path should not be loop, check your input")
	}
	loops := 0
	//replace each tile of the path after first with obstacle and find new path
	for _, t := range path {
		if t == pos {
			continue
		}
		ti := int(real(t))
		tj := int(imag(t))
		//change tile from the original path to obstacle and find the new path
		m[ti][tj] = []byte("#")[0]
		_, isloop := findPath(m, pos)
		if isloop {
			loops++
		}
		//revert the obstacle back to normal tile
		m[ti][tj] = []byte(".")[0]
	}
	return loops
}

func findPath(m [][]byte, pos complex128) ([]complex128, bool) {
	h := len(m)
	w := len(m[0])

	//start by going up
	d := complex(-1, 0)

	//visited: maps coords to direction of first visit
	v := map[complex128]complex128{}
	//path
	path := []complex128{}

	//mark start as visited and add to path
	v[pos] = complex(-1, 0)
	path = append(path, pos)

	for {
		//get next tile coords
		n := pos + d
		//get cords as integers
		ni := int(real(n))
		nj := int(imag(n))
		//check if out of map
		if ni < 0 || nj < 0 || ni >= h || nj >= w {
			break
		}
		//check tile type
		tile := string(m[ni][nj])
		switch {
		case tile == "#":
			//change direction: always turn right by 90deg
			d = d * complex(0, -1)
		case tile == "." || tile == "^":
			//move in the direction, mark as visited and add to path if not visited already
			pos = n
			if _, ok := v[pos]; !ok {
				v[pos] = d
				path = append(path, pos)
			} else if v[pos] == d { //it is visited, check if the direction is the same
				//if direction is also same, the path is looping
				return path, true
			}
		}
	}

	return path, false
}

func parseInputDay6() ([][]byte, complex128) {
	file, err := os.Open("./aoc2024/input-day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var p complex128
	m := [][]byte{}

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		l := scanner.Text()
		m = append(m, []byte(l))
		for j, c := range l {
			if c == '^' {
				p = complex(float64(i), float64(j))
			}
		}
		i++
	}

	return m, p
}
