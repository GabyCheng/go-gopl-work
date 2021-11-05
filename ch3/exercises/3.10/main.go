package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "123456"
	fmt.Println(comma(s))

}

func comma(s string) string {

	if s == "" {
		return ""
	}

	var buf bytes.Buffer
	str := s[0]
	buf.WriteByte(str)

	for i := 1; i < len(s); i++ {
		if (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}

	return buf.String()
}
