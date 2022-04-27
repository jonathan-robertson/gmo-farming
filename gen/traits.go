package gen

import (
	"fmt"
)

type Trait struct {
	code         rune
	name         string
	incompatible []rune
}

// 7 T2 variants * 13 seeds = 91
// 6+4+5+4+2+2+1=24 T3 variant combos * 13 seeds = 312
var Traits []Trait = []Trait{
	{code: 'B', name: "Bonus", incompatible: []rune{'E'}},
	{code: 'U', name: "Underground", incompatible: []rune{'U', 'S'}},
	{code: 'F', name: "Fast"},
	{code: 'E', name: "Explosive", incompatible: []rune{'B'}},
	{code: 'R', name: "Renewable", incompatible: []rune{'R'}}, // TODO: compatible with S or not?
	{code: 'T', name: "Thorny"},
	{code: 'S', name: "Sweet", incompatible: []rune{'U'}}, // TODO: compatible with R or not?
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
		return `\n\nBountiful Bonus: produces 4x the typical crop yield.`
	case 'F':
		return `\n\nVery Fast Growth: reaches maturity in a quarter of the time.`
	case 'E':
		return `\n\nLarge Explosive: triggers a large, concealed explosive when stepped on, struck with a melee weapon, or hit with an arrow.\n- Due to the flexible nature of plants, the detonator will not trigger if struck with bullets or other explosives.`
	case 'T':
		return `\n\nExtra Thorny: integrates with many sharp, metal thorns. Touching them will cause one to receive damage and bleed.`
	case 'S':
		return fmt.Sprintf(`\n\nSuper Sweet: produces a super sweet aroma upon reaching maturity, providing the high likelihood of attracting an animal.\n- %s are especially attracted to this type of plant.`,
			preferredConsumer)
	}
	return ""
}

func (t *Trait) getTraitDescription(preferredConsumer string) string {
	switch t.code {
	case 'B':
		return `\n\nBonus: produces 2x the typical crop yield.`
	case 'U':
		return `\n\nUnderground: fused with mushroom dna, alowing growth without the need for sunlight.`
	case 'F':
		return `\n\nFast Growth: reaches maturity in half the time.`
	case 'E':
		return `\n\nExplosive: triggers a concealed explosive when stepped on, struck with a melee weapon, or hit with an arrow.\n- Due to the flexible nature of plants, the detonator will not trigger if struck with bullets or other explosives.`
	case 'R':
		return `\n\nRenewable: clean, healthy water allows this plant to spread out its roots and bolster its nutrition absorption and allowing it to produce crops endlessly.`
	case 'T':
		return `\n\nThorny: integrates with sharp, metal thorns. Touching them will cause one to bleed.`
	case 'S':
		return fmt.Sprintf(`\n\nSweet: produces a sweet aroma upon reaching maturity, providing the possibility of attracting an animal.\n- %s are especially attracted to this type of plant.`,
			preferredConsumer)
	}
	return ""
}
