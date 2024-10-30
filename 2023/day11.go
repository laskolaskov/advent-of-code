package main

import (
	"math"
)

func Day11part1and2() int {
	// read file and extract universe matrix (strings are slices of runes)
	var universe = scanFile("input-day11.txt")

	galaxies := [][]int{}
	expRows := []int{}
	expCols := []int{}
	total := 0

	//find "expanded" lines
	for i, l := range universe {
		hasGalaxy := false
		for j, r := range l {
			if r == '#' {
				hasGalaxy = true
				galaxies = append(galaxies, []int{i, j})
			}
		}
		if !hasGalaxy {
			expRows = append(expRows, i)
		}
	}
	//find "expanded" columns
	for j := 0; j < len(universe[0]); j++ {
		hasGalaxy := false
		for i := 0; i < len(universe); i++ {
			if universe[i][j] == '#' {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			expCols = append(expCols, j)
		}
	}

	//find distances between galaxy pairs
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			//distance is the difference between x and y coords of the galaxies
			d := math.Abs(float64((galaxies[j][0] - galaxies[i][0]))) + math.Abs(float64((galaxies[j][1] - galaxies[i][1])))
			//account for expanded rows, if any of them is between the galaxy coords
			//for part 1 each empty row/col is doubled
			//for part 2 each row/col is expanded 1 mil times (1000000)
			//the actual row is added above where d is calculated, hence the -1
			for _, r := range expRows {
				if math.Min(float64(galaxies[i][0]), float64(galaxies[j][0])) < float64(r) && float64(r) < math.Max(float64(galaxies[i][0]), float64(galaxies[j][0])) {
					//d++
					d += 1000000 - 1
				}
			}
			//account for expanded cols
			for _, c := range expCols {
				if math.Min(float64(galaxies[i][1]), float64(galaxies[j][1])) < float64(c) && float64(c) < math.Max(float64(galaxies[i][1]), float64(galaxies[j][1])) {
					//d++
					d += 1000000 - 1
				}
			}
			//add to total
			total += int(d)
		}
	}

	return total
}
