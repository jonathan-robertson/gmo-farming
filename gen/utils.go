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
		time *= 1000
	}
	return time
}

func optionallyAddRenewable(traits string, name ...string) string {
	if strings.Contains(traits, "R") {
		return fmt.Sprintf(`<property name="DowngradeBlock" value="%s" />`, strings.Join(name, ""))
	}
	return ""
}
