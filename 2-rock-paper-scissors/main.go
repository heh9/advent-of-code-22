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

type Output int

const (
	Lose Output = 0
	Draw        = 3
	Win         = 6
)

func PlayRound(p1, p2 Move) int {
	var o Output

	switch {
	case p1 == p2:
		o = Draw
	case (p2+1)%3 != p1:
		o = Win
	default:
		o = Lose
	}

	return int(o)
}

func NextMove(m Move, o Output) Move {
	nm := Rock

	for PlayRound(m, nm) != int(o) {
		nm++
	}

	return nm
}

func StringToOutput(s string) Output {
	return Output(byte(s[0])-88) * 3
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
	tf := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		pp := strings.Fields(s.Text())
		co, us := StringToMove(pp[0]), StringToMove(pp[1])

		t += PlayRound(co, us) + int(us) + 1

		nm := NextMove(co, StringToOutput(pp[1]))
		tf += PlayRound(co, nm) + int(nm) + 1
	}

	fmt.Printf("Part one: %d\n", t)
	fmt.Printf("Part two: %d\n", tf)
}
