package main

import (
	"bytes"
	"fmt"
	// "os"
	// "runtime/pprof"
)

const OFFSET byte = 48 // '0' = 48

func sum(input []byte) int {
	var ret byte
	for _, b := range input {
		ret += b - OFFSET
	}
	return int(ret)
}

func reverse(input []byte) string {
	l := len(input)
	output := make([]byte, l)

	for i := range input {
		output[l-i-1] = input[i]
	}

	return string(output)
}

func IsDescriptive(input []byte) bool {
	if sum(input) != len(input) {
		return false
	}

	for i, b := range input {
		count := int(b - OFFSET)
		toFind := byte(len(input)-i-1) + OFFSET
		if bytes.Count(input, []byte{toFind}) != count {
			return false
		}
	}
	return true
}

func Incr(input []byte, maxDigit byte) bool {
	var carryover byte
	input[0] += 1
	if input[0] > maxDigit {
		input[0] = OFFSET
		carryover = 1
	}

	i := 1
	for i < len(input) && carryover > 0 {
		input[i] += carryover
		carryover = 0
		if input[i] > maxDigit {
			input[i] = OFFSET
			carryover = 1
		}
		i++
	}

	// Returns false if we overflow the number
	return carryover == 0
}

// Heuristics
// Not possible to have all the same number, so we don't have to go past n-2.n-2 etc.
// Actually, don't need to go past n-1 0 ** n - 1
// Start from 1.0**n-1
func SearchRange(input int) (ret []string) {
	value := make([]byte, input)
	value[input-1] = OFFSET + 1
	value[input-2] = OFFSET + 2
	for i := 2; i < input; i++ {
		value[i] = OFFSET
	}

	maxDigit := byte(input-2) + OFFSET
	if maxDigit > OFFSET+9 {
		maxDigit = OFFSET + 9
	}

	for Incr(value, maxDigit) {
		if IsDescriptive(value) {
			ret = append(ret, reverse(value))
		}
	}
	return ret
}

func main() {
	// f, _ := os.Create("descriptivenumbers.pprof")
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	nums := SearchRange(13)
	if len(nums) > 0 {
		for i := range nums {
			fmt.Println(nums[i])
		}
	} else {
		fmt.Println("No valid numbers")
	}
}
