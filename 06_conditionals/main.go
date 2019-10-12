package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Println("x equals y")
	} else {
	}

	//switch

	color := "red"
	switch color {
	case "blue":
		fmt.Println("Color is blue")
	case "red":
		fmt.Println("Color is red")
	}
}
