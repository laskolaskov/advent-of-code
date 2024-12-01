package aoc2023

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Day8part1() int {
	directions, desertMap, _, _ := ParseDesertMap("./aoc2023/input-day8.txt")

	result := 0

	start := "AAA"
	end := "ZZZ"
	limit := len(desertMap) * 1000 //arbitrary big number to limit the cycle

	currentDir := ""
	currentDirections := directions
	currentPos := start

	for {
		//if no more directions, restart from beginning
		if len(currentDirections) < 1 {
			currentDirections = directions
		}
		//get next direction from the string
		currentDir = string(currentDirections[0])
		//and pop it out
		currentDirections = currentDirections[1:]
		//get next position
		if currentDir == "L" {
			nextPos := desertMap[currentPos].left
			currentPos = nextPos
			result++
		} else if currentDir == "R" {
			nextPos := desertMap[currentPos].right
			currentPos = nextPos
			result++
		} else {
			log.Fatal("Invalid direction: ", currentDir)
		}
		//check exit condition
		if currentPos == end {
			break
		}
		//emergency exit
		if result > limit {
			log.Fatal("Too many steps ???? ", limit)
		}
	}

	return result // 13301
}

func Day8part2() int {
	directions, desertMap, startNodes, _ := ParseDesertMap("./aoc2023/input-day8.txt")
	result := 0

	limit := len(desertMap) * 1000000 //arbitrary big number to limit the cycle

	routes := map[string]map[string]int{}
	pathsLengths := []int{}

	/*
		1. run to see all paths with their ending positions
		2. NOTICE (that seems the hard part ???) that each path has ONLY 1 start and end positions
		(also we can see that there are only 6 start and only 6 end nodes in the map)
		3. then each path must be cycled until all paths reach their end together, which is the LEAST COMMON MULTIPLE of all path lenghts
		4. so add code to extract all lenghts into 'pathsLengths'
	*/

	for _, startNode := range startNodes {

		routes[startNode] = map[string]int{}
		currentDir := ""
		currentDirections := directions
		currentPos := startNode

		steps := 0

		for {
			//if no more directions, restart from beginning
			if len(currentDirections) < 1 {
				currentDirections = directions
			}
			//get next direction from the string
			currentDir = string(currentDirections[0])
			//and pop it out
			currentDirections = currentDirections[1:]
			//get next position
			if currentDir == "L" {
				nextPos := desertMap[currentPos].left
				currentPos = nextPos
				steps++
			} else if currentDir == "R" {
				nextPos := desertMap[currentPos].right
				currentPos = nextPos
				steps++
			} else {
				log.Fatal("Invalid direction: ", currentDir)
			}
			//check exit condition
			if string(currentPos[2]) == "Z" {
				if routes[startNode][currentPos] != 0 {
					break
				}
				routes[startNode][currentPos] = steps
				pathsLengths = append(pathsLengths, steps)
			}
			//emergency exit
			if result > limit {
				log.Fatal("Too many steps ???? ", limit)
			}
		}
	}

	//calculate the LCM of all the path lengths, as established
	result = LCM(pathsLengths[0], pathsLengths[1], pathsLengths[2:]...)
	return result // 7309459565207
}

func ParseDesertMap(fileName string) (string, map[string]MapEntry, []string, []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	directions := ""
	desertMap := map[string]MapEntry{}
	startNodes := []string{}
	endNodes := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		spl := strings.Split(line, " = ")
		if len(spl) == 1 {
			directions = line
			continue
		}
		spl2 := strings.Split(spl[1][1:len(spl[1])-1], ", ")
		mapEntry := MapEntry{
			pos:   spl[0],
			left:  spl2[0],
			right: spl2[1],
		}
		desertMap[mapEntry.pos] = mapEntry
		if string(mapEntry.pos[2]) == "A" {
			startNodes = append(startNodes, mapEntry.pos)
		}
		if string(mapEntry.pos[2]) == "Z" {
			endNodes = append(endNodes, mapEntry.pos)
		}
	}
	return directions, desertMap, startNodes, endNodes
}

type MapEntry struct {
	pos   string
	left  string
	right string
}
