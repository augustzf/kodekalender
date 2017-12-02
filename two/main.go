package main

import (
	"fmt"
)

type row []bool
type matrix []row
type generator func(x, y int) bool
type cell struct {
	x, y                  int
	left, right, up, down *cell
}

func main() {
	labyrinth := newMatrix(1000, formula)
	visited := newMatrix(1000, formula)
	reachable(labyrinth, visited, 1, 1)
	fmt.Println(visited.unreachable())
}

func newMatrix(side int, fn generator) matrix {
	var m = make(matrix, side)
	for y := range m {
		m[y] = make(row, side)
		if fn != nil {
			for x := range m[y] {
				m[y][x] = fn(x+1, y+1)
			}
		}
	}
	return m
}

func formula(x, y int) bool {
	val := x*x*x + 12*x*y + 5*x*y*y
	ones := 0
	for val > 0 {
		val &= val - 1
		ones++
	}
	return ones%2 == 1
}

func reachable(m matrix, visited matrix, a, b int) *cell {

	if a < 1 || a > len(m) || b < 1 || b > len(m) {
		// a,b not reachable: outside of labyrinth
		return nil
	}
	if m.get(a, b) {
		// a, b not reachable: it's a wall
		return nil
	}
	if visited.get(a, b) {
		// cell already reached
		return nil
	}
	c := &cell{
		x: a,
		y: b,
	}
	visited.set(a, b)
	c.left = reachable(m, visited, a-1, b)
	c.right = reachable(m, visited, a+1, b)
	c.up = reachable(m, visited, a, b-1)
	c.down = reachable(m, visited, a, b+1)

	return c
}

func (m matrix) unreachable() int {
	u := 0
	for y := range m {
		r := m[y]
		for x := range r {
			if !r[x] {
				u++
			}
		}
	}
	return u
}

func (m matrix) print() {
	for y := range m {
		r := m[y]
		for x := range r {
			wall := r[x]
			if wall {
				fmt.Print("# ")
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
	}
}

// 1-indexed access
func (m matrix) get(a, b int) bool {
	return m[a-1][b-1]
}

func (m matrix) set(a, b int) {
	m[a-1][b-1] = true
}
