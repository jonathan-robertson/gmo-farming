package main

import (
	"fmt"
	"gen"
	"os"
)

func main() {
	for _, producer := range []gen.Producer{
		&gen.ResearcherBlocks{},
		&gen.ResearcherItems{},
		&gen.ResearcherLocalization{},
		&gen.ResearcherProgression{},
		&gen.ResearcherRecipes{},
		&gen.ResearcherUIDisplay{},
		&gen.StandardBlocks{},
		&gen.StandardLocalization{},
		&gen.StandardProgression{},
		&gen.StandardRecipes{},
		&gen.StandardUIDisplay{},
	} {
		if err := gen.Write(producer); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
	}
}
