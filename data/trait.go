package data

import (
	"fmt"
)

type Trait struct {
	Code         rune
	Name         string
	DoubleName   string
	Incompatible []rune
	Ingredients  []Ingredient
}

type Ingredient struct {
	Name  string
	Count int
}

var Traits []Trait = []Trait{
	{Code: 'B', Name: "Bonus", DoubleName: "Bountiful Bonus", Incompatible: []rune{'E'}, Ingredients: []Ingredient{
		{"foodRottingFlesh", 5},
		{"medicalBloodBag", 2},
	}},
	{Code: 'U', Name: "Underground", Incompatible: []rune{'U', 'S'}, Ingredients: []Ingredient{
		{"plantedMushroom1", 1},
	}},
	{Code: 'F', Name: "Fast Growth", DoubleName: "Rapid Growth", Ingredients: []Ingredient{
		{"drinkCanMegaCrush", 2},
	}},
	{Code: 'E', Name: "Explosive", DoubleName: "Extremely Explosive", Incompatible: []rune{'B', 'R'}, Ingredients: []Ingredient{
		{"resourceScrapIron", 4},
		{"resourceGunPowder", 4},
		{"resourceDuctTape", 1},
	}},
	{Code: 'R', Name: "Renewable", Incompatible: []rune{'R', 'E'}, Ingredients: []Ingredient{
		{"drinkJarPureMineralWater", 10},
	}}, // TODO: compatible with S or not?
	/* TODO
	{Code: 'T', Name: "Thorny", DoubleName: "Extra Thorny", Ingredients: []Ingredient{
		{"resourceScrapIron", 10},
		{"resourceNail", 10},
	}},
	{Code: 'S', Name: "Sweet", DoubleName: "Super Sweet", Incompatible: []rune{'U'}, Ingredients: []Ingredient{
		{"resourceTestosteroneExtract", 2},
	}}, // TODO: compatible with R or not?
	*/
}

func (t *Trait) IsCompatibleWith(other Trait) bool {
	for _, r := range other.Incompatible {
		if t.Code == r {
			return false
		}
	}
	return true
}

func (t *Trait) GetDoubleTraitDescription(preferredConsumer string) string {
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
		return fmt.Sprintf(`%s: integrates with many sharp, metal thorns. Touching them will cause one to receive damage and bleed.`,
			t.DoubleName)
	case 'S':
		return fmt.Sprintf(`%s: produces a super sweet aroma upon reaching maturity, providing the high likelihood of attracting an animal.\n- %s are especially attracted to this type of plant.`,
			t.DoubleName,
			preferredConsumer)
	}
	return ""
}

func (t *Trait) GetTraitDescription(preferredConsumer string) string {
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
	case 'S':
		return fmt.Sprintf(`%s: produces a sweet aroma upon reaching maturity, providing the possibility of attracting an animal.\n- %s are especially attracted to this type of plant.`,
			t.Name,
			preferredConsumer)
	}
	return ""
}
