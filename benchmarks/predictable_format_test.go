package main

import "testing"

var ret = 0
var retStr = ""

func BenchmarkParseInt(b *testing.B) {
	var total int
	for n := 0; n < b.N; n++ {
		for _, line := range predictableInts {
			var x1, x2, x3 int
			x1, x2, x3 = parseInt(line)
			total += x1 + x2 + x3
		}
	}
	ret = total
}
func BenchmarkParseSplitInt(b *testing.B) {
	var total int
	for n := 0; n < b.N; n++ {
		for _, line := range predictableInts {
			var x1, x2, x3 int
			x1, x2, x3 = parseSplitInt(line)
			total += x1 + x2 + x3
		}
	}
	ret = total
}

func BenchmarkParseSscanfInt(b *testing.B) {
	var total int
	for n := 0; n < b.N; n++ {
		for _, line := range predictableInts {
			var x1, x2, x3 int
			x1, x2, x3 = parseSscanfInt(line)
			total += x1 + x2 + x3
		}
	}
	ret = total
}
func BenchmarkParseString(b *testing.B) {
	var total string
	for n := 0; n < b.N; n++ {
		for _, line := range predictableInts {
			var x1, x2, x3 string
			x1, x2, x3 = parseString(line)

			if x1 < x2 {
				total = x3
			} else {
				total = x2
			}
		}
	}
	retStr = total
}
func BenchmarkParseSplitString(b *testing.B) {
	var total string
	for n := 0; n < b.N; n++ {
		for _, line := range predictableInts {
			var x1, x2, x3 string
			x1, x2, x3 = parseSplitString(line)

			if x1 < x2 {
				total = x3
			} else {
				total = x2
			}
		}
	}
	retStr = total
}

func BenchmarkParseSscanfString(b *testing.B) {
	var total string
	for n := 0; n < b.N; n++ {
		for _, line := range predictableInts {
			var x1, x2, x3 string
			x1, x2, x3 = parseSscanfString(line)
			if x1 < x2 {
				total = x3
			} else {
				total = x2
			}
		}

	}
	retStr = total
}
