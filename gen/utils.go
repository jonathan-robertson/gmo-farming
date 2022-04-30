package gen

import (
	"fmt"
	"os"
	"strings"
)

func getFile(filename string) (*os.File, error) {
	os.Remove(filename)
	return os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
}

func calculateStandardNameSuffix(traits string) string {
	return fmt.Sprintf("T%d%s", traits)
}

func calculateCropYield(count int, traits string) int {

	// All GMO Seeds start wtih double resources
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

func calculateCraftTime(time int, traits string) int {
	if traits == "" {
		time *= 450
	}
	return time
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

func getEnhancedSeedEffectDescription() string {
	return `All GMO Seeds twice the crop yield.`
}

func getInitialEnhancementCraftingTip() string {
	return `This can be crafted by hand or in a Hot Box.`
}

func getHotBoxRequirementTip() string {
	return `A Hot Box is required to craft this.`
}
