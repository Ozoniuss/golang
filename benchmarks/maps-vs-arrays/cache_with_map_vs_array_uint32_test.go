package mapsvsarrays

import (
	"math"
	"math/rand/v2"
	"testing"
)

var v = make([]byte, math.MaxUint32)

// Maps are slower when preallocated apparently, when we're talking about this
// size.
var m = make(map[uint32]struct{}, math.MaxUint32)

func withArray() int {

	size := 0

	for i := 0; i < 100000; i++ {
		r := rand.Uint32()
		if v[r] == 0 {
			v[r] = 1
			size++
		}
	}
	return size
}
func WithMap() int {
	for i := 0; i < 100000; i++ {
		r := rand.Uint32()
		m[r] = struct{}{}
	}
	return len(m)
}

var ret int

func BenchmarkWithArray(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		r += withArray() % 761
	}
	ret = r
}

func BenchmarkWithMap(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		r += WithMap() % 761
	}
	ret = r
}
