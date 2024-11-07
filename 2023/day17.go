package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"

	"slices"
)

type Tile struct {
	Pos   complex128
	Dir   complex128
	Val   int
	Total int
}

type TileHeap []Tile

func (h TileHeap) Len() int           { return len(h) }
func (h TileHeap) Less(i, j int) bool { return h[i].Total < h[j].Total }
func (h TileHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TileHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Tile))
}

func (h *TileHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// https://www.reddit.com/r/adventofcode/comments/18k9ne5/comment/kdqp7jx/
// TODO: my implementation is too slow :(
func Day17part1() int {
	plain, end := scanComplexPlain("input-day17.txt")

	seen := []Tile{}
	todo := &TileHeap{
		Tile{Dir: complex(0, 1)},
		Tile{Dir: complex(1, 0)},
	}

	//part 1
	/*
		min := 1
		max := 3
	*/
	//part 2
	min := 4
	max := 10
	result := 0

	for todo.Len() > 0 {
		t := heap.Pop(todo).(Tile)

		if t.Pos == end {
			result = t.Total
			break
		}

		if slices.ContainsFunc(seen, func(s Tile) bool {
			return t.Pos == s.Pos && t.Dir == s.Dir
		}) {
			continue
		}
		seen = append(seen, t)

		for _, d := range []complex128{complex(0, 1) / t.Dir, complex(0, -1) / t.Dir} {
			for i := min; i <= max; i++ {
				newPos := t.Pos + d*complex(float64(i), 0)
				//check if end is in the plain
				if val, ok := plain[newPos]; ok {
					//sum the value of each tile passed
					sum := 0
					for j := 1; j <= i; j++ {
						p := t.Pos + d*complex(float64(j), 0)
						sum += plain[p]
					}
					heap.Push(todo, Tile{Total: t.Total + sum, Pos: newPos, Dir: d, Val: val})
				}
			}
		}
	}

	return result
}

// retuns the map of (complex coords)=>value and the coords of the last element (down right)
func scanComplexPlain(fileName string) (map[complex128]int, complex128) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	plain := make(map[complex128]int)

	scanner := bufio.NewScanner(file)

	var i float64 = 0
	var j float64 = 0
	for scanner.Scan() {
		l := scanner.Text()
		j = float64(len(l) - 1)
		for j, c := range l {
			key := complex(i, float64(j))
			val, _ := strconv.Atoi(string(c))
			plain[key] = val
		}
		i++
	}

	return plain, complex(i-1, j)
}

func printComplexPlain(plain map[complex128]int) {
	fmt.Println("----------")
	for k, v := range plain {
		fmt.Println(k, v)
	}
}
