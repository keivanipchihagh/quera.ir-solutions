package main

import "fmt"

func main() {
	var number int
	fmt.Scanf("%d", &number)

	for {
		number = sum_digits(number)
		if number < 10 {
			break
		}
	}

	fmt.Println(sum_digits(number))
}

func sum_digits(number int) int {
	sum := 0
	for i := number; i >= 1; i /= 10 {
		sum += i % 10
	}
	return sum
}