package main

import (
	"fmt"
)

const OFFSET byte = 48 // '0' = 48

func BuildDescription(num []int) []int {
	number := make([]int, len(num))
	for _, n := range num {
		number[n]++
	}
	return number
}

var Partitions [][]int = [][]int{}

func GetPartitions(target, maxValue, digits int, soFar []int) {
	if target == 0 {
		// Filter cases that can't be self descriptive
		if len(soFar) == 1 || len(soFar) == digits {
			return
		}

		// Zero pad partition
		if len(soFar) < digits {
			soFar = append(soFar, make([]int, digits-len(soFar))...)
		}
		Partitions = append(Partitions, soFar)
	} else {
		if maxValue > 1 {
			GetPartitions(target, maxValue-1, digits, soFar)
		}
		if maxValue <= target {
			GetPartitions(target-maxValue, maxValue, digits, append(soFar, maxValue))
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func toString(num []int) string {
	ret := make([]byte, len(num))
	for i := range num {
		ret[i] = byte(num[i]) + OFFSET
	}
	return string(ret)
}

func main() {
	total := 15

	for input := 15; input <= total; input++ {
		Partitions = [][]int{}
		GetPartitions(input, input-1, input, []int{})

		selfdescriptors := []string{}
		for i := range Partitions {
			num1 := BuildDescription(Partitions[i])
			num2 := BuildDescription(num1)
			if equal(num1, num2) {
				selfdescriptors = append(selfdescriptors, toString(num1))
			}
		}

		if len(selfdescriptors) > 0 {
			for _, s := range selfdescriptors {
				fmt.Println(input, s)
			}
		}
	}
}
