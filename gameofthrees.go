package main

import (
	"fmt"
)

func main() {
	num := 31337357
	for num > 1 {
		switch num % 3 {
		case 0:
			fmt.Println(num, 0)
			num = num / 3
		case 1:
			fmt.Println(num, -1)
			num = (num - 1) / 3
		case 2:
			fmt.Println(num, 1)
			num = (num + 1) / 3
		}
	}
	fmt.Println(num)
}
