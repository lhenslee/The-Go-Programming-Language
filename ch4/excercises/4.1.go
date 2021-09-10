// Write a function that counts the number of bits that are in two SHA256 hashes
package main

import (
	"fmt"
	"learngo/ch2"
)

const (
	//hello = 0x2CF24DBA5FB0A30E26E83B2AC5B9E29E1B161E5C1FA7425E73043362938B9824
	hi = "98EA6E4F216F2FB4B69FFF9B3A44842C38686CA685F3F55DC48C5D3FB1107BE4"
)

func main() {
	var hello uint64 = 0x2CF24DBA5FB0A30E26E83B2AC5B9E29E1B161E5C1FA7425E73043362938B9824
	fmt.Println(popcount.PopCount(uint64(hello)))
}