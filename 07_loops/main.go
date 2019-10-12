package main

import "fmt"

func main() {

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Number is %d \n", i)
	// }

	for i := 1; i < 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("Fiz Buzz %d \n", i)
		} else if i%3 == 0 {
			fmt.Printf("Fizz %d \n", i)
		} else if i%5 == 0 {
			fmt.Printf("buzz %d \n", i)
		}
	}
}
