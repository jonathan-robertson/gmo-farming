package gen

import (
	"data"
	"fmt"
	"strings"
)

// ResearcherProgression is responsible for producing content for progression.xml
type ResearcherProgression struct{}

// GetPath returns file path for this producer
func (*ResearcherProgression) GetPath() string {
	return "Config-Researcher"
}

// GetFilename returns filename for this producer
func (*ResearcherProgression) GetFilename() string {
	return "progression.xml"
}

// Produce xml data to the provided channel
func (p *ResearcherProgression) Produce(c chan string) {
	defer close(c)
	c <- `<config>`
	c <- `<set xpath="/progression/perks/perk[@name='perkLivingOffTheLand']/@max_level">5</set>`
	c <- `<set xpath="/progression/perks/perk[@name='perkLivingOffTheLand']/effect_group/passive_effect[@name='HarvestCount' and @tags='cropHarvest,wildCropsHarvest']/@level">1,2,3,4,5</set>`
	c <- `<set xpath="/progression/perks/perk[@name='perkLivingOffTheLand']/effect_group/passive_effect[@name='HarvestCount' and @tags='cropHarvest,wildCropsHarvest']/@value">1,1,2,2,2</set>`
	c <- `<set xpath="/progression/perks/perk[@name='perkLivingOffTheLand']/effect_group/passive_effect[@name='HarvestCount' and @tags='bonusCropHarvest']/@level">2,3,4,5</set>`
	c <- `<set xpath="/progression/perks/perk[@name='perkLivingOffTheLand']/effect_group/passive_effect[@name='HarvestCount' and @tags='bonusCropHarvest']/@value">1,1,1,1</set>`
	c <- `<append xpath="/progression/perks/perk[@name='perkLivingOffTheLand']/effect_group">
    <passive_effect name="RecipeTagUnlocked" operation="base_set" level="3,5" value="1" tags="hotbox" />`
	c <- fmt.Sprintf(`<passive_effect name="RecipeTagUnlocked" operation="base_set" level="3,5" value="1" tags="%s" />`,
		strings.Join(p.getTraitTagsEnhanced(), ","))
	c <- fmt.Sprintf(`<passive_effect name="RecipeTagUnlocked" operation="base_set" level="4,5" value="1" tags="%s" />`,
		strings.Join(p.getTraitTagsSingles(), ","))
	c <- fmt.Sprintf(`<passive_effect name="RecipeTagUnlocked" operation="base_set" level="5,5" value="1" tags="%s" />`,
		strings.Join(p.getTraitTagsDoubles(), ","))
	c <- `</append>`
	c <- `</config>`
}

func (*ResearcherProgression) getTraitTagsEnhanced() (tags []string) {
	for _, plant := range data.Plants {
		tags = append(tags, plant.GetSchematicName(""))
	}
	return
}

func (*ResearcherProgression) getTraitTagsSingles() (tags []string) {
	for _, plant := range data.Plants {
		for _, trait := range data.Traits {
			if plant.IsCompatibleWith(trait) {
				tags = append(tags, plant.GetSchematicName(string(trait.Code)))
			}
		}
	}
	return
}

func (*ResearcherProgression) getTraitTagsDoubles() (tags []string) {
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
