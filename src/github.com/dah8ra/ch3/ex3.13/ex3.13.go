package main

import (
	"fmt"
)

type ByteSize float64

const (
	KB ByteSize = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Printf("%g\n", KB)
	fmt.Printf("%g\n", MB)
	fmt.Printf("%g\n", GB)
	fmt.Printf("%g\n", TB)
	fmt.Printf("%g\n", PB)
	fmt.Printf("%g\n", EB)
	fmt.Printf("%g\n", ZB)
	fmt.Printf("%g\n", YB)
}
