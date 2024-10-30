package main

import "testing"

//var result int

func BenchmarkDay13part2(b *testing.B) {
	//var r int
	for i := 0; i < b.N; i++ {
		Day13part2()
	}
	//result = r
}
