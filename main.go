package main

import (
	"gen"
)

func main() {
	if err := gen.WritePlantBlocks("Config-Vanilla/blocks.xml"); err != nil {
		panic(err)
	}
	if err := gen.WritePlantRecipes("Config-Vanilla/recipes.xml"); err != nil {
		panic(err)
	}
	if err := gen.WritePlantLocalization("Config-Vanilla/Localization.txt"); err != nil {
		panic(err)
	}
}
