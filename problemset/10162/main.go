package main

import "fmt"

func main() {
	var number int
	fmt.Scanf("%d", &number)

	if number % 2 == 0 {
    fmt.Println("Bala Barare")
  } else {
    fmt.Println("Payin Barare")
  }
}
