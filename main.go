package main

import "fmt"

func main() {
	s := "Hello, Golang"
	for _, v := range s {
		// v is rune (rune is alias for int32)
		fmt.Printf("Unicode code point : %U - character: %c - binary %b hex %x Decimal %d\n", v, v, v, v, v)
	}
}
