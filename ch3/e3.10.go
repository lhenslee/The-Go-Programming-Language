package main

import (
	"fmt"
	"bytes"
	"os"
)

func main() {
	fmt.Println(string(comma(os.Args[1])))
}

func comma(s string) []byte {
	var buf bytes.Buffer
	var j int
	for i := len(s)-1; i>=0; i-- {
		j++
		buf.WriteByte(s[i])
		if j%3 == 0 {
			buf.WriteByte(',')
		}
	}
	var rbuf bytes.Buffer
	for i := len(buf.Bytes())-1; i>=0; i-- {
		rbuf.WriteByte(buf.Bytes()[i])
	}
	return rbuf.Bytes()
}