package main

import (
	"fmt"
	"math/rand"
	"time"
)

var SOURCE = []byte{'O', 'I', 'S', 'Z', 'L', 'J', 'T'}

type Bag struct {
	Values    []byte
	Visited   map[int]bool
	Remaining int
}

func (b *Bag) Next() byte {
	if b.Remaining == 1 {
		// Reset visited
		var last byte
		for k := range b.Values {
			if b.Visited[k] == false {
				last = b.Values[k]
			}
			b.Visited[k] = false
		}
		b.Remaining = len(b.Values)
		return last
	}

	index := rand.Intn(len(b.Values))
	var ret byte
	if !b.Visited[index] {
		b.Visited[index] = true
		b.Remaining -= 1
		ret = b.Values[index]
	} else {
		// Search for the closest unvisited
		for i := 1; i < len(b.Values); i++ {
			testIndex := (index + i) % len(b.Values)
			if !b.Visited[testIndex] {
				b.Visited[testIndex] = true
				b.Remaining -= 1
				ret = b.Values[testIndex]
				break
			}
		}
	}
	return ret
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var visited = make(map[int]bool)
	for i := range SOURCE {
		visited[i] = false
	}

	bag := &Bag{
		Values:    SOURCE,
		Visited:   visited,
		Remaining: len(SOURCE),
	}

	iterations := 49
	var output string
	for i := 0; i < iterations; i++ {
		output += fmt.Sprintf("%c", bag.Next())
	}
	fmt.Println(output)

	// Verify output
	for i := 0; i < iterations; i += 7 {
		toTest := output[i : i+7]
		seenBefore := make(map[rune]bool)
		for _, t := range toTest {
			if _, ok := seenBefore[t]; !ok {
				seenBefore[t] = true
			} else {
				fmt.Println("INVALID SECTION", i, toTest)
			}
		}
	}
}
