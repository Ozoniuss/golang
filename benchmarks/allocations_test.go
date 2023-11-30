package main

import "testing"

// go test -run=XXX -bench=. -benchmem allocations_test.go

//go:noinline
func ReturnInt() int {
	return 3
}

//go:noinline
func ReturnIntPtr() *int {
	// Return the address of something that is allocated on the stack
	var i int
	i = 3
	return &i
}

//go:noinline
func ReturnIntNew() int {
	// allocate memory
	var i = new(int)
	*i = 3
	return *i
}

var ret1 int = 40

// All these get inlined, must add go:noinline to avoid.
func BenchmarkReturnInt(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		o := ReturnInt()
		r = o
	}
	ret1 = r
}

func BenchmarkReturnIntPtr(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		o := ReturnIntPtr()
		r = *o
	}
	ret1 = r
}

func BenchmarkReturnIntNew(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		o := ReturnIntNew()
		r = o
	}
	ret1 = r
}
