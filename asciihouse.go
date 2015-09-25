package main

import (
	"fmt"
	// "math/rand"
)

type Room struct {
	left  bool
	right bool
	roof  bool
	floor bool
}

func (r *Room) Draw() []string {
	output := make([]string, 3)

	// floor
	output[0] = "   "
	if r.floor {
		output[0] = "---"
		if !r.left {
			output[0] = "-" + output[0]
		} else {
			output[0] = "+" + output[0]
		}
		if !r.right {
			output[0] = output[0] + "-"
		} else {
			output[0] = output[0] + "+"
		}
	} else {
		if r.left {
			output[0] = "|" + output[0]
		} else {
			output[0] = " " + output[0]
		}
		if r.right {
			output[0] = output[0] + "|"
		} else {
			output[0] = output[0] + " "
		}
	}

	// middle
	output[1] = "   "
	if r.left {
		output[1] = "|" + output[1]
	} else {
		output[1] = " " + output[1]
	}
	if r.left {
		output[1] = "|" + output[1]
	} else {
		output[1] = " " + output[1]
	}

	return output
}

func main() {
	r := Room{true, false, true, false}
	fmt.Println(r.Draw())
}
