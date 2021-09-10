package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(0)
	}
	fmt.Println(anagram(os.Args[1], os.Args[2]))
}

func anagram(s1, s2 string) bool {
	var i int
	if len(s1) != len(s2) {
		return false
	}
	for _, v := range s1 {
		if has(s2, v) {
			i++
		}
	}
	if i == len(s1) {
		return true
	}
	return false
}

func has(s string, l rune) bool {
	for _, v := range s {
		if v == l {
			return true
		}
	}
	return false
}