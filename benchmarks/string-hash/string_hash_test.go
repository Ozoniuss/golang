package stringhash

import (
	"math/rand/v2"
	"strings"
	"testing"
	"unsafe"
)

// Here I'm essentially testing ways of converting strings to ints. That is
// because sometimes graph nodes are strings, and storing strings or byte
// arrays is more expensive than storing ints. Any 8-character UTF-8 string
// fits into 64 bits, which is just an int in golang.

var sstrings = genStrings(1000)

func genStrings(cap int) []string {
	if cap <= 0 {
		cap = 1000000
	}
	sb := &strings.Builder{}
	strs := make([]string, 0, cap)
	for range cap {
		length := rand.UintN(8)
		for range length {
			rb := byte('a' + rand.UintN(26))
			sb.WriteByte(rb)
		}
		strs = append(strs, sb.String())
		sb.Reset()
	}
	return strs
}

// note that "" is considered the same as the string with character 0
func hash8(s string) uint {
	l := len(s)
	if l == 0 {
		return 0
	}
	r := uint(0)
	for i := 0; i < 8; i++ {
		if i != 0 {
			r = r << 8
		}

		if i < l {
			r = r | uint(s[i])
		}
	}
	return r
}

func hash8unsafe(s string) uint {
	l := len(s)
	if l == 0 {
		return 0
	}
	sdata := unsafe.StringData(s)
	r := uint(0)

	b1 := unsafe.Pointer(uintptr(unsafe.Pointer(sdata)) + 0)
	b2 := unsafe.Pointer(uintptr(unsafe.Pointer(sdata)) + 1)
	b3 := unsafe.Pointer(uintptr(unsafe.Pointer(sdata)) + 2)
	b4 := unsafe.Pointer(uintptr(unsafe.Pointer(sdata)) + 3)
	b5 := unsafe.Pointer(uintptr(unsafe.Pointer(sdata)) + 4)
	b6 := unsafe.Pointer(uintptr(unsafe.Pointer(sdata)) + 5)
	b7 := unsafe.Pointer(uintptr(unsafe.Pointer(sdata)) + 6)
	b8 := unsafe.Pointer(uintptr(unsafe.Pointer(sdata)) + 7)

	r = r | uint(*(*byte)(b1))
	r = r << 8
	r = r | uint(*(*byte)(b2))
	r = r << 8
	r = r | uint(*(*byte)(b3))
	r = r << 8
	r = r | uint(*(*byte)(b4))
	r = r << 8
	r = r | uint(*(*byte)(b5))
	r = r << 8
	r = r | uint(*(*byte)(b6))
	r = r << 8
	r = r | uint(*(*byte)(b7))
	r = r << 8
	r = r | uint(*(*byte)(b8))

	return r
}

// uhash8be unhashes the string assuming big endian representation of the hash.
func uhash8be(hash uint) string {
	c1 := byte(hash >> 56)
	c2 := byte(hash >> 48)
	c3 := byte(hash >> 40)
	c4 := byte(hash >> 32)
	c5 := byte(hash >> 24)
	c6 := byte(hash >> 16)
	c7 := byte(hash >> 8)
	c8 := byte(hash)
	li := 7
	chrs := []byte{c1, c2, c3, c4, c5, c6, c7, c8}
	for li >= 0 {
		if chrs[li] != 0 {
			break
		}
		li--
	}
	if li == -1 {
		return ""
	}
	return string(chrs[:li+1])
}

// uhash8le unhashes the string assuming little endian representation of the hash.
func uhash8le(hash uint) string {
	c1 := byte(hash >> 56)
	c2 := byte(hash >> 48)
	c3 := byte(hash >> 40)
	c4 := byte(hash >> 32)
	c5 := byte(hash >> 24)
	c6 := byte(hash >> 16)
	c7 := byte(hash >> 8)
	c8 := byte(hash)
	li := 7
	chrs := []byte{c8, c7, c6, c5, c4, c3, c2, c1}
	for li >= 0 {
		if chrs[li] != 0 {
			break
		}
		li--
	}
	if li == -1 {
		return ""
	}
	return string(chrs[:li+1])
}

// Sanity checks

func TestHashLittleEndian(t *testing.T) {
	t.Run("hash8unsafev2", func(t *testing.T) {
		for _, s := range sstrings {
			if r := uhash8le(hash8unsafev2(s)); s != r {
				t.Fatalf("expected %s, got %s\n", s, r)
			}
		}
	})
}

func TestHashBigEndian(t *testing.T) {
	t.Run("hash8unsafev2", func(t *testing.T) {
		for _, s := range sstrings {
			if r := uhash8be(hash8(s)); s != r {
				t.Fatalf("expected %s, got %s\n", s, r)
			}
		}
	})
}

func hash8unsafev2(s string) uint {
	l := len(s)
	if l == 0 {
		return 0
	}
	sdata := unsafe.StringData(s)
	// There never seems to be any non-zero data after the string. So maybe
	// a workaround is not necessary (a workaround is easy by doing a bitwise
	// and).
	return *(*uint)(unsafe.Pointer(uintptr(unsafe.Pointer(sdata))))

}

var mi = make(map[uint]struct{})

var ret = 0

func BenchmarkAddToIntMapUnsafev2(b *testing.B) {
	mi := make(map[uint]struct{})
	r := 0
	for range 500 {

		for n := 0; n < b.N; n++ {
			for _, s := range sstrings {
				h := hash8unsafev2(s)
				if _, ok := mi[h]; ok {
					r += 16
				} else {
					mi[h] = struct{}{}
					r += 7
				}
				r = r % 9971
			}
			for k := range mi {
				r += len(uhash8be(k))
				r = r % 47
			}
		}
	}
	ret = r
}

func BenchmarkAddToIntMap(b *testing.B) {
	mi := make(map[uint]struct{})

	r := 0
	for range 500 {

		for n := 0; n < b.N; n++ {
			for _, s := range sstrings {
				h := hash8(s)
				if _, ok := mi[h]; ok {
					r += 16
				} else {
					mi[h] = struct{}{}
					r += 7
				}
				r = r % 9971
			}
			for k := range mi {
				r += len(uhash8be(k))
				r = r % 47
			}
		}
	}
	ret = r
}

func BenchmarkAddToIntMapUnsafe(b *testing.B) {
	mi := make(map[uint]struct{})

	r := 0
	for range 500 {
		for n := 0; n < b.N; n++ {
			for _, s := range sstrings {
				h := hash8unsafe(s)
				if _, ok := mi[h]; ok {
					r += 16
				} else {
					mi[h] = struct{}{}
					r += 7
				}
				r = r % 9971
			}
			for k := range mi {
				r += len(uhash8be(k))
				r = r % 47
			}
		}
	}
	ret = r
}

func BenchmarkAddToStringMap(b *testing.B) {
	ms := make(map[string]struct{})
	r := 0
	for range 500 {
		for n := 0; n < b.N; n++ {
			for _, s := range sstrings {
				if _, ok := ms[s]; ok {
					r += 16
				} else {
					ms[s] = struct{}{}
					r += 7
				}
				r = r % 9971
			}
			for k := range ms {
				r += len(k)
				r = r % 47
			}
		}
	}
	ret = r
}
