package aoc2023

import (
	"slices"
)

var LEFT string = "left"
var RIGHT string = "right"
var UP string = "up"
var DOWN string = "down"

var LS byte = byte('\\')
var RS byte = byte('/')
var HS byte = byte('-')
var VS byte = byte('|')
var G byte = byte('.')

var energized = [][]int{}
var energizedFrom = []string{}

func Day16part1() int {
	state := scanBytes("./aoc2023/input-day16.txt")

	//reset globals
	energized = [][]int{}
	energizedFrom = []string{}

	tile := []int{0, 0}
	dir := RIGHT

	scanTile(tile, dir, state)

	return len(unique(energized))
}

func Day16part2() int {
	state := scanBytes("./aoc2023/input-day16.txt")

	best := 0

	//top row starters
	for j := 0; j < len(state[0]); j++ {
		//initialize
		i := 0      //row is fixed to first one
		dir := DOWN //direction is also fixed
		//reset globals
		energized = [][]int{}
		energizedFrom = []string{}

		start := []int{i, j}

		scanTile(start, dir, state)

		result := len(unique(energized))
		if result > best {
			best = result
		}
	}

	//bottom row starters
	for j := 0; j < len(state[0]); j++ {
		//initialize
		i := len(state) //row is fixed to last one
		dir := UP       //direction is also fixed
		//reset globals
		energized = [][]int{}
		energizedFrom = []string{}

		start := []int{i, j}

		scanTile(start, dir, state)

		result := len(unique(energized))
		if result > best {
			best = result
		}
	}

	//first column starters
	for i := 0; i < len(state); i++ {
		//initialize
		j := 0       //col is fixed to first one
		dir := RIGHT //direction is also fixed
		//reset globals
		energized = [][]int{}
		energizedFrom = []string{}

		start := []int{i, j}

		scanTile(start, dir, state)

		result := len(unique(energized))
		if result > best {
			best = result
		}
	}

	//last column starters
	for i := 0; i < len(state); i++ {
		//initialize
		j := len(state[0]) //col is fixed to last one
		dir := LEFT        //direction is also fixed
		//reset globals
		energized = [][]int{}
		energizedFrom = []string{}

		start := []int{i, j}

		scanTile(start, dir, state)

		result := len(unique(energized))
		if result > best {
			best = result
		}
	}

	//feels a bit slow, but still bearable (~8-9 seconds ??)
	return best
}

func scanTile(tile []int, dir string, state [][]byte) {
	//check if still in state bounds
	if tile[0] < 0 || tile[0] >= len(state) || tile[1] < 0 || tile[1] >= len(state[0]) {
		//out of bounds, stop progress on this beam
		return
	}
	//check if already energized
	for i := 0; i < len(energized); i++ {
		if energized[i][0] == tile[0] && energized[i][1] == tile[1] {
			//isEnergized = true
			//if already energized from the same direction, don't need to go more, another beam already scanned this tile
			if energizedFrom[i] == dir {
				//stop beam
				return
			}
		}
	}
	//energize
	energized = append(energized, tile)
	energizedFrom = append(energizedFrom, dir)
	//switch by tile symbol
	t := state[tile[0]][tile[1]]
	switch t {
	case G:
		//go to next tile, without changing direction
		next := getNext(tile, dir)
		scanTile(next, dir, state)
	case LS:
		//switch direction
		switch dir {
		case UP:
			dir = LEFT
		case DOWN:
			dir = RIGHT
		case LEFT:
			dir = UP
		case RIGHT:
			dir = DOWN
		}
		//go to next tile
		next := getNext(tile, dir)
		scanTile(next, dir, state)
	case RS:
		//switch direction
		switch dir {
		case UP:
			dir = RIGHT
		case DOWN:
			dir = LEFT
		case LEFT:
			dir = DOWN
		case RIGHT:
			dir = UP
		}
		//go to next tile
		next := getNext(tile, dir)
		scanTile(next, dir, state)
	case HS:
		if dir == LEFT || dir == RIGHT {
			//continue in the same direction
			next := getNext(tile, dir)
			scanTile(next, dir, state)
		} else {
			//split horizontaly
			d1 := LEFT
			d2 := RIGHT

			scanTile(getNext(tile, d1), d1, state)
			scanTile(getNext(tile, d2), d2, state)
		}
	case VS:
		if dir == UP || dir == DOWN {
			//continue in the same direction
			next := getNext(tile, dir)
			scanTile(next, dir, state)
		} else {
			//split horizontaly
			d1 := UP
			d2 := DOWN

			scanTile(getNext(tile, d1), d1, state)
			scanTile(getNext(tile, d2), d2, state)
		}
	}
}

func getNext(tile []int, dir string) []int {
	i := tile[0]
	j := tile[1]
	switch dir {
	case UP:
		i--
	case DOWN:
		i++
	case LEFT:
		j--
	case RIGHT:
		j++
	}
	return []int{i, j}
}

func unique(s [][]int) [][]int {
	slices.SortFunc(s, func(e1 []int, e2 []int) int {
		if e1[0] < e2[0] {
			return -1
		}
		if e1[0] == e2[0] && e1[1] < e2[1] {
			return -1
		}
		if e1[0] == e2[0] && e1[1] == e2[1] {
			return 0
		}
		return 1
	})

	s = slices.CompactFunc(energized, func(e1 []int, e2 []int) bool {
		return e1[0] == e2[0] && e1[1] == e2[1]
	})
	return s
}
