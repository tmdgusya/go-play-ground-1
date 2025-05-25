package main

import (
	"fmt"

	"go-play-ground-1/usepkg/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	graph := asciigraph.Plot(data, asciigraph.Height(10), asciigraph.Caption("Sample Graph"))

	fmt.Println(graph)
}
