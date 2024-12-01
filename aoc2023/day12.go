package aoc2023

import (
	"fmt"
	"strconv"
	"strings"
)

var dp [104][104][104]int

type Entry struct {
	Line   string
	Groups []int
}

func Day12part1() int {
	total := 0
	var entries = []Entry{}
	var lines = scanFile("./aoc2023/input-day12.txt")
	//split lines into entries with text part(string) and array of group lengths with integer values
	for _, l := range lines {
		s := strings.Split(l, " ")
		entries = append(entries, Entry{Line: s[0], Groups: MapSlice(strings.Split(s[1], ","), func(el string) int {
			i, _ := strconv.Atoi(el)
			return i
		})})
	}
	//fmt.Println(entries)
	for _, e := range entries {
		dp = [104][104][104]int{}
		result := recurse(e.Line, 0, e.Groups, 0, 0)
		fmt.Println(e)
		fmt.Println(result)
		total += result
	}
	return total
}

// https://github.com/vipul0092/advent-of-code-2023/blob/main/day12/day12.go#L57
func recurse(pattern string, pidx int, numbers []int, nidx int, grouplen int) int {
	if len(pattern) == pidx {
		if (nidx == len(numbers)-1 && numbers[nidx] == grouplen) || (nidx == len(numbers) && grouplen == 0) {
			return 1
		}
		return 0
	}

	if dp[pidx][nidx][grouplen] != 0 {
		return dp[pidx][nidx][grouplen] - 1
	}
	sum := 0
	char := pattern[pidx]

	if char == '?' || char == '#' {
		// place a '#' and increment the grouplen
		sum += recurse(pattern, pidx+1, numbers, nidx, grouplen+1)
	}
	if char == '?' || char == '.' {
		// if grouplen > 0, we can place a '.' and close the group if the count matches
		if grouplen > 0 && nidx < len(numbers) && numbers[nidx] == grouplen {
			sum += recurse(pattern, pidx+1, numbers, nidx+1, 0)
		}
		// if no group, place a '.' and simply move ahead without any matching
		if grouplen == 0 {
			sum += recurse(pattern, pidx+1, numbers, nidx, 0)
		}
	}

	dp[pidx][nidx][grouplen] = sum + 1
	return sum
}
