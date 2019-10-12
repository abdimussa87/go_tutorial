package main

import "fmt"

func main() {

	emails := make(map[string]string)

	emails["abdi"] = "abdi@gmail.com"
	fmt.Println(emails)

	address := map[string]string{"mussa": "mussa@gmail.com"}

	fmt.Println(address)
}
