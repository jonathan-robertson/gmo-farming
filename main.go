package main

import (
	"fmt"
	"gen"
	"os"
)

func main() {
	for _, target := range []string{"Vanilla", "CrystalHell"} {
		if err := generateFiles(target); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
	}
}

func generateFiles(target string) error {
	if err := gen.WritePlantBlocks(target); err != nil {
		return err
	}
	if err := gen.WritePlantRecipes(target); err != nil {
		return err
	}
	if err := gen.WritePlantLocalization(target); err != nil {
		return err
	}
	if err := gen.WriteItems(target); err != nil {
		return err
	}
	return nil
}
