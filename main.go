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

		&gen.ResearcherBlocks{Path: "Config-Researcher-Rewards"},
		&gen.ResearcherItems{Path: "Config-Researcher-Rewards"},
		&gen.ResearcherRewardsLocalization{Path: "Config-Researcher-Rewards"},
		&gen.ResearcherRecipes{Path: "Config-Researcher-Rewards"},
		&gen.ResearcherUIDisplay{Path: "Config-Researcher-Rewards"},
		&gen.ResearcherRewardsLoot{},

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
