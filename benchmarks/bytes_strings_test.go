package main

import (
	"testing"
)

var sstrings = []string{
	"OAm",
	"6VD",
	"eyb",
	"l5H",
	"WZg",
	"p4V",
	"Jre",
	"oJ7",
	"y61",
	"41E",
	"eW7",
	"0Jc",
	"Zpp",
	"FTN",
	"jLK",
	"6go",
	"sTF",
	"Q5P",
	"xF4",
	"8aM",
	"woU",
	"JQK",
	"kmh",
	"utp",
	"7Pg",
	"dhX",
	"as8",
	"Ig2",
	"qa9",
	"oAX",
	"Gd9",
	"Pj7",
	"KKy",
	"R6s",
	"dlm",
	"2Gq",
	"0eG",
	"HGz",
	"Ky5",
	"LNu",
	"3ZE",
	"Up5",
	"1GL",
	"G7P",
	"FDM",
	"zKu",
	"3BV",
	"FHP",
	"xpO",
	"lJx",
	"7Yd",
	"qoJ",
	"kp4",
	"RVU",
	"jyS",
	"KqQ",
	"ZhK",
	"iy2",
	"mno",
	"SvF",
	"7W7",
	"uMc",
	"FF8",
	"X1R",
	"f8r",
	"lzS",
	"Zo4",
	"0ei",
	"EfV",
	"VRp",
	"eaG",
	"89O",
	"8Ic",
	"GMo",
	"Lrr",
	"I2k",
	"MLv",
	"sat",
	"vU9",
	"fRL",
	"LU3",
	"VxW",
	"Amu",
	"BBG",
	"VBl",
	"9bg",
	"tEQ",
	"rRT",
	"yqW",
	"tV8",
	"PnB",
	"Gib",
	"jQc",
	"5gb",
	"I8k",
	"JjJ",
	"jGx",
	"v95",
	"RVH",
	"pxH",
}

// This tests the speed of conversion between strings and bytes. I'm essentially
// looking for the fastest way of converting a string to an array of bytes.

// toByteArray takes a string and converts it to a byte array of 4 bytes. The
// number 4 was chosen because computers tend to optimize better "round"
// numbers.
//
// go:noinline
func toByteArr(s string) [4]byte {
	var arr [4]byte
	for i := 0; i < len(s); i++ {
		arr[i] = s[i]
	}
	return arr
}

// toByteSlices takes a string and converts it to a byte slice of capacity 4.
//
// go:noinline
func toByteSlice(s string) []byte {
	var arr = make([]byte, 4, 4)
	for i := 0; i < len(s); i++ {
		arr[i] = s[i]
	}
	return arr
}

// toByteSlice2 does the same as toByteSlice, except it uses the build in copy
// method.
func toByteSlice2(s string) []byte {
	var arr = make([]byte, 4, 4)
	copy(arr, s)
	return arr
}

var out [4]byte
var out2 []byte

// Simply convert to byte array and assign to the byte array. I wasn't expecting
// this to be the slowest.
func BenchmarkByteArr(b *testing.B) {
	var in [4]byte
	for n := 0; n < b.N; n++ {
		for _, s := range sstrings {
			in = toByteArr(s)
		}
	}
	out = in
}

// First convert to slice, then assign to array.
func BenchmarkByteSliceThenArr(b *testing.B) {
	var in [4]byte
	for n := 0; n < b.N; n++ {
		for _, s := range sstrings {
			inslc := toByteSlice(s)
			for i := 0; i < 4; i++ {
				in[i] = inslc[i]
			}
		}
	}
	out = in
}

// First convert to slice using copy, then assign to array.
func BenchmarkByteSliceCopy(b *testing.B) {
	var in [4]byte

	for n := 0; n < b.N; n++ {
		for _, s := range sstrings {
			inslc := toByteSlice2(s)
			for i := 0; i < 4; i++ {
				in[i] = inslc[i]
			}
		}
	}
	out = in
}

// Simply convert to byte slice, then assign to in.
func BenchmarkByteSlice2(b *testing.B) {
	var in []byte
	for n := 0; n < b.N; n++ {
		for _, s := range sstrings {
			// Allocates because it's pretty hard to track in. Essentially
			// in will point to some new byte slice, and it's not possible
			// to see what in will do with that slice.
			in = toByteSlice(s)

			// If we copy however we don't allocate because we store everything
			// in the original slice in points to.
			// copy(in, toByteSlice(s))
		}
	}
	out2 = in
}
