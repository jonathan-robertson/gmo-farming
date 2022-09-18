package data

import (
	"fmt"
	"strings"
)

// Plant interface describes the common interactions one can take with all plants
type Plant interface {
	GetCraftTime() int
	GetDescription() string
	GetDisplayName() string
	GetName() string
	GetSchematicName(string) string
	IsCompatibleWith(Trait) bool
	WriteBlockStages(chan string, string, string)
}

// Plants is a collection of all plants to be processed
var Plants []Plant = []Plant{
	&Aloe{
		Name:        "Aloe",
		DisplayName: "Aloe Vera",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
	&Blueberry{
		Name:        "Blueberry",
		DisplayName: "Blueberry",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
	&Chrysanthemum{
		Name:        "Chrysanthemum",
		DisplayName: "Chrysanthemum",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
	&Coffee{
		Name:        "Coffee",
		DisplayName: "Coffee",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
	&Corn{
		Name:        "Corn",
		DisplayName: "Corn",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
	&Cotton{
		Name:        "Cotton",
		DisplayName: "Cotton",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
	&Goldenrod{
		Name:        "Goldenrod",
		DisplayName: "Goldenrod",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
	&GraceCorn{
		Name:        "GraceCorn",
		DisplayName: "Super Corn",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
	&Hop{
		Name:        "Hop",
		DisplayName: "Hop",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
	&Mushroom{
		Name:               "Mushroom",
		DisplayName:        "Mushroom Spores",
		Description:        `Mushroom spores can be planted on all surfaces and walls and will grow without sunlight.`,
		CropYield:          2,
		BonusYield:         1,
		CraftTime:          2,
		incompatibleTraits: []rune{'U'},
	},
	&Potato{
		Name:        "Potato",
		DisplayName: "Potato",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
	&Pumpkin{
		Name:        "Pumpkin",
		DisplayName: "Pumpkin",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
	&Yucca{
		Name:        "Yucca",
		DisplayName: "Yucca",
		CropYield:   2,
		BonusYield:  1,
		CraftTime:   2,
	},
}

func calculateCropYield(count int, traits string) int {

	// All Enhanced Seeds start with double resources
	count *= 2

	// [B] Bonus
	if strings.Contains(traits, "BB") {
		count *= 4
	} else if strings.Contains(traits, "B") {
		count *= 2
	}

	return count
}

func calculateBonusYield(count int, traits string) int {
	return calculateCropYield(count, traits)
}

func calculatePlantTier(traits string) (tier int) {
	return len(traits) + 1
}

func getRenewableAndDropTags(traits string, plant Plant) string {
	if strings.ContainsRune(traits, 'R') {
		return fmt.Sprintf(`
    <property name="DowngradeBlock" value="planted%s1_%s" />
    <drop event="Destroy" count="0" />`,
			plant.GetName(),
			traits)
	}
	return fmt.Sprintf(`<drop event="Destroy" name="planted%s1_%s" count="1" prob="0.5"/>`,
		plant.GetName(),
		traits)
}

func getUnlock(plant Plant, target, traits string) string {
	switch target {
	case "Researcher":
		return fmt.Sprintf(`
    <property name="UnlockedBy" value="%s"/>`, plant.GetSchematicName(traits))
	default:
		return ""
	}
}

func getDefaultSeedDescription() string {
	return `Plant these seeds on a craftable Farm Plot block.`
}

func getCraftingGroup(traits string) string {
	return fmt.Sprintf("Tier%dSeeds", len(traits)+1)
}

func getItemTypeIcon(traits string) string {
	switch len(traits) {
	case 0:
		return "block_upgrade"
	case 1:
		return "add"
	case 2:
		return "healing_factor"
	default:
		return ""
	}
}
