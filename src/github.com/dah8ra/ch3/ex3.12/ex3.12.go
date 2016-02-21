package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	key := os.Args[1]
	target := os.Args[2]
	if findAnagram(key, target) {
		fmt.Printf("RESULT: TRUE -> %s %s\n", key, target)
	} else {
		fmt.Printf("RESULT: FALSE ->  %s %s\n", key, target)
	}
}

func findAnagram(key string, target string) bool {
	var keyBuf bytes.Buffer
	keyBuf.WriteString(key)
	var flag bool = false
	for {
		vkey := keyBuf.Next(1)
		svkey := string(vkey)
		if svkey == "" {
			break
		}
		var targetBuf bytes.Buffer
		targetBuf.WriteString(target)
		for {
			vtarget := targetBuf.Next(1)
			svtarget := string(vtarget)
			if svtarget == "" {
				break
			}
			if svkey == svtarget {
				flag = true
				break
			} else {
				flag = false
			}
		}
		if !flag {
			break
		}
	}
	return flag
}
