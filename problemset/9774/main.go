package main

import (
	"fmt"
	"strings"
)

func main() {
	var number string
	fmt.Scanf("%s", &number)

	for _, char := range number {
		fmt.Println(string(char) + ":", strings.Repeat(string(char), int(char - '0')))
	}
}
