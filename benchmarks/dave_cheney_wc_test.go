package main

import "testing"

// go test -run=XXX -bench=. -benchmem wc_dave_cheney*

func BenchmarkApproach1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		approach1(false)
	}
}
func BenchmarkApproach2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		approach2(false)
	}
}
func BenchmarkApproach4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		approach4(false)
	}
}
