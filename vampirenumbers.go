package main

import (
	"fmt"
	"math"
	"sort"
)

func FindFactorials(n, digits int) [][]int {
	// Check to make sure n has the same or more digits as required
	if GetDigitCount(n) < digits {
		return nil
	}
	start := int(math.Pow(10.0, float64(digits-1)))
	var output [][]int
	for i := start; i < (n/2)+1 && GetDigitCount(i) == digits; i++ {
		if j := n / i; j*i == n {
			if j < i {
				break
			}

			jd := GetDigitCount(j)
			if jd > digits {
				// It has more digits than we require. Factor it further.
				if ret := FindFactorials(j, digits); ret != nil {
					for _, factorials := range ret {
						factorials = append(factorials, i)
						output = append(output, factorials)
					}
				}
			} else if jd == digits {
				// It has the exact number of digits we require.
				output = append(output, []int{i, j})
			}
		}
	}
	return output
}

func GetDigitCount(n int) int {
	var i int = 1
	var count int
	for {
		val := n / i
		if val == 0 {
			break
		}
		count++
		i *= 10
	}
	return count
}

func GetDigits(n int) []int {
	var i int = 1
	digits := []int{}
	for {
		val := n / i
		if val == 0 {
			break
		}
		remainder := (n / i) % 10
		digits = append(digits, remainder)
		i *= 10
	}
	return digits
}

func CheckFactorialContained(n int, size int, factorials []int) bool {
	if len(factorials) != size {
		return false
	}
	nDigits := GetDigits(n)
	facDigits := []int{}
	for i := range factorials {
		f := GetDigits(factorials[i])
		facDigits = append(facDigits, f...)
	}
	return CheckIntSliceEqual(nDigits, facDigits)
}

func CheckIntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func CheckHaveFactorial(fac []int, factorials [][]int) bool {
	for i := range factorials {
		if CheckIntSliceEqual(fac, factorials[i]) {
			return true
		}
	}
	return false
}

func PrintFactorial(n int, factorial []int) {
	output := fmt.Sprintf("%d=%d", n, factorial[0])
	for i := 1; i < len(factorial); i++ {
		output += fmt.Sprintf("*%d", factorial[i])
	}
	fmt.Println(output)
}

func main() {
	digits := 6
	size := 3

	// output := [][]int{}
	start := int(math.Pow(10, float64(digits-1)))
	end := int(math.Pow(10, float64(digits)))

	for n := start; n < end; n++ {
		found := [][]int{}
		factorials := FindFactorials(n, 2)
		for i := range factorials {
			if CheckFactorialContained(n, size, factorials[i]) {
				if !CheckHaveFactorial(factorials[i], found) {
					PrintFactorial(n, factorials[i])
					found = append(found, factorials[i])
				}
			}
		}
	}
}
