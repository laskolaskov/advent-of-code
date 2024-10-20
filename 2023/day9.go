package main

import (
	"log"
	"strconv"
	"strings"
)

type Seq []int

func (s Seq) isZeroes() bool {
	for _, el := range s {
		if el != 0 {
			return false
		}
	}
	return true
}

func Day9part1() int {
	total := 0
	seqs := make(map[int]Seq)
	steps := make(map[int][]Seq)
	isZeroes := make(map[int]bool)
	nextInSeq := make(map[int]int)

	//read file and extract sequences
	lines := scanFile("input-day9.txt")
	for i, l := range lines {
		sequence := MapSlice(strings.Split(l, " "), func(el string) int {
			n, err := strconv.Atoi(el)
			if err != nil {
				log.Fatal(err)
			}
			return n
		})
		seqs[i] = sequence
	}

	//calculate first interpolation sequence
	for seqIndex, s := range seqs {
		//create empty step sequence
		step := Seq{}
		//calculate interpolation steps
		for i, el := range s {
			//skip calculation for first element
			if i == 0 {
				continue
			}
			//calculate diff between current element and previous element
			diff := el - s[i-1]
			//append to step sequence
			step = append(step, diff)
		}
		//append into steps map where key is the sequence index
		steps[seqIndex] = append(steps[seqIndex], step)
		isZeroes[seqIndex] = step.isZeroes()
	}

	//calculate next step for non-zero previous steps
	//itearate over staps map
	for seqIndex, s := range steps {
		//take the last step sequence
		lastStep := s[len(s)-1]
		//iterate to find next steps, until last step becomes zero sequence
		for {
			//create empty step sequence
			step := Seq{}
			//calculate interpolation steps
			for i, el := range lastStep {
				//skip calculation for first element
				if i == 0 {
					continue
				}
				//calculate diff between current element and previous element
				diff := el - lastStep[i-1]
				//append to step sequence
				step = append(step, diff)
			}
			//append into steps map where key is the sequence index
			steps[seqIndex] = append(steps[seqIndex], step)
			isZeroes[seqIndex] = step.isZeroes()
			//if last step is zero sequence, break
			if step.isZeroes() {
				break
			}
			//last step becomes the current step
			lastStep = step
		}
	}
	//calculate result
	//iterate over the steps map
	for seqIndex, s := range steps {
		result := 0
		//iterate over each step sequence in reverse, skipping the last (the sero sequence)
		for i := len(s) - 2; i >= 0; i-- {
			//add last element of the current sequence to the total result
			result += s[i][len(s[i])-1]
		}
		//add the result to the last element of the original sequence
		result += seqs[seqIndex][len(seqs[seqIndex])-1]
		//add the result in the intermediate map
		nextInSeq[seqIndex] = result
		//add to total result
		total += result
	}

	return total
}

func Day9part2() int {
	total := 0
	seqs := make(map[int]Seq)
	steps := make(map[int][]Seq)
	isZeroes := make(map[int]bool)
	nextInSeq := make(map[int]int)

	//read file and extract sequences
	lines := scanFile("input-day9.txt")
	for i, l := range lines {
		sequence := MapSlice(strings.Split(l, " "), func(el string) int {
			n, err := strconv.Atoi(el)
			if err != nil {
				log.Fatal(err)
			}
			return n
		})
		seqs[i] = sequence
	}

	//calculate first interpolation sequence
	for seqIndex, s := range seqs {
		//create empty step sequence
		step := Seq{}
		//calculate interpolation steps
		for i, el := range s {
			//skip calculation for first element
			if i == 0 {
				continue
			}
			//calculate diff between current element and previous element
			diff := el - s[i-1]
			//append to step sequence
			step = append(step, diff)
		}
		//append into steps map where key is the sequence index
		steps[seqIndex] = append(steps[seqIndex], step)
		isZeroes[seqIndex] = step.isZeroes()
	}

	//calculate next step for non-zero previous steps
	//itearate over staps map
	for seqIndex, s := range steps {
		//take the last step sequence
		lastStep := s[len(s)-1]
		//iterate to find next steps, until last step becomes zero sequence
		for {
			//create empty step sequence
			step := Seq{}
			//calculate interpolation steps
			for i, el := range lastStep {
				//skip calculation for first element
				if i == 0 {
					continue
				}
				//calculate diff between current element and previous element
				diff := el - lastStep[i-1]
				//append to step sequence
				step = append(step, diff)
			}
			//append into steps map where key is the sequence index
			steps[seqIndex] = append(steps[seqIndex], step)
			isZeroes[seqIndex] = step.isZeroes()
			//if last step is zero sequence, break
			if step.isZeroes() {
				break
			}
			//last step becomes the current step
			lastStep = step
		}
	}
	//calculate result
	//same as first part, but we extrapolate new first elements instead (by substracting from original first elements)
	//iterate over the steps map
	for seqIndex, s := range steps {
		result := 0
		//iterate over each step sequence in reverse, skipping the last (the sero sequence)
		for i := len(s) - 2; i >= 0; i-- {
			//substract result from the first element of the current sequence to calculate the new result
			result = s[i][0] - result
		}
		//substract result from the first element of the original sequence to calculate the new result
		result = seqs[seqIndex][0] - result
		//add the result in the intermediate map
		nextInSeq[seqIndex] = result
		//add to total result
		total += result
	}

	return total
}
