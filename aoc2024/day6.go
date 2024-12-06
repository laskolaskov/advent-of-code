package aoc2024

import (
	"bufio"
	"log"
	"os"
)

func Day6part1() int {
	m, pos := parseInputDay6()
	h := len(m)
	w := len(m[0])

	//start by going up
	d := complex(-1, 0)

	//visited
	v := map[complex128]bool{}

	//mark start as visited
	v[pos] = true

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
		switch string(m[ni][nj]) {
		case "#":
			//change direction: always turn right by 90deg
			d = d * complex(0, -1)
		case ".":
			//move in the direction and mark as visited
			pos = n
			v[pos] = true
		default:
			//move in the direction without marking as visited
			pos = n
		}
	}

	return len(v)
}

func parseInputDay6() ([]string, complex128) {
	file, err := os.Open("./aoc2024/input-day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var p complex128
	m := []string{}

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		l := scanner.Text()
		m = append(m, l)
		for j, c := range l {
			if c == '^' {
				p = complex(float64(i), float64(j))
			}
		}
		i++
	}

	return m, p
}
