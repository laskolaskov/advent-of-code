package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day5part1() int {
	result := -1

	seeds, keys, maps := ParseSeedMaps("input-day5.txt")

	for _, seed := range seeds {
		location := getSeedLocation(seed, keys, maps)
		if result == -1 || result > location {
			result = location
		}
	}

	return result // 535088217
}

func Day5part2() int {

	result := -1
	seeds, keys, maps := ParseSeedMaps("input-day5.txt")

	for i := 0; i < len(seeds)-1; i += 2 {
		startSeed := seeds[i]
		endSeed := seeds[i] + seeds[i+1]
		for seed := startSeed; seed <= endSeed; seed++ {
			location := getSeedLocation(seed, keys, maps)
			if result == -1 || result > location {
				result = location
			}
		}
	}

	return result // 51399228 so slow, needs optimizations !!! TODO
}

func ParseSeedMaps(fileName string) ([]int, []string, map[string][]AlmanachEntry) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seeds := []int{}
	keys := []string{}
	maps := map[string][]AlmanachEntry{}
	var currentKey string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		//skip empty lines
		if line == "" {
			continue
		}

		//extract seeds
		if s, found := strings.CutPrefix(line, "seeds: "); found {
			seeds = MapSlice(strings.Split(s, " "), func(el string) int {
				n, err := strconv.Atoi(el)
				if err != nil {
					log.Fatal(err)
				}
				return n
			})
			continue
		}

		//check if line is header and set it as current key
		if s, found := strings.CutSuffix(line, " map:"); found {
			currentKey = s
			keys = append(keys, s)
			continue
		}

		//extract data from current line and push it in the corresponding map
		spl := strings.Split(line, " ")
		source, err := strconv.Atoi(spl[1])
		if err != nil {
			log.Fatal(err)
		}
		dest, err := strconv.Atoi(spl[0])
		if err != nil {
			log.Fatal(err)
		}
		delta, err := strconv.Atoi(spl[2])
		if err != nil {
			log.Fatal(err)
		}
		maps[currentKey] = append(maps[currentKey], AlmanachEntry{
			source: source,
			dest:   dest,
			delta:  delta,
		})
	}

	return seeds, keys, maps
}

func MapSlice[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

type AlmanachEntry struct {
	source int
	dest   int
	delta  int
}

func getSeedLocation(seed int, keys []string, maps map[string][]AlmanachEntry) int {

	current := seed

	//for each key
	for _, key := range keys {
		//iterate the corresponding map to find mapping
		for _, entry := range maps[key] {
			//check if the current value is in the entry source range
			if entry.source <= current && current <= entry.source+entry.delta {
				//change it to the corresponding destination value
				current = current + (entry.dest - entry.source)
				break
			}
		}
	}

	return current
}
