package main

import (
	"fmt"
	"math"
)

func IsPerfectSquare(f float64) bool {
	val := math.Sqrt(f)
	return float64(int64(val)) == val
}

func IsFibonacci(n int64) bool {
	// Awesome formula, courtesy of
	// https://www.quora.com/What-is-the-most-efficient-algorithm-to-check-if-a-number-is-a-Fibonacci-Number

	first := 5.0*math.Pow(float64(n), 2.0) + 4
	if IsPerfectSquare(first) {
		return true
	}

	second := 5.0*math.Pow(float64(n), 2.0) - 4
	return IsPerfectSquare(second)
}

func FindLowestFactor(n int64) int64 {
	// Take advantage of the fact that n = c * F(i), where F is the fibonacci
	// sequence. Read about this on
	// http://www.maths.surrey.ac.uk/hosted-sites/R.Knott/Fibonacci/fibGen.html
	var i int64
	for i = 2; i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			val := n / i
			if IsFibonacci(i) {
				return val
			}
			if IsFibonacci(val) {
				return i
			}
		}
	}
	return -1
}

func GenerateSequence(start, end int64) string {
	var output string
	output += "0" + fmt.Sprintf(" %d", start)
	curval := start
	var prevval int64 = 0
	for curval < end {
		tmp := curval
		curval += prevval
		prevval = tmp
		output += fmt.Sprintf(" %d", curval)
	}
	if curval != end {
		output += "... INVALID SEQUENCE"
	}
	return output
}

func main() {
	var input int64 = 37889062373143906
	if input == 0 {
		fmt.Println("0")
	} else if IsFibonacci(input) {
		fmt.Println(GenerateSequence(1, input))
	} else {
		factor := FindLowestFactor(input)
		fmt.Println(GenerateSequence(factor, input))
	}
}
