package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	buf.WriteString(s)
	var result bytes.Buffer
	for {
		temp := buf.Next(3)
		fmt.Println(temp)
		if len(temp) < 3 {
			result.Write(temp)
			return result.String()
		}
		result.Write(temp)
		result.WriteByte(44)
	}
	return result.String()
}
