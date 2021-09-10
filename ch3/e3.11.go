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
	sub := basename(s)
	var j int
	for i := len(sub)-1; i>=0; i-- {
		j++
		buf.WriteByte(sub[i])
		if j%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
	}
	var rbuf bytes.Buffer
	for i := len(buf.Bytes())-1; i>=0; i-- {
		rbuf.WriteByte(buf.Bytes()[i])
	}
	for _, v := range s[len(sub):] {
		rbuf.WriteByte(byte(v))
	}
	return rbuf.Bytes()
}

func basename(s string) string {
	for i := len(s)-1; i>=0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}