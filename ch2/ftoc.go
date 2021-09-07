package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g F = %g C\n", freezingF, FToC(freezingF))
	fmt.Printf("%g F = %g C\n", boilingF, FToC(boilingF))
}

func FToC(f float64) float64 {
	return (f-32*5/9)
}