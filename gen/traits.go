package gen

import (
	"fmt"
)

type Trait struct {
	code         rune
	name         string
	doubleName   string
	incompatible []rune
}

// 7 T2 variants * 13 seeds = 91
// 6+4+5+4+2+2+1=24 T3 variant combos * 13 seeds = 312
var Traits []Trait = []Trait{
	{code: 'B', name: "Bonus", doubleName: "Bountiful Bonus", incompatible: []rune{'E'}},
	{code: 'U', name: "Underground", incompatible: []rune{'U', 'S'}},
	{code: 'F', name: "Fast", doubleName: "Very Fast"},
	{code: 'E', name: "Explosive", doubleName: "Extremely Explosive", incompatible: []rune{'B', 'R'}},
	{code: 'R', name: "Renewable", incompatible: []rune{'R', 'E'}}, // TODO: compatible with S or not?
	{code: 'T', name: "Thorny", doubleName: "Extra Thorny"},
	{code: 'S', name: "Sweet", doubleName: "Super Sweet", incompatible: []rune{'U'}}, // TODO: compatible with R or not?
}

func (t *Trait) isCompatibleWith(t2 Trait) bool {
	for _, r := range t2.incompatible {
		if t.code == r {
			return false
		}
	}
	return true
}

func (t *Trait) getDoubleTraitDescription(preferredConsumer string) string {
	switch t.code {
	case 'B':
		return fmt.Sprintf(`\n\n%s: produces 4x the typical crop yield.`,
			t.doubleName)
	case 'F':
		return fmt.Sprintf(`\n\n%s: reaches maturity in a quarter of the time.`,
			t.doubleName)
	case 'E':
		return fmt.Sprintf(`\n\n%s: triggers a large, concealed explosive when stepped on, struck with a melee weapon, or hit with an arrow.\n- Due to the flexible nature of plants, the detonator will not trigger if struck with bullets or other explosives.`,
			t.doubleName)
	case 'T':
		return fmt.Sprintf(`\n\n%s: integrates with many sharp, metal thorns. Touching them will cause one to receive damage and bleed.`,
			t.doubleName)
	case 'S':
		return fmt.Sprintf(`\n\n%s: produces a super sweet aroma upon reaching maturity, providing the high likelihood of attracting an animal.\n- %s are especially attracted to this type of plant.`,
			t.doubleName,
			preferredConsumer)
	}
	return ""
}

func (t *Trait) getTraitDescription(preferredConsumer string) string {
	switch t.code {
	case 'B':
		return fmt.Sprintf(`\n\n%s: produces 2x the typical crop yield.`,
			t.doubleName)
	case 'U':
		return fmt.Sprintf(`\n\n%s: fused with mushroom dna, alowing growth without the need for sunlight.`,
			t.doubleName)
	case 'F':
		return fmt.Sprintf(`\n\n%s: reaches maturity in half the time.`,
			t.doubleName)
	case 'E':
		return fmt.Sprintf(`\n\n%s: triggers a concealed explosive when stepped on, struck with a melee weapon, or hit with an arrow.\n- Due to the flexible nature of plants, the detonator will not trigger if struck with bullets or other explosives.`,
			t.doubleName)
	case 'R':
		return fmt.Sprintf(`\n\n%s: clean, healthy water allows this plant to spread out its roots and bolster its nutrition absorption and allowing it to produce crops endlessly.`,
			t.doubleName)
	case 'T':
		return fmt.Sprintf(`\n\n%s: integrates with sharp, metal thorns. Touching them will cause one to bleed.`,
			t.doubleName)
	case 'S':
		return fmt.Sprintf(`\n\n%s: produces a sweet aroma upon reaching maturity, providing the possibility of attracting an animal.\n- %s are especially attracted to this type of plant.`,
			t.doubleName,
			preferredConsumer)
	}
	return ""
}
