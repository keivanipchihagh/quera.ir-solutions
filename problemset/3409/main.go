package main

import "fmt"

func main() {
	var number int
	fmt.Scanf("%d", &number)

	for i := 1; i <= number; i++ {
		for j :=1; j <= number; j++ {
			fmt.Printf("%d ", i*j)
		}
		fmt.Println()
	}
}