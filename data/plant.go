package data

import (
	"fmt"
	"strings"
)

type Plant interface {
	GetCraftTime() int
	GetDescription() string
	GetDisplayName() string
	GetName() string
	GetPreferredConsumer() string
	IsCompatibleWith(Trait) bool
	WriteBlockStages(chan string, string)
}

var Plants []Plant = []Plant{
	CreateAloe(),
	CreateBlueberry(),
	CreateChrysanthemum(),
	CreateCoffee(),
	CreateCorn(),
	// CreateCotton(),
	// CreateGoldenrod(),
	// CreateGraceCorn(),
	// CreateHops(),
	CreateMushroom(),
	// CreatePotato(),
	// CreatePumpkin(),
	// CreateYucca(),
}

func calculateCropYield(count int, traits string) int {

	// All Enhanced Seeds start wtih double resources
	count *= 2

	// [B] Bonus
	if strings.Contains(traits, "BB") {
		count *= 4
	} else if strings.Contains(traits, "B") {
		count = int(float64(count) * 2)
	}

	// [R] Renewable
	if strings.Contains(traits, "R") {
		count = int(float64(count) * .75)
	}

	return count
}

func calculateBonusYield(count int, traits string) int {
	return calculateCropYield(count, traits)
}

func optionallyAddRenewable(traits string, plant Plant) string {
	if strings.ContainsRune(traits, 'R') {
		return fmt.Sprintf(`<property name="DowngradeBlock" value="planted%s1_%s" />`,
			plant.GetName(),
			traits)
	}
	return ""
}

func getDefaultSeedDescription() string {
	return `Plant these seeds on a craftable Farm Plot block to grow plants for you to harvest.\n\nWhen harvested, there is a 50% chance to get a seed back for replanting.`
}

func getCraftingGroup(traits string) string {
	if traits == "" {
		return "Food/Cooking"
	} else {
		return "SeedEnhancement"
	}
}
