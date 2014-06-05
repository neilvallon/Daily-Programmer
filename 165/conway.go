package main

import (
	"fmt"
	"time"
)

func main() {
	w := NewWorld(10, 10)
	w.m = Map{
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, false, false, false},
		{false, true, true, true, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
	}

	go w.Start()
	for i := 0; i <= 700; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(i)
		fmt.Println(w.Next())
	}
}

type Cell bool

func (c Cell) String() string {
	if c {
		return "#"
	}
	return "."
}

type Map [][]Cell

func (m Map) String() (s string) {
	for _, y := range m {
		s += fmt.Sprintf("%s\n", y)
	}
	return
}

func (m Map) NeighborsOf(x, y int) (c int) {
	maxy, maxx := len(m), len(m[0])
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			xx := (((x + j) % maxx) + maxx) % maxx
			yy := (((y + i) % maxy) + maxy) % maxy
			if m[yy][xx] {
				c++
			}
		}
	}

	return
}

func (m Map) next() Map {
	x, y := len(m[0]), len(m)
	nm := make(Map, y)
	for i := 0; i < y; i++ {
		nm[i] = make([]Cell, x)
	}

	for y := range m {
		for x := range m[y] {
			n := m.NeighborsOf(x, y)
			if m[y][x] {
				nm[y][x] = Cell(2 <= n && n <= 3)
			} else {
				nm[y][x] = n == 3
			}
		}
	}

	return nm
}

type World struct {
	m Map
	c chan Map
}

func NewWorld(x, y int) *World {
	m := make(Map, y)
	for i := 0; i < y; i++ {
		m[i] = make([]Cell, x)
	}

	return &World{
		m: m,
		c: make(chan Map),
	}
}

func (w *World) Start() {
	for {
		w.c <- w.m
		w.m = w.m.next()
	}
}

func (w *World) Next() Map {
	return <-w.c
}
