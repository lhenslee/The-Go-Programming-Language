package main

import (
	"fmt"
	"ch2/unitconv"
)

func main() {
	yd := unitconv.Yard(3)
	m := unitconv.YToM(yd)
	ft := unitconv.YToF(yd)
	in := unitconv.FToI(ft)
	fmt.Printf("%s = %s = %s = %s\n", yd, m, ft, in)
}