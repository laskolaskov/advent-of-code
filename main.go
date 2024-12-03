package main

import (
	"fmt"
	"log"
	"os"
)

// https://adventofcode.com/2023
// https://adventofcode.com/2024
func main() {
	if len(os.Args) < 2 {
		log.Fatal("provide date argument: YYYY-DD-part(1 or 2) ; ex.: '2024-01-1' -> day 1 part 1 for 2024 (the month is known)")
	}
	key := os.Args[1]
	f := AocMap[key]
	if f == nil {
		log.Fatal("No func for: " + key)
	}
	if key == "2023-03-2" ||
		key == "2023-18-1" ||
		key == "2023-18-2" {
		result := f.(func() int64)()
		fmt.Printf("Result for %s : %d \n", key, result)
	} else {
		result := f.(func() int)()
		fmt.Printf("Result for %s : %d \n", key, result)
	}

}
