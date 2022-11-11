package gen

import (
	"data"
	"fmt"
)

// ResearcherLocalization is responsible for producing content for Localization.txt
type ResearcherLocalization struct{}

// GetPath returns file path for this producer
func (*ResearcherLocalization) GetPath() string {
	return "Config-Researcher"
}

// GetFilename returns filename for this producer
func (*ResearcherLocalization) GetFilename() string {
	return "Localization.txt"
}

// Produce xml data to the provided channel
func (p *ResearcherLocalization) Produce(c chan string) {
	defer close(c)
	c <- "Key,File,Type,english"
	p.produceJournalTipsLocalization(c)
	p.produceHotBoxLocalization(c)
	p.produceThornyBuffLocalization(c)
	for _, plant := range data.Plants {
		p.produceSchematicLocalization(c, plant)
		p.producePlantLocalization(c, plant)
		for _, trait1 := range data.Traits {
			if plant.IsCompatibleWith(trait1) {
				p.produceSchematicLocalization(c, plant, trait1)
				p.producePlantLocalization(c, plant, trait1)
				for _, trait2 := range data.Traits {
					if trait1.IsCompatibleWith(trait2) && plant.IsCompatibleWith(trait2) {
						p.produceSchematicLocalization(c, plant, trait1, trait2)
						p.producePlantLocalization(c, plant, trait1, trait2)
					}
				}
			}
		}
	}
}

func (*ResearcherLocalization) produceJournalTipsLocalization(c chan string) {
	c <- `gmoJournalTip_title,Journal Tip,,"GMO Farming [ff8000][MOD]"`
	c <- `gmoJournalTip,Journal Tip,,"Craft the [ff8000]Hot Box[-] to modify seeds with a variety of special traits: [007fff]many of which can be combined[-] and [007fff]a few can even be doubled-up to maximize their effects[-]!\n\n[00ff80]Bonus[-]: further doubles crop yield (4x yield over unmodified crops)\n\n[00ff80]Explosive[-]: [ff007f]triggers a concealed explosive when stepped on, struck with a melee weapon, or hit with an arrow[-]. Due to the flexible nature of plants, the detonator will not trigger if struck with bullets or other explosives.\n\n[00ff80]Fast[-]: plant reaches maturity in half the time.\n\n[00ff80]Renewable[-]: clean, fresh water allows this plant to spread out its roots and produce crops endlessly.\n\n[00ff80]Thorny[-]: integrates with sharp, metal thorns. [ff007f]Touching them will cause one to bleed[-].\n\n[00ff80]Underground[-]: fused with mushroom dna, allowing growth without the need for sunlight."`
}

func (*ResearcherLocalization) produceHotBoxLocalization(c chan string) {
	c <- `hotbox,blocks,Workstation,Hot Box`
	c <- `hotboxDesc,blocks,Workstation,"The Hot Box is a simple workstation that allows enhanced seeds to absorb various materials and take on new traits."`
	c <- `hotboxTip,Journal Tip,,"The Hot Box is a simple workstation that allows enhanced seeds to absorb various materials and take on new traits."`
	c <- `hotboxTip_title,Journal Tip,,Hot Box`

	c <- `lblCategoryTier1SeedResearch,UI,Tooltip,Seed Enhancement Research`
	c <- `lblCategoryTier1Seeds,UI,Tooltip,Seed Enhancement`
	c <- `lblCategoryTier2SeedResearch,UI,Tooltip,First Trait Research`
	c <- `lblCategoryTier2Seeds,UI,Tooltip,First Trait`
	c <- `lblCategoryTier3SeedResearch,UI,Tooltip,Second Trait Research`
	c <- `lblCategoryTier3Seeds,UI,Tooltip,Second Trait`

}

func (*ResearcherLocalization) produceThornyBuffLocalization(c chan string) {
	c <- `buffInjuryThornsName,buffs,Buff,Thorns`
	c <- `buffInjuryCriticalThornsName,buffs,Buff,Critical Thorns`
	c <- `buffInjuryThornsDesc,buffs,Buff,"Your skin is pierced by the thorny barbs of an aggressively engineered plant.\n\nStep away from the plant to avoid further injury."`
	c <- `buffInjuryThornsTooltip,buffs,Buff,The thorns on this plant are cutting into your skin.`
}

