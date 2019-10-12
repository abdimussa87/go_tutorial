package main

import "fmt"

func main() {

	ids := []int{33, 25, 88}

	for _, id := range ids {

		fmt.Println(id)
	}

	address := map[string]string{"mussa": "mussa@gmail.com"}

	for k, v := range address {
		fmt.Printf("%s : %s \n", k, v)
	}
}
