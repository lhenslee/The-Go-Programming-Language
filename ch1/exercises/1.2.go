// Modify the echo program to print the index and value of each of its arguments, one per line.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, "\n"))
}