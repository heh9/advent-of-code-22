package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type Stack []byte

type Warehouse []Stack

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack) Pop() (byte, error) {
	b, err := s.Head()
	if err != nil {
		return b, err
	}

	h := len(*s)
	*s = (*s)[:h-1]

	return b, nil
}

func (s Stack) Head() (byte, error) {
	var b byte

	if s.IsEmpty() {
		return b, errors.New("stack is empty")
	}

	return s[len(s)-1], nil
}

func (s *Stack) Push(b byte) {
	*s = append(*s, b)
}

func (s *Stack) Reverse() {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func (w *Warehouse) AddStringToStacks(s string) {
	st := 0

	for {
		if st > len(*w)-1 {
			(*w) = append(*w, Stack{})
		}
		if s[0] == '[' {
			(*w)[st] = append((*w)[st], s[1])
		}
		if s = s[3:]; len(s) == 0 {
			break
		}

		s = s[1:]
		st++
	}
}

func (w *Warehouse) Reverse() {
	for _, s := range *w {
		s.Reverse()
	}
}

func (w *Warehouse) MoveSeq(n, f, t int) error {
	for s := 0; s < n; s++ {
		h, err := (*w)[f].Pop()
		if err != nil {
			return fmt.Errorf("unable to move, err: %v", err)
		}

		(*w)[t].Push(h)
	}

	return nil
}

func (w *Warehouse) Move(n, f, t int) error {
	buf := Stack{}

	for s := 0; s < n; s++ {
		h, err := (*w)[f].Pop()
		if err != nil {
			return fmt.Errorf("unable to move, err: %v", err)
		}

		buf.Push(h)
	}

	for !buf.IsEmpty() {
		h, _ := buf.Pop()
		(*w)[t].Push(h)
	}

	return nil
}

func (w Warehouse) TopLayer() []byte {
	t := []byte{}

	for _, s := range w {
		h, _ := s.Head()
		t = append(t, h)
	}

	return t
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file, err: %v", err)
	}

	var w Warehouse

	s := bufio.NewScanner(f)
	for s.Scan() {
		buf := s.Text()

		if buf == "" {
			break
		}

		w.AddStringToStacks(buf)
	}

	w.Reverse()

	for s.Scan() {
		var n, f, t int

		_, err := fmt.Sscanf(s.Text(), "move %d from %d to %d", &n, &f, &t)
		if err != nil {
			log.Printf("unable to read command %s\n", err)
		}

		w.Move(n, f-1, t-1)
	}

	fmt.Printf("Top containers: %s\n", string(w.TopLayer()))
}
