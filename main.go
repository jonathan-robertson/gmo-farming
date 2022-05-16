package main

import (
	"fmt"
	"gen"
	"os"
)

func main() {
	for _, producer := range []gen.Producer{
		&gen.CrystalHellBlocks{},
		&gen.CrystalHellLocalization{},
		&gen.CrystalHellProgression{},
		&gen.CrystalHellRecipes{},
		&gen.CrystalHellUIDisplay{},
		&gen.VanillaBlocks{},
		&gen.VanillaItems{},
		&gen.VanillaLocalization{},
		&gen.VanillaProgression{},
		&gen.VanillaRecipes{},
		&gen.VanillaUIDisplay{},
	} {
		if err := gen.Write(producer); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
	}
}
