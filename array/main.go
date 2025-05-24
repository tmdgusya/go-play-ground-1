package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.0, 26.0, 27.0, 28.0}

	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}
}
