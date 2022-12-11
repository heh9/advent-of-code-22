package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

func NewInterval(s string) Interval {
	pp := strings.Split(s, "-")
	x, _ := strconv.Atoi(pp[0])
	y, _ := strconv.Atoi(pp[1])

	return Interval{
		Start: x,
		End:   y,
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Min(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func (a Interval) Contains(b Interval) bool {
	if a.Start <= b.Start && b.End <= a.End {
		return true
	}

	return false
}

func (a Interval) Overlaps(b Interval) bool {
	if a.Start <= b.End && b.Start <= a.End {
		return true
	}

	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file, err: %v", err)
	}

	fc := 0
	o := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		pp := strings.Split(s.Text(), ",")
		a, b := NewInterval(pp[0]), NewInterval(pp[1])

		if a.Contains(b) || b.Contains(a) {
			fc++
		}

		if a.Overlaps(b) {
			o++
		}
	}

	fmt.Printf("Part one: %d\n", fc)
	fmt.Printf("Part two: %d\n", o)
}
