package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var FUNCS = map[string]func(x float64) float64{
	"1": math.Cos,
	"2": func(x float64) float64 { return x - math.Tan(x) },
	"3": func(x float64) float64 { return 1.0 + (1.0 / x) },
	"4": func(x float64) float64 { return 4.0 * x * (1.0 - x) },
}

func DottieNumber(x float64, f func(x float64) float64) float64 {
	var cur = f(x)
	var prev = cur - 1 // Don't want them to accidentally start the same
	for i := 0; i < 1000; i++ {
		prev = cur
		cur = f(cur)
		if cur == prev {
			return cur
		}
	}
	return cur
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var key string
	var exists bool
	var err error
	for !exists {
		fmt.Print("\nPick a function:\n1) cos(x)\n2) x - tan(x)\n3) 1 + (1 / x)\n4) 4x(1 - x)\nEnter number: ")
		key, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Could not read input")
			continue
		}

		key = strings.TrimSpace(key)
		_, exists := FUNCS[key]

		if !exists {
			fmt.Print(key, " is not a valid choice.\n")
		} else {
			break
		}
	}

	f := FUNCS[key]

	var value string
	for {
		fmt.Print("Enter a value for x (q to quit): ")
		value, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Could not read input")
			return
		}
		value = strings.TrimSpace(value)
		if value == "q" {
			return
		}

		x, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println("Invalid value for x.")
			return
		}

		dottie := DottieNumber(x, f)
		fmt.Printf("Calculated %.04f\n", dottie)
	}
}
