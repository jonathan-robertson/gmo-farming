package gen

import (
	"data"
	"fmt"
)

// ResearcherItems is responsible for producing content for items.xml
type ResearcherItems struct{}

// GetPath returns file path for this producer
func (p *ResearcherItems) GetPath() string {
	return "Config-Researcher"
}

// GetFilename returns filename for this producer
func (p *ResearcherItems) GetFilename() string {
	return "items.xml"
}

// Produce xml data to the provided channel
func (p *ResearcherItems) Produce(c chan string) {
	defer close(c)
	c <- `<config><append xpath="/items">`
	for _, plant := range data.Plants {
		p.produceSchematic(c, plant)
		for _, trait1 := range data.Traits {
			if plant.IsCompatibleWith(trait1) {
				p.produceSchematic(c, plant, trait1)
				for _, trait2 := range data.Traits {
					if trait1.IsCompatibleWith(trait2) {
						if plant.IsCompatibleWith(trait1) && plant.IsCompatibleWith(trait2) {
							p.produceSchematic(c, plant, trait1, trait2)
						}
					}
				}
			}
		}
	}
	c <- `</append></config>`
}

func (p *ResearcherItems) produceSchematic(c chan string, plant data.Plant, traits ...data.Trait) {
	var iconName string
	if plant.GetName() == "GraceCorn" {
		iconName = `plantedCorn1"/><property name="CustomIconTint" value="ff9f9f`
	} else {
		iconName = fmt.Sprintf(`planted%s1`, plant.GetName())
	}

	var traitsStr, group, unlocks string
	switch len(traits) {
	case 0:
		traitsStr = ""
		group = "Tier1SeedResearch"
		unlocks = fmt.Sprintf(`planted%s1_`, plant.GetName())
	case 1:
		traitsStr = string(traits[0].Code)
		group = "Tier2SeedResearch"
		unlocks = fmt.Sprintf(`planted%s1_%c`, plant.GetName(), traits[0].Code)
	case 2:
		traitsStr = string(traits[0].Code) + string(traits[1].Code)
		group = "Tier3SeedResearch"
		unlocks = fmt.Sprintf(`planted%s1_%c%c`, plant.GetName(), traits[0].Code, traits[1].Code)
	}

	c <- fmt.Sprintf(`<item name="%s">
    <property name="Extends" value="schematicNoQualityRecipeMaster"/>
    <property name="CreativeMode" value="Player"/>
    <property name="CustomIcon" value="%s"/>
    <property name="Group" value="%s"/>
    <property name="UnlockedBy" value="perkLivingOffTheLand"/>
    <property name="Unlocks" value="%s"/>
	<effect_group>
        <triggered_effect trigger="onSelfPrimaryActionEnd" action="ModifyCVar" cvar="%s" operation="set" value="1"/>
        <triggered_effect trigger="onSelfPrimaryActionEnd" action="GiveExp" exp="50"/>	
    </effect_group>
</item>`,
		plant.GetSchematicName(traitsStr), iconName, group, unlocks, unlocks)
}
