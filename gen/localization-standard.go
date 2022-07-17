package gen

import (
	"data"
	"fmt"
)

// StandardLocalization is responsible for producing content for Localization.txt
type StandardLocalization struct{}

// GetPath returns file path for this producer
func (*StandardLocalization) GetPath() string {
	return "Config-Standard"
}

// GetFilename returns filename for this producer
func (*StandardLocalization) GetFilename() string {
	return "Localization.txt"
}

// Produce xml data to the provided channel
func (p *StandardLocalization) Produce(c chan string) {
	defer close(c)
	c <- "Key,File,Type,english"
	p.produceHotBoxLocalization(c)
	p.produceThornyBuffLocalization(c)
	for _, plant := range data.Plants {
		p.producePlantLocalization(c, plant)
		for _, trait1 := range data.Traits {
			if plant.IsCompatibleWith(trait1) {
				p.producePlantLocalization(c, plant, trait1)
				for _, trait2 := range data.Traits {
					if trait1.IsCompatibleWith(trait2) && plant.IsCompatibleWith(trait2) {
						p.producePlantLocalization(c, plant, trait1, trait2)
					}
				}
			}
		}
	}
}

func (*StandardLocalization) produceHotBoxLocalization(c chan string) {
	c <- `hotbox,blocks,Workstation,Hot Box`
	c <- `hotboxDesc,blocks,Workstation,"The Hot Box is a simple workstation that allows enhanced seeds to absorb various materials and take on new traits."`
	c <- `hotboxTip,Journal Tip,,"The Hot Box is a simple workstation that allows enhanced seeds to absorb various materials and take on new traits."`
	c <- `hotboxTip_title,Journal Tip,,Hot Box`

	c <- `perkLivingOffTheLandName,progression,perk For,Living off the Land [FF8000][MOD]`
	c <- `perkLivingOffTheLandDesc,progression,perk For,Specialize in harvesting more crops using your hands or a tool.\n\n[FF8000][MOD] This perk has 2 new skill levels and added unlocks for level 3.`

	c <- `perkLivingOffTheLandRank3Desc,progression,perk For,Farmer [FF8000][MOD]`
	c <- `perkLivingOffTheLandRank3LongDesc,progression,perk For,Triple the harvest of wild or planted crops.\n\nCraft Hot Boxes and Enhanced Seeds that you'll be able to add special traits to.`
	c <- `perkLivingOffTheLandRank4Desc,progression,perk For,Mad Scientist [FF8000][MOD]`
	c <- `perkLivingOffTheLandRank4LongDesc,progression,perk For,Add a new Trait to enhanced seeds.\n\nAdded Traits can provide a wide variety of properties to a seed; ranging from increasing crop yield to allowing plants to grow without sunlight.`
	c <- `perkLivingOffTheLandRank5Desc,progression,perk For,Agricultural Genius [FF8000][MOD]`
	c <- `perkLivingOffTheLandRank5LongDesc,progression,perk For,Add a Second Trait to enhanced seeds.\n\nAdded Traits can provide a wide variety of properties to a seed; ranging from increasing crop yield to allowing plants to grow without sunlight.`

	c <- `lblCategoryTier1Seeds,UI,Tooltip,Seed Enhancement`
	c <- `lblCategoryTier2Seeds,UI,Tooltip,First Trait`
	c <- `lblCategoryTier3Seeds,UI,Tooltip,Second Trait`
}

func (*StandardLocalization) produceThornyBuffLocalization(c chan string) {
	c <- `buffInjuryThornsName,buffs,Buff,Thorns`
	c <- `buffInjuryCriticalThornsName,buffs,Buff,Critical Thorns`
	c <- `buffInjuryThornsDesc,buffs,Buff,"Your skin is pierced by the thorny barbs of an aggressively engineered plant.\n\nStep away from the plant to avoid further injury."`
	c <- `buffInjuryThornsTooltip,buffs,Buff,The thorns on this plant are cutting into your skin.`
}

func (*StandardLocalization) producePlantLocalization(c chan string, plant data.Plant, traits ...data.Trait) {
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
