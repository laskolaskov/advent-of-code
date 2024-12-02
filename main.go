package main

import (
	"fmt"
	"log"
)

// https://adventofcode.com/2023
// https://adventofcode.com/2024
func main() {
	key := "2024-02-1"
	f := AocMap[key]
	if f == nil {
		log.Fatal("No func stored for: " + key)
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
