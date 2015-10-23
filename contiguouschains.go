package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	Label int
}

type Grid struct {
	Matrix   [][]*Point
	Label    int
	PrevLoc  []int
	Location []int
}

// FindChain searches the grid for the beginning of a chain. The beginning of
// a chain is a Point that is unlabelled, and has at most one unlabelled
// neighbour. If there are unlabelled points, but none with a single partner
func (g *Grid) FindChain() bool {
	for r := range g.Matrix {
		for c := range g.Matrix[r] {
			if p := g.Matrix[r][c]; p != nil && p.Label == 0 {
				g.Location[0] = r
				g.Location[1] = c
				g.PrevLoc[0] = r
				g.PrevLoc[1] = c
				g.Label++
				p.Label = g.Label
				return true
			}
		}
	}
	return false
}

func (g *Grid) TraceChain() {
	// Assumes that the start of the chain has already been assigned

}

func (g *Grid) String() string {
	message := ""
	for r := range g.Matrix {
		for c := range g.Matrix[r] {
			if g.Matrix[r][c] != nil {
				message += strconv.Itoa(g.Matrix[r][c].Label)
			} else {
				message += " "
			}
		}
		message += "\n"
	}
	return message
}

func CreateGrid(input string) *Grid {
	rows := strings.Split(input, "\n")
	grid := &Grid{
		Matrix:   make([][]*Point, len(rows)),
		Label:    0,
		Location: []int{0, 0},
	}

	for r := range rows {
		columns := strings.Split(rows[r], "")
		grid.Matrix[r] = make([]*Point, len(columns))
		for c := range columns {
			if columns[c] == "x" {
				grid.Matrix[r][c] = &Point{}
			}
		}
	}

	return grid
}

func main() {
	input := `xxxx xxxx
   xxx   
x   x   x
xxxxxxxxx`

	grid := CreateGrid(input)
	fmt.Println(grid)

}
