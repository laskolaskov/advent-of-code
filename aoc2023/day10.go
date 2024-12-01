package aoc2023

import (
	"fmt"
	"log"
)

var tiles = map[rune]map[string]bool{
	'|': {
		"UP":    true,
		"DOWN":  true,
		"LEFT":  false,
		"RIGHT": false,
	},
	'-': {
		"UP":    false,
		"DOWN":  false,
		"LEFT":  true,
		"RIGHT": true,
	},
	'L': {
		"UP":    true,
		"DOWN":  false,
		"LEFT":  false,
		"RIGHT": true,
	},
	'J': {
		"UP":    true,
		"DOWN":  false,
		"LEFT":  true,
		"RIGHT": false,
	},
	'7': {
		"UP":    false,
		"DOWN":  true,
		"LEFT":  true,
		"RIGHT": false,
	},
	'F': {
		"UP":    false,
		"DOWN":  true,
		"LEFT":  false,
		"RIGHT": true,
	},
	'.': {
		"UP":    false,
		"DOWN":  false,
		"LEFT":  false,
		"RIGHT": false,
	},
	//'S': [][]int{[]int{1, 0}, []int{-1, 0}},
}

var reverseDirection = map[string]string{
	"UP":    "DOWN",
	"DOWN":  "UP",
	"RIGHT": "LEFT",
	"LEFT":  "RIGHT",
}

var directions = map[string][]int{
	"UP":    {-1, 0},
	"DOWN":  {1, 0},
	"RIGHT": {0, 1},
	"LEFT":  {0, -1},
}

func GetNext(cell []int, inDir string, lines []string) ([]int, string) {
	var outDir string
	//get tile
	tile := rune(lines[cell[0]][cell[1]])
	//get the connection direction which is not the incoming direction
	for d, v := range tiles[tile] {
		if v && d != inDir {
			outDir = d
			break
		}
	}
	//get the delta for the next cell coords
	delta := directions[outDir]
	//calculate next cell coords
	next := []int{cell[0] + delta[0], cell[1] + delta[1]}
	//return the next cell coords and the reversed outgoing direction
	//ex. if we are going from the cell to the LEFT cell, we are coming to the next cell from the RIGHT
	return next, reverseDirection[outDir]
}

func CheckAdj(cell []int, dir string, lines []string) ([]int, error) {
	//get the movement delta for the direction
	d := directions[dir]
	//get adj cell coords
	adj := []int{cell[0] + d[0], cell[1] + d[1]}
	//check out of bounds
	if adj[0] < 0 || adj[1] < 0 || adj[0] >= len(lines) || adj[1] >= len(lines[0]) {
		return nil, fmt.Errorf("tile out of bounds")
	}
	//check if tile has reverse connection
	tile := rune(lines[adj[0]][adj[1]])
	if !tiles[tile][reverseDirection[dir]] {
		return nil, fmt.Errorf("no possible connection with this tile")
	}
	return adj, nil
}

func Day10part1and2() int {
	// read file and extract maze matrix (strings are slices of runes)
	var lines = scanFile("./aoc2023/input-day10.txt")

	//starting coords
	var start []int
	//starting connections
	var starts [][]int
	//starting incoming directions
	var dirs []string
	//paths
	var p1, p2 [][]int

	//find starting coords
	for i, l := range lines {
		for j, r := range l {
			if r == 'S' {
				start = append(start, i)
				start = append(start, j)
			}
		}
	}

	//find the 2 possible connectins to the start tile
	//check up
	up, _ := CheckAdj(start, "UP", lines)
	if up != nil {
		starts = append(starts, up)
		dirs = append(dirs, "DOWN")
	}
	//check down
	down, _ := CheckAdj(start, "DOWN", lines)
	if down != nil {
		starts = append(starts, down)
		dirs = append(dirs, "UP")
	}
	//check left
	left, _ := CheckAdj(start, "LEFT", lines)
	if left != nil {
		starts = append(starts, left)
		dirs = append(dirs, "RIGHT")
	}
	//check right
	right, _ := CheckAdj(start, "RIGHT", lines)
	if right != nil {
		starts = append(starts, right)
		dirs = append(dirs, "LEFT")
	}

	c1 := starts[0]
	c2 := starts[1]

	d1 := dirs[0]
	d2 := dirs[1]

	//add start cells to paths
	p1 = append(p1, c1)
	p2 = append(p2, c2)

	for {
		//get next cell for both paths
		nextCell1, nextDir1 := GetNext(c1, d1, lines)
		nextCell2, nextDir2 := GetNext(c2, d2, lines)
		//change current to next
		c1 = nextCell1
		c2 = nextCell2
		d1 = nextDir1
		d2 = nextDir2
		//check if each cell is the same as the last in the other path and exit the iteration - odd number of cells in the full loop
		//this should not happen, given how the puzzle is described, the loop will always have even number of cells
		if (c1[0] == p2[len(p2)-1][0] && c1[1] == p2[len(p2)-1][1]) || (c2[0] == p1[len(p1)-1][0] && c2[1] == p1[len(p1)-1][1]) {
			break
		}
		//add them to paths
		p1 = append(p1, c1)
		p2 = append(p2, c2)
		//check if the cells are the same(the paths merge in the same cell to form the path loop) and exit the iteration - even number of cells in the full loop
		if c1[0] == c2[0] && c1[1] == c2[1] {
			break
		}
	}

	if len(p1) != len(p2) {
		log.Fatal("paths are with different lengths")
	}

	//result for part one is the lenght of the either path
	//resultPart1 := len(p1)
	//return resultPart1

	//create full path
	path := p1[:]
	//add second path in reverse
	for i := len(p2) - 1; i >= 0; i-- {
		//skip last (it is the same as last in p1)
		if i == len(p2)-1 {
			continue
		}
		path = append(path, p2[i])
	}
	//add start
	path = append(path, start)

	// https://en.wikipedia.org/wiki/Shoelace_formula
	//calculate area with Shoelace formula
	area := Shoelace(path, 1, 0)

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	//calculate number of inside points with Pick's theorem
	//we can do that, because each point in the path is a point from the maze matrix
	//and is boundary point for the theorem
	insidePoins := InsidePointsPickTheorem(area, len(path))

	return int(insidePoins)
}
