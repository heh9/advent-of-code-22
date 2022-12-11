package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func TopNSnacks(n int, r io.Reader) []int {
	sum := 0
	var sums []int

	s := bufio.NewScanner(r)
	for s.Scan() {
		if s.Text() == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}

		if v, err := strconv.Atoi(s.Text()); err == nil {
			sum += v
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sums)))

	return sums[:n]
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file, err: %v", err)
	}

	top := TopNSnacks(3, f)

	fmt.Printf("Part one: %d\n", top[0])
	fmt.Printf("Part two: %d\n", top[0]+top[1]+top[2])
}
