package main

import "fmt"

func computerPrice(rate float32, nights int) float32 {
	return rate * float32(nights)
}

func main() {
	fmt.Println(computerPrice(100, 2))
}
