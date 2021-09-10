package main

import (
	"fmt"
)

func main() {
	var a [3]int = [3]int{420, 69}
	for i, v := range a {
		fmt.Println(i, v)
	}
}