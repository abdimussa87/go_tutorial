package main

import "fmt"

func main() {
	fmt.Println(greeting("Abdi"))
	fmt.Println(sum(1, 3))
}

func greeting(name string) string {
	return "Hello" + name
}

func sum(num1, num2 int) int {
	return num1 + num2
}
