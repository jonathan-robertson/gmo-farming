package data

import (
	"fmt"
)

// Trait is a collection of data representing a special skill or
// attribute that can be applied to Plants
type Trait struct {
	Code               rune
	Name               string
	DoubleName         string
	IncompatibleTraits []rune
	Ingredients        []Ingredient
}

// Ingredient represents the name and quantity of item needed to craft something
type Ingredient struct {
	Name  string
	Count int
}

// Traits is a collection of all Trait types
var Traits []Trait = []Trait{
	{Code: 'B', Name: "Bonus", DoubleName: "Bountiful Bonus", IncompatibleTraits: []rune{'E'}, Ingredients: []Ingredient{
		{"foodRottingFlesh", 5},
		{"medicalBloodBag", 2},
	}},
	{Code: 'U', Name: "Underground", IncompatibleTraits: []rune{'U'}, Ingredients: []Ingredient{
		{"plantedMushroom1", 1},
	}},
	{Code: 'F', Name: "Fast Growth", DoubleName: "Rapid Growth", Ingredients: []Ingredient{
		{"drinkCanMegaCrush", 2},
	}},
	{Code: 'E', Name: "Explosive", DoubleName: "Extremely Explosive", IncompatibleTraits: []rune{'B', 'R'}, Ingredients: []Ingredient{
		{"resourceScrapIron", 4},
		{"resourceGunPowder", 4},
		{"resourceDuctTape", 1},
	}},
	{Code: 'R', Name: "Renewable", IncompatibleTraits: []rune{'R', 'E'}, Ingredients: []Ingredient{
		{"drinkJarPureMineralWater", 10},
	}},
	{Code: 'T', Name: "Thorny", DoubleName: "Extra Thorny", Ingredients: []Ingredient{
		{"resourceScrapIron", 10},
		{"resourceNail", 10},
	}},
}

// IsCompatibleWith determines if one trait is compatible with another
func (t *Trait) IsCompatibleWith(other Trait) bool {
	for _, r := range t.IncompatibleTraits {
		if r == other.Code {
			return false
		}
	}
	return true
}

// GetDoubleTraitDescription returns the description for a trait that's doubled
func (t *Trait) GetDoubleTraitDescription() string {
	switch t.Code {
	case 'B':
		return fmt.Sprintf(`%s: further quadruples crop yield.`,
			t.DoubleName)
	case 'F':
		return fmt.Sprintf(`%s: reaches maturity in a quarter of the time.`,
			t.DoubleName)
	case 'E':
		return fmt.Sprintf(`%s: triggers a concealed explosive with a large payload when stepped on, struck with a melee weapon, or hit with an arrow.\n- Due to the flexible nature of plants, the detonator will not trigger if struck with bullets or other explosives.`,
			t.DoubleName)
	case 'T':
		return fmt.Sprintf(`%s: integrates with many sharp, metal thorns. Touching them will cause one to bleed profusely.`,
			t.DoubleName)
	}
	return ""
}

// GetTraitDescription returns the description for a trait
func (t *Trait) GetTraitDescription() string {
	switch t.Code {
	case 'B':
		return fmt.Sprintf(`%s: further doubles crop yield.`,
			t.Name)
	case 'U':
		return fmt.Sprintf(`%s: fused with mushroom dna, alowing growth without the need for sunlight.`,
			t.Name)
	case 'F':
		return fmt.Sprintf(`%s: reaches maturity in half the time.`,
			t.Name)
	case 'E':
		return fmt.Sprintf(`%s: triggers a concealed explosive when stepped on, struck with a melee weapon, or hit with an arrow.\n- Due to the flexible nature of plants, the detonator will not trigger if struck with bullets or other explosives.`,
			t.Name)
	case 'R':
		return fmt.Sprintf(`%s: clean, healthy water allows this plant to spread out its roots and bolster its nutrition absorption and allowing it to produce crops endlessly.`,
			t.Name)
	case 'T':
		return fmt.Sprintf(`%s: integrates with sharp, metal thorns. Touching them will cause one to bleed.`,
			t.Name)
	}
	return ""
}
