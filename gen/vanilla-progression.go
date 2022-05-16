package gen

import (
	"data"
	"fmt"
	"strings"
)

type VanillaProgression struct{}

func (*VanillaProgression) GetPath() string {
	return "Config-Vanilla/progression.xml"
}

func (p *VanillaProgression) Produce(c chan string) {
	defer close(c)
	c <- `<config>`
	c <- `<set xpath="/progression/perks/perk[@name='perkLivingOffTheLand']/@max_level">5</set>`
	c <- `<append xpath="/progression/perks/perk[@name='perkLivingOffTheLand']/effect_group">
    <passive_effect name="RecipeTagUnlocked" operation="base_set" level="3" value="1" tags="hotbox" />`
	c <- fmt.Sprintf(`<passive_effect name="RecipeTagUnlocked" operation="base_set" level="3" value="1" tags="%s" />`,
		strings.Join(p.getTraitTagsEnhanced(), ","))
	c <- fmt.Sprintf(`<passive_effect name="RecipeTagUnlocked" operation="base_set" level="4" value="1" tags="%s" />`,
		strings.Join(p.getTraitTagsSingles(), ","))
	c <- fmt.Sprintf(`<passive_effect name="RecipeTagUnlocked" operation="base_set" level="5" value="1" tags="%s" />`,
		strings.Join(p.getTraitTagsDoubles(), ","))
	c <- `</append>`
	c <- `</config>`
}

func (*VanillaProgression) getTraitTagsEnhanced() (tags []string) {
	for _, plant := range data.Plants {
		tags = append(tags, plant.GetSchematicName(""))
	}
	return
}

func (*VanillaProgression) getTraitTagsSingles() (tags []string) {
	for _, plant := range data.Plants {
		for _, trait := range data.Traits {
			if plant.IsCompatibleWith(trait) {
				tags = append(tags, plant.GetSchematicName(string(trait.Code)))
			}
		}
	}
	return
}

func (*VanillaProgression) getTraitTagsDoubles() (tags []string) {
	for _, plant := range data.Plants {
		for _, first := range data.Traits {
			if plant.IsCompatibleWith(first) {
				for _, second := range data.Traits {
					if first.IsCompatibleWith(second) && plant.IsCompatibleWith(second) {
						tags = append(tags, plant.GetSchematicName(string(first.Code)+string(second.Code)))
					}
				}
			}
		}
	}
	return
}
