package main

import (
	"fmt"

	"github.com/gaspiman/uniPOS"
)

func main() {
	// Map a POS tag from one tagset to the Universal tagset. Ex. WRB from Penn Treebank
	pos := uniPOS.Map("en-ptb", "WRB")
	fmt.Println("PennTreebank tag: WRB = UniversalPOS:", pos)
	// Mapping for only a specific tagset. ex: Penn Treebank
	uTagset := uniPOS.GetMap("en-ptb")
	fmt.Println(uTagset)
}
