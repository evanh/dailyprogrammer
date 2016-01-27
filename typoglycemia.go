package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	input := []byte("Some text, do you see the really really long words?")
	words := bytes.Split(input, ' ')

}
