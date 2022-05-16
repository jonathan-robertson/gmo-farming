package gen

import (
	"data"
	"fmt"
	"strings"
)

// CrystalHellProgression is responsible for producing content for progression.xml
type CrystalHellProgression struct{}

// GetPath returns file path for this producer
func (*CrystalHellProgression) GetPath() string {
	return "Config-CrystalHell"
}

// GetFilename returns filename for this producer
func (*CrystalHellProgression) GetFilename() string {
	return "progression.xml"
}

// Produce xml data to the provided channel
func (p *CrystalHellProgression) Produce(c chan string) {
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

func (*CrystalHellProgression) getTraitTagsEnhanced() (tags []string) {
	for _, plant := range data.Plants {
		tags = append(tags, fmt.Sprintf("planted%s1_", plant.GetName()))
	}
	return
}

func (*CrystalHellProgression) getTraitTagsSingles() (tags []string) {
	for _, plant := range data.Plants {
		for _, trait := range data.Traits {
			if plant.IsCompatibleWith(trait) {
				tags = append(tags, fmt.Sprintf("planted%s1_%c", plant.GetName(), trait.Code))
			}
		}
	}
	return
}

func (*CrystalHellProgression) getTraitTagsDoubles() (tags []string) {
	for _, plant := range data.Plants {
		for _, first := range data.Traits {
			if plant.IsCompatibleWith(first) {
				for _, second := range data.Traits {
					if first.IsCompatibleWith(second) && plant.IsCompatibleWith(second) {
						tags = append(tags, fmt.Sprintf("planted%s1_%c%c", plant.GetName(), first.Code, second.Code))
					}
				}
			}
		}
	}
	return
}
