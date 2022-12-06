package main

import (
	"fmt"
	"strings"
)

func main() {

	var number int

	fmt.Scanf("%d", &number)
	fmt.Println("W" + strings.Repeat("o", number) + "w!")
}
