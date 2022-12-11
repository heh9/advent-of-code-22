package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file, err: %v", err)
	}

	lmax := 0
	max := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == "" {
			max = Max(lmax, max)
			lmax = 0
			continue
		}

		v, _ := strconv.Atoi(s.Text())
		lmax += v
	}

	fmt.Println(max)
}
