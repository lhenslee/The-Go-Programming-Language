package main

import (
	"fmt"
	"tempconv"
)

func main() {
	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToK(tempconv.FreezingC))
}