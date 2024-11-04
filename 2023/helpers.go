package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func scanFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func scanBytes(fileName string) [][]byte {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes := [][]byte{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//TODO: check more about this bug
		//reading bytes from scanner and then appending them, results in strange corrupted data (input-day14test2-bug-with-bytes-slice-append.txt)
		//despite that the line we read is identical(!) to the input, something happens when appending (???)
		/* b := scanner.Bytes()
		fmt.Println(b)
		fmt.Println(string(b))
		bytes = append(bytes, b) (???? seems to mess the data completely - input-day14test2-bug-with-bytes-slice-append.txt) */

		l := scanner.Text()
		bytes = append(bytes, []byte(l))
	}

	return bytes
}

func printBytes(state [][]byte) {
	fmt.Println("----------")
	for _, l := range state {
		fmt.Println(string(l))
	}
}

func deepCopy(sl [][]byte) [][]byte {
	r := [][]byte{}
	for _, l := range sl {
		new := append([]byte{}, l...)
		r = append(r, new)
	}
	return r
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
