package main

import "testing"

//var result int

func BenchmarkDay14part1(b *testing.B) {
	//var r int
	for i := 0; i < b.N; i++ {
		Day14part1()
	}
	//result = r
}
