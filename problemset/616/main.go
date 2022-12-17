package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scanf("%d", &number)

	for i := number + 1; true; i++ {
		if PowerOfTwo(i) == 0 {
			fmt.Println(i)
			break
		}
	}
}

func PowerOfTwo(n int) int {
	return n & (n - 1)
 }