package main

import (
	"gen"
)

func main() {
	if err := gen.WritePlantBlocks(); err != nil {
		panic(err)
	}
	if err := gen.WritePlantRecipes(); err != nil {
		panic(err)
	}
	if err := gen.WritePlantLocalization(); err != nil {
		panic(err)
	}
}
