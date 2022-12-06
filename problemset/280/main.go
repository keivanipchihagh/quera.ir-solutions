package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if (a*a + b*b == c*c) || (b*b + c*c == a*a) || (a*a + c*c == b*b) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
