package main

import (
	"fmt"
	"math/rand"
	"time"
)

var SOURCE = []byte{'O', 'I', 'S', 'Z', 'L', 'J', 'T'}

func main() {
	rand.Seed(time.Now().UnixNano())

	indices := rand.Perm(len(SOURCE))
	var output string
	for i := 0; i < 50; i++ {
		if i != 0 && i%7 == 0 {
			indices = rand.Perm(len(SOURCE))
		}
		output += fmt.Sprintf("%c", SOURCE[indices[i%7]])
	}
	fmt.Println(output)
}
