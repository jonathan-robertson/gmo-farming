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
		&gen.ResearcherRecipes{},
		&gen.ResearcherUIDisplay{},
		&gen.StandardBlocks{},
		&gen.StandardLocalization{},
		&gen.StandardRecipes{},
		&gen.StandardUIDisplay{},
	} {
		if err := gen.Write(producer); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
	}
}
