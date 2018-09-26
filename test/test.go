package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, i := range a {
		fmt.Println(i)
	}

	fmt.Println(a)
}

func aa(s int) {
	fmt.Println(s)
}
