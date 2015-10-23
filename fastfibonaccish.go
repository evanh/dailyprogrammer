package main

import (
	"fmt"
)

func FindHighestFactor(n int64, factors []int64) int64 {
	for i := len(factors) - 1; i >= 0; i-- {
		if n%factors[i] == 0 {
			return n / factors[i]
		}
	}
	return -1
}

func GenerateSequence(start, end int64) []int64 {
	var output []int64 = []int64{0}
	curval := start
	var prevval int64 = 0
	for curval < end {
		tmp := curval
		curval += prevval
		prevval = tmp
		output = append(output, curval)
	}
	return output
}

func Stringify(sequence []int64) string {
	var output string = "0"
	for i := range sequence {
		output += fmt.Sprintf(" %d", sequence[i])
	}
	return output
}

func main() {
	var input int64 = 62610760266540248
	if input == 0 {
		fmt.Println("0")
		return
	}

	possibleFactors := GenerateSequence(1, input)
	if possibleFactors[len(possibleFactors)-1] == input {
		fmt.Println(Stringify(possibleFactors))
	}

	factor := FindHighestFactor(input, possibleFactors)
	fmt.Println(Stringify(GenerateSequence(factor, input)))
}
