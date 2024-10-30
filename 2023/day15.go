package main

import (
	"log"
	"strconv"
	"strings"
)

var EQUAL = byte('=')
var DASH = byte('-')

type Lens struct {
	Label string
	Value int
}

func Day15part1() int {
	h := 0
	input := scanFile("input-day15.txt")
	if len(input) > 1 {
		log.Fatal("the input must be one long line")
	}
	split := strings.Split(input[0], ",")
	for _, s := range split {
		h += hash(s)
	}
	return h
}

func Day15part2() int {
	boxes := make([][]Lens, 256)
	input := scanFile("input-day15.txt")
	if len(input) > 1 {
		log.Fatal("the input must be one long line")
	}

	split := strings.Split(input[0], ",")
	for _, s := range split {
		var label []byte
		var operation byte
		var value int
		for _, c := range s {
			if byte(c) == EQUAL || byte(c) == DASH {
				operation = byte(c)
			} else {
				if operation != 0 {
					//have operation, this is the value
					v, _ := strconv.Atoi(string(c))
					value = v
				} else {
					//it is label
					label = append(label, byte(c))
				}
			}
		}
		doOperation(boxes, string(label), operation, value)
	}
	//fmt.Println(boxes)

	//calculate focusing power
	p := 0
	for boxIndex, box := range boxes {
		for lensIndex, lens := range box {
			p += (boxIndex + 1) * (lensIndex + 1) * lens.Value
		}
	}
	return p
}

func doOperation(boxes [][]Lens, label string, operation byte, value int) {
	boxIndex := hash(label)
	found := false
	for i := 0; i < len(boxes[boxIndex]); i++ {
		if boxes[boxIndex][i].Label == label && operation == DASH {
			//remove this element and stop
			if i == len(boxes[boxIndex])-1 {
				boxes[boxIndex] = boxes[boxIndex][:i]
			} else if i == 0 {
				boxes[boxIndex] = boxes[boxIndex][1:]
			} else {
				boxes[boxIndex] = append(boxes[boxIndex][:i], boxes[boxIndex][i+1:]...)
			}
			found = true
			break
		}
		if boxes[boxIndex][i].Label == label && operation == EQUAL {
			//change the value and stop
			boxes[boxIndex][i].Value = value
			found = true
			break
		}
	}
	if !found && operation == EQUAL {
		//label was not found - append it
		boxes[boxIndex] = append(boxes[boxIndex], Lens{Label: label, Value: value})
	}
}

func hash(s string) int {
	var h int
	for _, c := range s {
		i := int(c)
		h += i
		h *= 17
		h = h % 256
	}
	return h
}
