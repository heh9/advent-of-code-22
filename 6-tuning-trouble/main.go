package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Queue []byte

func (q *Queue) Push(b byte) {
	*q = append(*q, b)
}

func (q *Queue) Pop() bool {
	if len(*q) > 0 {
		*q = (*q)[1:]
		return true
	}

	return false
}

func (q Queue) Contains(b byte) bool {
	for _, v := range q {
		if b == v {
			return true
		}
	}

	return false
}

func IsUnique(s string) bool {
	u := map[rune]bool{}

	for _, v := range s {
		if _, ok := u[v]; ok == true {
			return false
		}

		u[v] = true
	}

	return true
}

func IndexOfMarker(s string, m int) int {
	d := 0
	q := Queue{}

	for _, v := range s {
		d++

		if len(q) == m {
			if IsUnique(string(q)) {
				return d - 1
			}

			q.Pop()
		}

		b := byte(v)
		if q.Contains(b) {
			q.Pop()
		}

		q.Push(b)
	}

	return 0
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file, err: %v", err)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(IndexOfMarker(s.Text(), 14))
	}
}
