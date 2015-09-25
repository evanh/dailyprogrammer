package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Game struct {
	Grid   [][]string
	Width  int
	Height int
}

func (g *Game) Iterate() {
	newgrid := make([][]string, g.Height)

	for r := range g.Grid {
		newgrid[r] = make([]string, g.Width)
		for c := range g.Grid[r] {
			cell := g.Grid[r][c]
			neighbours := g.GetNeighbours(r, c)
			if cell != " " {
				if len(neighbours) <= 1 {
					newgrid[r][c] = " "
				} else if len(neighbours) >= 4 {
					newgrid[r][c] = " "
				} else {
					newgrid[r][c] = cell
				}
			} else {
				if len(neighbours) == 3 {
					index := rand.Intn(3)
					newgrid[r][c] = neighbours[index]
				} else {
					newgrid[r][c] = " "
				}
			}
		}
	}
	g.Grid = newgrid
}

func (g *Game) GetNeighbours(r, c int) []string {
	neighbours := []string{}
	locations := [][]int{
		[]int{r - 1, c - 1},
		[]int{r - 1, c},
		[]int{r - 1, c + 1},
		[]int{r, c - 1},
		[]int{r, c + 1},
		[]int{r + 1, c - 1},
		[]int{r + 1, c},
		[]int{r + 1, c + 1},
	}
	for _, loc := range locations {
		if val := g.GetValue(loc[0], loc[1]); val != " " {
			neighbours = append(neighbours, val)
		}
	}

	return neighbours
}

func (g *Game) GetValue(r, c int) string {
	if r < 0 || r >= g.Height {
		return " "
	} else if c < 0 || c >= g.Width {
		return " "
	} else {
		return g.Grid[r][c]
	}
}

func (g *Game) IsDead() bool {
	for r := range g.Grid {
		for c := range g.Grid[r] {
			if g.GetValue(r, c) != " " {
				return false
			}
		}
	}
	return true
}

func (g *Game) String() string {
	output := ""
	for r := range g.Grid {
		for c := range g.Grid[r] {
			output += g.GetValue(r, c)
		}
		output += "\n"
	}
	return output
}

func MakeGame(lines []string) *Game {
	height := len(lines)
	width := 0
	for i := range lines {
		if len(lines[i]) > width {
			width = len(lines[i])
		}
	}
	g := &Game{
		Height: height,
		Width:  width,
		Grid:   make([][]string, height),
	}
	for i := range lines {
		cells := strings.Split(lines[i], "")
		if len(cells) < width {
			for i := width - len(cells); i > 0; i-- {
				cells = append(cells, " ")
			}
		}
		g.Grid[i] = cells
	}

	return g
}

func main() {
	input := `What? 
This is exceedingly silly. 

Really, we would like some ACTUAL programming challenges around here.`

	lines := strings.Split(input, "\n")
	game := MakeGame(lines)
	prev := ""
	for {
		latest := game.String()
		if latest == prev || game.IsDead() {
			break
		}
		prev = latest
		fmt.Printf("\r--------\n%s\n--------\n", latest)
		game.Iterate()
	}
}
