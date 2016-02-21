package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", decimalComma(os.Args[i]))
	}
}

func decimalComma(s string) string {
	split := strings.Split(s, ".")
	var buf bytes.Buffer
	for i, _ := range split {
		if i == 0 {
			buf.WriteString(commaFromTail(split[0]))
		} else {
			buf.WriteString(commaFromTop(split[1]))
			break
		}
		buf.WriteString(".")
	}
	return buf.String()
}

func commaFromTail(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaFromTail(s[:n-3]) + "," + s[n-3:]
}

func commaFromTop(s string) string {
	var buf bytes.Buffer
	buf.WriteString(s)
	var result bytes.Buffer
	for {
		temp := buf.Next(3)
		if len(temp) < 3 {
			result.Write(temp)
			return result.String()
		}
		result.Write(temp)
		result.WriteByte(44)
	}
	return result.String()
}
