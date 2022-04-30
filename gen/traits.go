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
	{code: 'F', name: "Fast Growth", doubleName: "Rapid Growth"},
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
		return fmt.Sprintf(`%s: further quadruples crop yield.`,
			t.doubleName)
	case 'F':
		return fmt.Sprintf(`%s: reaches maturity in a quarter of the time.`,
			t.doubleName)
	case 'E':
		return fmt.Sprintf(`%s: triggers a concealed explosive with a large payload when stepped on, struck with a melee weapon, or hit with an arrow.\n- Due to the flexible nature of plants, the detonator will not trigger if struck with bullets or other explosives.`,
			t.doubleName)
	case 'T':
		return fmt.Sprintf(`%s: integrates with many sharp, metal thorns. Touching them will cause one to receive damage and bleed.`,
			t.doubleName)
	case 'S':
		return fmt.Sprintf(`%s: produces a super sweet aroma upon reaching maturity, providing the high likelihood of attracting an animal.\n- %s are especially attracted to this type of plant.`,
			t.doubleName,
			preferredConsumer)
	}
	return ""
}

func (t *Trait) getTraitDescription(preferredConsumer string) string {
	switch t.code {
	case 'B':
		return fmt.Sprintf(`%s: further doubles crop yield.`,
			t.name)
	case 'U':
		return fmt.Sprintf(`%s: fused with mushroom dna, alowing growth without the need for sunlight.`,
			t.name)
	case 'F':
		return fmt.Sprintf(`%s: reaches maturity in half the time.`,
			t.name)
	case 'E':
		return fmt.Sprintf(`%s: triggers a concealed explosive when stepped on, struck with a melee weapon, or hit with an arrow.\n- Due to the flexible nature of plants, the detonator will not trigger if struck with bullets or other explosives.`,
			t.name)
	case 'R':
		return fmt.Sprintf(`%s: clean, healthy water allows this plant to spread out its roots and bolster its nutrition absorption and allowing it to produce crops endlessly.`,
			t.name)
	case 'T':
		return fmt.Sprintf(`%s: integrates with sharp, metal thorns. Touching them will cause one to bleed.`,
			t.name)
	case 'S':
		return fmt.Sprintf(`%s: produces a sweet aroma upon reaching maturity, providing the possibility of attracting an animal.\n- %s are especially attracted to this type of plant.`,
			t.name,
			preferredConsumer)
	}
	return ""
}
