package gen

import (
	"data"
	"fmt"
)

// ResearcherItems is responsible for producing content for items.xml
type ResearcherRewardsLoot struct{}

// GetPath returns file path for this producer
func (p *ResearcherRewardsLoot) GetPath() string {
	return "Config-Researcher-Rewards"
}

// GetFilename returns filename for this producer
func (p *ResearcherRewardsLoot) GetFilename() string {
	return "loot.xml"
}

// Produce xml data to the provided channel
func (p *ResearcherRewardsLoot) Produce(c chan string) {
	defer close(c)
	c <- `<config><insertBefore xpath="/lootcontainers/lootgroup[1]">`
	p.ProduceEnhancedSeedSchematics(c)
	p.ProduceSingleTraitSeedSchematics(c)
	p.ProduceDoubleTraitSeedSchematics(c)
	c <- `
		<lootgroup name="groupReinforcedChestT2_GMO" count="all">
			<item group="enhancedSeedSchematics" prob=".01" force_prob="true"/>
		</lootgroup>
		<lootgroup name="groupReinforcedChestT3_GMO" count="all">
			<item group="enhancedSeedSchematics" prob=".05" force_prob="true"/>
			<item group="singleTraitSeedSchematics" prob=".01" force_prob="true"/>
		</lootgroup>
		<lootgroup name="groupHardenedChestT4_GMO" count="all">
			<item group="enhancedSeedSchematics" prob=".15" force_prob="true"/>
			<item group="singleTraitSeedSchematics" prob=".05" force_prob="true"/>
			<item group="doubleTraitSeedSchematics" prob=".01" force_prob="true"/>
		</lootgroup>
		<lootgroup name="groupHardenedChestT5_GMO" count="all">
			<item group="enhancedSeedSchematics" prob=".25" force_prob="true"/>
			<item group="singleTraitSeedSchematics" prob=".15" force_prob="true"/>
			<item group="doubleTraitSeedSchematics" prob=".05" force_prob="true"/>
		</lootgroup>
	</insertBefore>

	<append xpath="/lootcontainers/lootgroup[@name='groupReinforcedChestT2']">
		<item group="groupReinforcedChestT2_GMO"/>
	</append>
	<append xpath="/lootcontainers/lootgroup[@name='groupReinforcedChestT3']">
		<item group="groupReinforcedChestT3_GMO"/>
	</append>
	<append xpath="/lootcontainers/lootgroup[@name='groupHardenedChestT4']">
		<item group="groupHardenedChestT4_GMO"/>
	</append>
	<append xpath="/lootcontainers/lootgroup[@name='groupHardenedChestT5']">
		<item group="groupHardenedChestT5_GMO"/>
	</append>
</config>`
}

func (p *ResearcherRewardsLoot) ProduceEnhancedSeedSchematics(c chan string) {
	c <- `<lootgroup name="enhancedSeedSchematics">`
	for _, plant := range data.Plants {
		p.produceSchematic(c, plant)
	}
	c <- `</lootgroup>`
}

func (p *ResearcherRewardsLoot) ProduceSingleTraitSeedSchematics(c chan string) {
	c <- `<lootgroup name="singleTraitSeedSchematics">`
	for _, plant := range data.Plants {
		for _, trait1 := range data.Traits {
			if plant.IsCompatibleWith(trait1) {
				p.produceSchematic(c, plant, trait1)
			}
		}
	}
	c <- `</lootgroup>`
}

func (p *ResearcherRewardsLoot) ProduceDoubleTraitSeedSchematics(c chan string) {
	c <- `<lootgroup name="doubleTraitSeedSchematics">`
	for _, plant := range data.Plants {
		for _, trait1 := range data.Traits {
			if plant.IsCompatibleWith(trait1) {
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
	c <- `</lootgroup>`
}

func (p *ResearcherRewardsLoot) produceSchematic(c chan string, plant data.Plant, traits ...data.Trait) {
	var traitsStr string
	switch len(traits) {
	case 0:
		traitsStr = ""
	case 1:
		traitsStr = string(traits[0].Code)
	case 2:
		traitsStr = string(traits[0].Code) + string(traits[1].Code)
	}
	c <- fmt.Sprintf(`<item name="%s" />`, plant.GetSchematicName(traitsStr))
}
