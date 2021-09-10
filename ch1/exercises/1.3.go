// Experiment to measure the difference in running time between our 
// potentially inefficient versions and the one that uses strings.Join.
// Use the time package and write benchmark tests for systematic performance evaluation.
package main

import (
	"strings"
	"time"
	"os"
	"fmt"
)

func main() {
	start := time.Now()
	OldMethod()
	fmt.Println("Time for OldMethod =", time.Since(start).Microseconds(), "MS")
	start = time.Now()
	NewMethod()
	fmt.Println("Time for NewMethod =", time.Since(start).Microseconds(), "MS")
}

func OldMethod() {
	var s, sep string
	for i := 1; i<len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func NewMethod() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}