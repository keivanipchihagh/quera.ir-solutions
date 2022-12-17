package main

import "fmt"

func main() {

	list := make([]int, 1000)

	for {
		var number int
		fmt.Scanf("%d", &number)

		if number == 0 {
			break
		}
		list = append(list, number)
	}

	for i := range list {
		i = len(list) - 1 - i
		if list[i] != 0 {
			fmt.Println(list[i])
		}
	}
}