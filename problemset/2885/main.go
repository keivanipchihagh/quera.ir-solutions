package main

import (
	"fmt"
	"strings"
)

func main() {

	var number int

	fmt.Scanf("%d", &number)
	fmt.Println(strings.Repeat("man khoshghlab hastam\n", number))
}
