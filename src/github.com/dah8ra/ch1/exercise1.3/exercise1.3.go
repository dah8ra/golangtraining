package main 

import (
	"time"
	"fmt"
	"strings"
	"os"
)

func main() {
	str := "a"
	elapsed := time.Now()
	for i:=1; i<10 ; i++ { 
		str += str
		//fmt.Println(i)
	}
	//nanosec:=time.Since(elapsed).Nanoseconds()
	//fmt.Printf("%d sec", nanosec)
	sec:=time.Since(elapsed).Seconds()
	fmt.Printf("No effective  : %.8f sec", sec)
	
	fmt.Println()
	
	elapsedJoin := time.Now()
	for i:=1 ; i< 10 ; i++{
		strings.Join(os.Args[1:], "a")
	}
	//nanosecJoin := time.Since(elapsedJoin).Nanoseconds()
	//fmt.Printf("Join effective: %d nano sec", nanosecJoin)
	secJoin := time.Since(elapsedJoin).Seconds()
	fmt.Printf("Join effective: %.8f sec", secJoin)
	
	fmt.Println()
	
	diff := sec-secJoin
	fmt.Printf("Diff = %.8f sec", diff)
}
