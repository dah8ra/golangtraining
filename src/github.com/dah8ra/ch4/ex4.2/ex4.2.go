package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var s = flag.Int("sha", 0, "Hashing by sha256 (default)")

func main() {
	flag.Parse()
	//fmt.Printf("%g\n", flag.Args())
	//fmt.Printf("s: %v\n", *s)
	if *s == 384 {
		fmt.Printf("SHA384 -> %x\n", sha512.Sum384([]byte(flag.Arg(0))))
	} else if *s == 512 {
		fmt.Printf("SHA512 -> %x\n", sha512.Sum512([]byte(flag.Arg(0))))
	} else {
		fmt.Printf("SHA256 -> %x\n", sha256.Sum256([]byte(flag.Arg(0))))
	}
}
