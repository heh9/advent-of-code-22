package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Move int

const (
	Rock Move = iota
	Paper
	Scissors
)

func PlayRound(p1, p2 Move) int {
	s := p2 + 1

	switch {
	case p1 == p2:
		s += 3
	case (p2+1)%3 != p1:
		s += 6
	default:
	}

	return int(s)
}

func StringToMove(s string) Move {
	d := byte(s[0])
	if d > 67 {
		d -= 23
	}

	return Move(d - 65)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file, err: %v", err)
	}

	t := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		pp := strings.Fields(s.Text())
		co, us := StringToMove(pp[0]), StringToMove(pp[1])
		t += PlayRound(co, us)
	}

	fmt.Printf("Part one: %d\n", t)
}