func (*ResearcherLocalization) producePlantLocalization(c chan string, plant data.Plant, traits ...data.Trait) {
	switch len(traits) {
	case 0:
		c <- fmt.Sprintf(`planted%s1_,blocks,Farming,"%s (Seed, Enhanced)"`,
			plant.GetName(), plant.GetDisplayName())
		c <- fmt.Sprintf(`planted%s2_,blocks,Farming,"%s (Growing, Enhanced)"`,
			plant.GetName(), plant.GetDisplayName())
		c <- fmt.Sprintf(`planted%s3_,blocks,Farming,"%s (Enhanced)"`,
			plant.GetName(), plant.GetDisplayName())
		c <- fmt.Sprintf(`planted%s1_Desc,blocks,Farming,"%s%s%s\n\n%s"`,
			plant.GetName(),
			plant.GetDescription(),
			getEnhancedSeedEffectDescription(),
			getSeedReturnDescription(traits),
			getInitialEnhancementCraftingTip())
	case 1:
		c <- fmt.Sprintf(`planted%s1_%c,blocks,Farming,"%s (Seed, %s)"`,
			plant.GetName(), traits[0].Code, plant.GetDisplayName(), traits[0].Name)
		c <- fmt.Sprintf(`planted%s2_%c,blocks,Farming,"%s (Growing, %s)"`,
			plant.GetName(), traits[0].Code, plant.GetDisplayName(), traits[0].Name)
		c <- fmt.Sprintf(`planted%s3_%c,blocks,Farming,"%s (%s)"`,
			plant.GetName(), traits[0].Code, plant.GetDisplayName(), traits[0].Name)
		c <- fmt.Sprintf(`planted%s1_%cDesc,blocks,Farming,"%s%s%s\n\n%s\n\n%s"`,
			plant.GetName(),
			traits[0].Code,
			plant.GetDescription(),
			getEnhancedSeedEffectDescription(),
			getSeedReturnDescription(traits),
			traits[0].GetTraitDescription(),
			getHotBoxRequirementTip())
	case 2:
		if traits[0].Code == traits[1].Code {
			c <- fmt.Sprintf(`planted%s1_%c%c,blocks,Farming,"%s (Seed, %s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].DoubleName)
			c <- fmt.Sprintf(`planted%s2_%c%c,blocks,Farming,"%s (Growing, %s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].DoubleName)
			c <- fmt.Sprintf(`planted%s3_%c%c,blocks,Farming,"%s (%s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].DoubleName)
			c <- fmt.Sprintf(`planted%s1_%c%cDesc,blocks,Farming,"%s%s%s\n\n%s\n\n%s"`,
				plant.GetName(),
				traits[0].Code,
				traits[1].Code,
				plant.GetDescription(),
				getEnhancedSeedEffectDescription(),
				getSeedReturnDescription(traits),
				traits[0].GetDoubleTraitDescription(),
				getHotBoxRequirementTip())
		} else {
			c <- fmt.Sprintf(`planted%s1_%c%c,blocks,Farming,"%s (Seed, %s, %s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].Name, traits[1].Name)
			c <- fmt.Sprintf(`planted%s2_%c%c,blocks,Farming,"%s (Growing, %s, %s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].Name, traits[1].Name)
			c <- fmt.Sprintf(`planted%s3_%c%c,blocks,Farming,"%s (%s, %s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].Name, traits[1].Name)
			c <- fmt.Sprintf(`planted%s1_%c%cDesc,blocks,Farming,"%s%s%s\n\n%s\n\n%s\n\n%s"`,
				plant.GetName(),
				traits[0].Code,
				traits[1].Code,
				plant.GetDescription(),
				getEnhancedSeedEffectDescription(),
				getSeedReturnDescription(traits),
				traits[0].GetTraitDescription(),
				traits[1].GetTraitDescription(),
				getHotBoxRequirementTip())
		}
	}
}

func (*ResearcherLocalization) produceSchematicLocalization(c chan string, plant data.Plant, traits ...data.Trait) {
	switch len(traits) {
	case 0:
		c <- fmt.Sprintf(`%s,blocks,Farming,"%s (Seed, Enhanced) Recipe"`,
			plant.GetSchematicName(""), plant.GetDisplayName())
	case 1:
		c <- fmt.Sprintf(`%s,blocks,Farming,"%s (Seed, %s) Recipe"`,
			plant.GetSchematicName(string(traits[0].Code)), plant.GetDisplayName(), traits[0].Name)
	case 2:
		if traits[0].Code == traits[1].Code {
			c <- fmt.Sprintf(`%s,blocks,Farming,"%s (Seed, %s) Recipe"`,
				plant.GetSchematicName(string(traits[0].Code)+string(traits[1].Code)), plant.GetDisplayName(), traits[0].DoubleName)
		} else {
			c <- fmt.Sprintf(`%s,blocks,Farming,"%s (Seed, %s, %s) Recipe"`,
				plant.GetSchematicName(string(traits[0].Code)+string(traits[1].Code)), plant.GetDisplayName(), traits[0].Name, traits[1].Name)
		}
	}
}
