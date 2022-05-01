package gen

import (
	"os"
)

func getFile(filename string) (*os.File, error) {
	os.Remove(filename)
	return os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
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
