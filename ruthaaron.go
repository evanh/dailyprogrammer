package main

import (
	"fmt"
)

func GetDistinctPrimeFactors(n int) []int {
	current := n
	factors := []int{}
	more_factors := true
	for more_factors {
		found_factor := false
		for i := 2; i < (current/2)+1; i++ {
			if current%i == 0 {
				if !Contains(i, factors) {
					factors = append(factors, i)
				}
				current = current / i
				found_factor = true
				break
			}
		}
		if !found_factor {
			if !Contains(current, factors) {
				factors = append(factors, current)
			}
			more_factors = false
		}
	}

	return factors
}

func Contains(n int, list []int) bool {
	for i := range list {
		if list[i] == n {
			return true
		}
	}
	return false
}

func Sum(ints []int) int {
	sum := 0
	for i := range ints {
		sum += ints[i]
	}
	return sum
}

func main() {
	var prev_factors []int = []int{1}
	var cur_factors []int

	for i := 2; i < 5000; i++ {
		cur_factors = GetDistinctPrimeFactors(i)
		if Sum(prev_factors) == Sum(cur_factors) {
			fmt.Printf("(%d,%d) VALID\n", i-1, i)
		}
		prev_factors = cur_factors
	}
}
