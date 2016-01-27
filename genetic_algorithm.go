package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const PARENT_CHOICE = 0.5
const GENERATION_SIZE = 10000
const PARENT_POOL_SIZE = 2

var TARGET = []byte("Hello, World!")
var TARGET_LEN = len(TARGET)
var CHILD_MUTATION_RATE = 1.0 / float64(TARGET_LEN)

type Pool [][]byte

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Less(i, j int) bool {
	return Fitness(p[i]) < Fitness(p[j])
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var ParentPool = make(Pool, PARENT_POOL_SIZE)
var Generation = 0
var Candidates = make(Pool, GENERATION_SIZE)

func Fitness(candidate []byte) int {
	distance := 0
	for i := range candidate {
		if candidate[i] != TARGET[i] {
			distance += 1
		}
	}
	return distance
}

func Reproduce(p1, p2 []byte) []byte {
	child := make([]byte, TARGET_LEN)

	var best, worst []byte
	if Fitness(p1) < Fitness(p2) {
		best = p1
		worst = p2
	} else {
		best = p2
		worst = p1
	}

	for i := range p1 {
		if rand.Float64() < CHILD_MUTATION_RATE {
			child[i] = byte(rand.Intn(128))
		} else if rand.Float64() < PARENT_CHOICE {
			child[i] = best[i]
		} else {
			child[i] = worst[i]
		}
	}

	return child
}

func CreateNewGeneration() {
	for i := range Candidates[:PARENT_POOL_SIZE] {
		ParentPool[i] = Candidates[i]
	}

	// Replace all the non-parents with the children
	for i := PARENT_POOL_SIZE; i < len(Candidates); i++ {
		p1 := ParentPool[rand.Intn(PARENT_POOL_SIZE)]
		p2 := ParentPool[rand.Intn(PARENT_POOL_SIZE)]
		Candidates[i] = Reproduce(p1, p2)
	}
}

func Life() bool {
	sort.Sort(Candidates)
	if bytes.Equal(Candidates[0], TARGET) {
		return true
	}
	CreateNewGeneration()
	return false
}

func Initialize() Pool {
	p := make(Pool, GENERATION_SIZE)
	for i := range p {
		p[i] = make([]byte, TARGET_LEN)
		for j := range p[i] {
			p[i][j] = byte(rand.Intn(128))
		}
	}
	return p
}

func main() {
	rand.Seed(time.Now().UnixNano())
	Candidates = Initialize()
	for !Life() {
		fmt.Printf("Gen: %02d | Fitness: %02d | %s\n", Generation, Fitness(Candidates[0]), Candidates[0])
		Generation++
	}
	fmt.Printf("Gen: %02d | Fitness: %02d | %s\n", Generation, Fitness(Candidates[0]), Candidates[0])
}
