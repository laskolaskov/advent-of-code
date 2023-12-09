package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

var handPower = map[string]int{
	"five of kind":  7,
	"four of kind":  6,
	"full house":    5,
	"three of kind": 4,
	"two pair":      3,
	"one pair":      2,
	"high card":     1,
}

var cardPower = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'1': 0,
}

var cardPowerJoker = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
	'1': 0,
}

func day7part1() int {
	hands := parseCardHands("input-day7.txt", false)

	//custom sorting of the hands
	sort.Slice(hands, func(i, j int) bool {
		if handPower[hands[i].result] < handPower[hands[j].result] {
			return true
		} else if handPower[hands[i].result] > handPower[hands[j].result] {
			return false
		} else {
			// if equal, sort by first card power and so on
			for c := 0; c < 5; c++ {
				if cardPower[rune(hands[i].old[c])] < cardPower[rune(hands[j].old[c])] {
					return true
				} else if cardPower[rune(hands[i].old[c])] > cardPower[rune(hands[j].old[c])] {
					return false
				} else {
					//if equal, continue with comparing next char
					continue
				}
			}
		}
		return false //should never return here on valid input
	})

	//calculate result
	result := 0
	for i, h := range hands {
		current := h.bid * (i + 1)
		result += current
	}
	return result // 252295678
}

func day7part2() int {
	hands := parseCardHands("input-day7.txt", true)

	//custom sorting of the hands
	sort.Slice(hands, func(i, j int) bool {
		if handPower[hands[i].result] < handPower[hands[j].result] {
			return true
		} else if handPower[hands[i].result] > handPower[hands[j].result] {
			return false
		} else {
			// if equal, sort by first card power and so on
			for c := 0; c < 5; c++ {
				if cardPowerJoker[rune(hands[i].old[c])] < cardPowerJoker[rune(hands[j].old[c])] {
					return true
				} else if cardPowerJoker[rune(hands[i].old[c])] > cardPowerJoker[rune(hands[j].old[c])] {
					return false
				} else {
					//if equal, continue with comparing next char
					continue
				}
			}
		}
		return false //should never return here on valid input
	})

	//calculate result
	result := 0
	for i, h := range hands {
		current := h.bid * (i + 1)
		result += current
	}

	return result //  250577259
}

type Hand struct {
	val    string
	old    string
	bid    int
	result string
}

func counter(s string, filtered bool) map[string]int {
	r := map[string]int{}
	f := map[string]int{}
	for _, c := range s {
		r[string(c)]++
	}
	if !filtered {
		return r
	}
	for c, v := range r {
		if v != 1 {
			f[c] = v
		}
	}
	return f
}

func parseCardHands(fileName string, withJoker bool) []Hand {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hands := []Hand{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		spl := strings.Split(line, " ")
		v := spl[0]
		oldVal := v
		b, err := strconv.Atoi(spl[1])
		if err != nil {
			log.Fatal(err)
		}
		if withJoker {
			v = replaceHighesCardWithJoker(v)
		}
		hand := Hand{
			val:    v,
			old:    oldVal,
			bid:    b,
			result: checkHandResult(v),
		}
		hands = append(hands, hand)
	}

	return hands
}

func checkHandResult(v string) string {
	charCount := maps.Values(counter(v, true))
	matches := len(charCount)
	if matches == 0 {
		return "high card"
	}
	if matches == 1 && charCount[0] == 5 {
		return "five of kind"
	}
	if matches == 1 && charCount[0] == 4 {
		return "four of kind"
	}
	if matches == 1 && charCount[0] == 3 {
		return "three of kind"
	}
	if matches == 1 && charCount[0] == 2 {
		return "one pair"
	}
	if matches == 2 && charCount[0] == charCount[1] {
		return "two pair"
	}
	if matches == 2 && charCount[0] != charCount[1] {
		return "full house"
	}
	fmt.Println(charCount)
	return "----------------------------" // should never return this on valid input
}

func replaceHighesCardWithJoker(cardValue string) string {
	charCount := counter(cardValue, false)
	if charCount["J"] == 0 {
		return cardValue
	}
	highestCount := 0
	highestChar := "J"
	for char, count := range charCount {
		if char == "J" {
			continue
		}
		if count > highestCount || (count == highestCount && cardPowerJoker[rune(char[0])] > cardPowerJoker[rune(highestChar[0])]) {
			highestCount = count
			highestChar = char
		}
	}
	//replace jokers with the highest
	return strings.ReplaceAll(cardValue, "J", highestChar)
}
