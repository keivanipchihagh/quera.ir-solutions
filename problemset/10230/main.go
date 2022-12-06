package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if (a + b > c) && (b + c > a) && (a + c > b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}