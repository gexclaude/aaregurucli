package main

import (
	"fmt"
	"time"
)

func main() {
	x := "This is the text to be printed to console"
	
	for i := 0; i < len(x); i++ {
		fmt.Printf("\r %-41s", x[i:])
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Printf("\r %-41s", "")
	fmt.Println()
}