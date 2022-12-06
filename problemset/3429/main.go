package main

import "fmt"

func state(degree int) string {
	switch {
	case degree > 100:
		return "Steam"
	case degree < 0:
		return "Ice"
	default:
		return "Water"
	}
}

func main() {
	var degree int
	fmt.Scanf("%d", &degree)
	fmt.Println(state(degree))
}