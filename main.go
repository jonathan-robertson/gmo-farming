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

		&gen.ResearcherBlocks{Path: "Config-Researcher-Loot"},
		&gen.ResearcherItems{Path: "Config-Researcher-Loot"},
		&gen.ResearcherLocalization{Path: "Config-Researcher-Loot"},
		&gen.ResearcherRecipes{Path: "Config-Researcher-Loot"},
		&gen.ResearcherUIDisplay{Path: "Config-Researcher-Loot"},
		&gen.ResearcherLoot{},

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
