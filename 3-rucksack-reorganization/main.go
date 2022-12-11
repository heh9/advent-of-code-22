package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
// ASCII A == 65, a = 97
func CharToPriority(c rune) int {
	d := byte(c)

	if d > 96 {
		return int(d - 96)
	}

	return int(d - 38)
}

func CompareBackpacks(a, b string) int {
	s := map[int]bool{}

	for _, v := range a {
		s[CharToPriority(v)] = true
	}

	for _, v := range b {
		p := CharToPriority(v)
		if _, ok := s[p]; ok != false {
			return p
		}
	}

	return 0
}

func FindBadgeSticker(bb ...string) int {
	f := make([]int, 53)

	for n, bp := range bb {
		for _, c := range bp {
			p := CharToPriority(c)

			if f[p] == n {
				f[p] = n + 1
			}
		}
	}

	for p, v := range f {
		if v == len(bb) {
			return p
		}
	}

	return 0
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file, err: %v", err)
	}

	sum := 0
	badges := 0
	var buf []string

	s := bufio.NewScanner(f)
	for s.Scan() {
		pp := s.Text()
		h := len(pp) / 2

		sum += CompareBackpacks(pp[0:h], pp[h:])

		buf = append(buf, pp)
		if len(buf) == 3 {
			badges += FindBadgeSticker(buf...)
			buf = nil
		}
	}

	fmt.Printf("Part one: %d\n", sum)
	fmt.Printf("Part two: %d\n", badges)
}
