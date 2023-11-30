package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/pkg/profile"
)

// Only run 1 profile at a time! It is known that profiles perform worse if
// doing multiple profiles simultaneously.

// time go run wc_dave_cheney.go assets/moby.txt
// go tool pprof -http=:8080 wcprofile/cpu.pprof
// profiler rate may be changed to get different results

// Profiler samples the callstacks of the active goroutines at each iteration.

var PATH = "assets/moby.txt"

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func approach1(shouldProfile bool) {
	if shouldProfile {
		defer profile.Start(profile.ProfilePath("./wcprofile/"), profile.CPUProfile).Stop()
	}

	f, err := os.Open(PATH)
	if err != nil {
		log.Fatalf("could not open file %q: %v", PATH, err)
	}

	words := 0
	inword := false
	for {
		// In go, operations are not buffered by default, so this calls an
		// operating system call every time, which is slow. For security
		// reasons os operations are likely to only get slower.
		r, err := readbyte(f)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", PATH, err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	// fmt.Printf("%q: %d words\n", PATH, words)
}

// buffered
func approach2(shouldProfile bool) {
	if shouldProfile {
		defer profile.Start(profile.ProfilePath("./wcprofile/"), profile.CPUProfile).Stop()
	}

	f, err := os.Open(PATH)
	if err != nil {
		log.Fatalf("could not open file %q: %v", PATH, err)
	}

	words := 0
	inword := false
	b := bufio.NewReader(f)
	for {
		// In go, operations are not buffered by default, so this calls an
		// operating system call every time, which is slow. For security
		// reasons os operations are likely to only get slower.
		r, err := readbyte(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", PATH, err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	// fmt.Printf("%q: %d words\n", PATH, words)
}

// buffered, memprofile
func approach3(shouldProfile bool) {
	if shouldProfile {
		defer profile.Start(profile.ProfilePath("./wcprofile/"), profile.MemProfile, profile.MemProfileRate(1)).Stop()
	}

	f, err := os.Open(PATH)
	if err != nil {
		log.Fatalf("could not open file %q: %v", PATH, err)
	}

	words := 0
	inword := false
	b := bufio.NewReader(f)
	for {
		// In go, operations are not buffered by default, so this calls an
		// operating system call every time, which is slow. For security
		// reasons os operations are likely to only get slower.
		r, err := readbyte(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", PATH, err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	// fmt.Printf("%q: %d words\n", PATH, words)
}

var extbuf [1]byte

func readbyteExternalBuf(r io.Reader) (rune, error) {
	// The issue with this being a generic reader is that the compiler doesn't
	// know what's going to happen with this buffer. It's possible to capture
	// an address from the buffer and keep it. This forces the allocation to
	// be done on the heap.
	//
	// Note that I'm talking about the slice being created.
	_, err := r.Read(extbuf[:])
	return rune(extbuf[0]), err
}

// buffered, memprofile, noallocs
func approach4(shouldProfile bool) {
	if shouldProfile {
		defer profile.Start(profile.ProfilePath("./wcprofile/"), profile.MemProfile, profile.MemProfileRate(1)).Stop()
	}

	f, err := os.Open(PATH)
	if err != nil {
		log.Fatalf("could not open file %q: %v", PATH, err)
	}

	words := 0
	inword := false
	b := bufio.NewReader(f)
	for {
		// In go, operations are not buffered by default, so this calls an
		// operating system call every time, which is slow. For security
		// reasons os operations are likely to only get slower.
		r, err := readbyteExternalBuf(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", PATH, err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	// fmt.Printf("%q: %d words\n", PATH, words)
}

func main() {
	// approach1(true)
	// approach2(true)
	// approach3(true)
	approach4(true)
}
