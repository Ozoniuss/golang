package main

import (
	"fmt"
	"strconv"
	"strings"
)

var predictableInts = []string{
	"0 154 19",
	"0 161 63",
	"0 306 50",
	"0 437 53",
	"0 670 52",
	"0 886 7",
	"0 908 52",
	"1 5 9",
	"1 268 64",
	"1 581 58",
	"2 96 24",
	"2 206 4",
	"2 681 40",
	"2 687 66",
	"2 794 10",
	"2 900 20",
	"3 64 95",
	"3 298 10",
	"3 537 98",
	"3 706 38",
	"3 716 23",
	"4 183 35",
	"4 783 20",
	"4 923 64",
	"5 487 27",
	"5 649 97",
	"6 43 40",
	"6 137 96",
	"6 547 98",
	"6 606 89",
	"7 601 48",
	"7 888 17",
	"8 853 52",
	"8 996 97",
	"9 609 69",
	"9 923 3",
	"10 393 3",
	"10 444 21",
	"10 682 12",
	"11 522 66",
	"11 886 78",
	"12 11 35",
	"12 800 69",
	"13 93 18",
	"13 144 100",
	"13 223 79",
	"13 447 37",
	"13 628 95",
	"14 128 89",
	"14 154 35",
	"14 176 32",
	"14 425 11",
	"14 748 70",
	"15 240 52",
	"15 334 75",
	"15 609 96",
	"15 939 48",
	"16 242 15",
	"16 526 92",
	"16 802 46",
	"16 845 3",
	"16 891 39",
	"17 621 62",
	"17 650 54",
	"17 717 35",
	"17 801 31",
	"17 851 57",
	"18 17 38",
	"18 102 51",
	"18 343 68",
	"18 435 4",
	"19 99 29",
	"19 246 2",
	"19 404 73",
	"19 594 46",
	"20 53 2",
	"20 268 51",
	"20 413 83",
	"20 434 96",
	"20 635 39",
	"21 116 17",
	"21 383 27",
	"21 969 50",
	"22 73 86",
	"22 130 78",
	"22 164 24",
	"22 870 81",
	"23 424 26",
	"24 167 38",
	"24 650 73",
	"25 55 72",
	"25 100 40",
	"25 128 11",
	"25 476 33",
	"25 590 94",
	"26 300 46",
	"26 397 63",
	"26 455 50",
}

func parseInt(line string) (int, int, int) {
	sep1 := 0
	sep2 := 0
	var i1, i2, i3 int

	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			if sep1 == 0 {
				sep1 = i
			} else {
				sep2 = i
				break
			}
		}
	}

	i1, _ = strconv.Atoi(line[0:sep1])
	i2, _ = strconv.Atoi(line[sep1+1 : sep2])
	i3, _ = strconv.Atoi(line[sep2+1 : len(line)])
	return i1, i2, i3
}

func parseSplitInt(line string) (int, int, int) {
	var i1, i2, i3 int
	parts := strings.Split(line, " ")
	i1, _ = strconv.Atoi(parts[0])
	i2, _ = strconv.Atoi(parts[1])
	i3, _ = strconv.Atoi(parts[2])
	return i1, i2, i3
}

func parseSscanfInt(line string) (int, int, int) {
	var i1, i2, i3 int
	fmt.Sscanf(line, "%d %d %d", &i1, &i2, &i3)
	return i1, i2, i3
}

func parseString(line string) (string, string, string) {
	sep1 := 0
	sep2 := 0
	var i1, i2, i3 string

	for i := 0; i < len(line); i++ {
		if line[i] == ' ' {
			if sep1 == 0 {
				sep1 = i
			} else {
				sep2 = i
				break
			}
		}
	}

	i1 = line[0:sep1]
	i2 = line[sep1+1 : sep2]
	i3 = line[sep2+1 : len(line)]
	return i1, i2, i3
}

func parseSplitString(line string) (string, string, string) {
	var i1, i2, i3 string
	parts := strings.Split(line, " ")
	i1 = parts[0]
	i2 = parts[1]
	i3 = parts[2]
	return i1, i2, i3
}

func parseSscanfString(line string) (string, string, string) {
	var i1, i2, i3 string
	fmt.Sscanf(line, "%s %s %s", &i1, &i2, &i3)
	return i1, i2, i3
}

func main() {
	for _, line := range predictableInts {
		fmt.Println(parseInt(line))
		fmt.Println(parseSplitInt(line))
		fmt.Println(parseSscanfInt(line))
	}
}
