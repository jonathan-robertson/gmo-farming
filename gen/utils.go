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

func calculateStandardNameSuffix(tier int, traits string) string {
	return fmt.Sprintf("T%d%s", tier, traits)
}

func calculateCropYield(count, tier int, traits string) int {
	switch tier {
	case 2:
		count *= 2
	case 3:
		count *= 4
	}

	// [B] Bonus
	if strings.Contains(traits, "BB") {
		count *= 2
	} else if strings.Contains(traits, "B") {
		count = int(float64(count) * 1.5)
	}

	// [R] Renewable
	if strings.Contains(traits, "R") {
		count = int(float64(count) * .75)
	}

	return count
}

func calculateBonusYield(count, tier int, traits string) int {
	return calculateCropYield(count, tier, traits)
}

func calculateCraftTime(time, tier int, traits string) int {
	if traits == "" && (tier == 2 || tier == 3) {
		time *= 100
	}
	return time
}

func optionallyAddRenewable(traits string, name ...string) string {
	if strings.Contains(traits, "R") {
		return fmt.Sprintf(`<property name="DowngradeBlock" value="%s" />`, strings.Join(name, ""))
	}
	return ""
}
