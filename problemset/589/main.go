package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n - 1)
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(factorial(number))
}