package gen

import (
	"data"
	"fmt"
)

type CrystalHellLocalization struct{}

func (*CrystalHellLocalization) GetPath() string {
	return "Config-CrystalHell"
}

func (*CrystalHellLocalization) GetFilename() string {
	return "Localization.txt"
}

func (p *CrystalHellLocalization) Produce(c chan string) {
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

func (*CrystalHellLocalization) produceHotBoxLocalization(c chan string) {
	c <- `hotbox,blocks,Workstation,Hot Box`
	c <- `hotboxDesc,blocks,Workstation,"The Hot Box is a simple workstation that allows enhanced seeds to absorb various materials and take on new traits."`
	c <- `hotboxTip,Journal Tip,,"The Hot Box is a simple workstation that allows enhanced seeds to absorb various materials and take on new traits."`
	c <- `hotboxTip_title,Journal Tip,,Hot Box`

	c <- `perkLivingOffTheLandRank3Desc,progression,perk For,Farmer`
	c <- `perkLivingOffTheLandRank3LongDesc,progression,perk For,Triple the harvest of wild or planted crops. Craft Hot Boxes and Enhanced Seeds that you'll be able to research special traits for.`
	c <- `perkLivingOffTheLandRank4Desc,progression,perk For,Mad Scientist`
	c <- `perkLivingOffTheLandRank4LongDesc,progression,perk For,Craft a Trait into enhanced seeds.\n\nTraits can be used to add a wide variety of properties to a seed; ranging from increasing crop yield to allowing plants to grow without sunlight.`
	c <- `perkLivingOffTheLandRank5Desc,progression,perk For,Agricultural Genius`
	c <- `perkLivingOffTheLandRank5LongDesc,progression,perk For,Craft a second Trait into enhanced seeds.\n\nDouble the Traits,\nDouble the fun!`

	c <- `lblCategoryTier1Seeds,UI,Tooltip,Tier 1 Seed Enhancements`
	c <- `lblCategoryTier2Seeds,UI,Tooltip,Tier 2 Seed Enhancements`
	c <- `lblCategoryTier3Seeds,UI,Tooltip,Tier 3 Seed Enhancements`
}

func (*CrystalHellLocalization) produceThornyBuffLocalization(c chan string) {
	c <- `buffInjuryThornsName,buffs,Buff,Thorns`
	c <- `buffInjuryCriticalThornsName,buffs,Buff,Critical Thorns`
	c <- `buffInjuryThornsDesc,buffs,Buff,"Your skin is pierced by the thorny barbs of an aggressively engineered plant.\n\nStep away from the plant to avoid further injury."`
	c <- `buffInjuryThornsTooltip,buffs,Buff,The thorns on this plant are cutting into your skin.`
}

func (*CrystalHellLocalization) producePlantLocalization(c chan string, plant data.Plant, traits ...data.Trait) {
	switch len(traits) {
	case 0:
		c <- fmt.Sprintf(`planted%s1_,blocks,Farming,"%s (Seed, Enhanced)"`,
			plant.GetName(), plant.GetDisplayName())
		c <- fmt.Sprintf(`planted%s2_,blocks,Farming,"%s (Growing, Enhanced)"`,
			plant.GetName(), plant.GetDisplayName())
		c <- fmt.Sprintf(`planted%s3_,blocks,Farming,"%s (Enhanced)"`,
			plant.GetName(), plant.GetDisplayName())
		c <- fmt.Sprintf(`planted%s1_Desc,blocks,Farming,"%s\n\n%s\n\n%s"`,
			plant.GetName(),
			plant.GetDescription(),
			getEnhancedSeedEffectDescription(),
			getInitialEnhancementCraftingTip())
	case 1:
		c <- fmt.Sprintf(`planted%s1_%c,blocks,Farming,"%s (Seed, %s)"`,
			plant.GetName(), traits[0].Code, plant.GetDisplayName(), traits[0].Name)
		c <- fmt.Sprintf(`planted%s2_%c,blocks,Farming,"%s (Growing, %s)"`,
			plant.GetName(), traits[0].Code, plant.GetDisplayName(), traits[0].Name)
		c <- fmt.Sprintf(`planted%s3_%c,blocks,Farming,"%s (%s)"`,
			plant.GetName(), traits[0].Code, plant.GetDisplayName(), traits[0].Name)
		c <- fmt.Sprintf(`planted%s1_%cDesc,blocks,Farming,"%s\n\n%s\n\n%s\n\n%s"`,
			plant.GetName(),
			traits[0].Code,
			plant.GetDescription(),
			getEnhancedSeedEffectDescription(),
			traits[0].GetTraitDescription(plant.GetPreferredConsumer()),
			getHotBoxRequirementTip())
	case 2:
		if traits[0].Code == traits[1].Code {
			c <- fmt.Sprintf(`planted%s1_%c%c,blocks,Farming,"%s (Seed, %s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].DoubleName)
			c <- fmt.Sprintf(`planted%s2_%c%c,blocks,Farming,"%s (Growing, %s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].DoubleName)
			c <- fmt.Sprintf(`planted%s3_%c%c,blocks,Farming,"%s (%s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].DoubleName)
			c <- fmt.Sprintf(`planted%s1_%c%cDesc,blocks,Farming,"%s\n\n%s\n\n%s\n\n%s"`,
				plant.GetName(),
				traits[0].Code,
				traits[1].Code,
				plant.GetDescription(),
				getEnhancedSeedEffectDescription(),
				traits[0].GetDoubleTraitDescription(plant.GetPreferredConsumer()),
				getHotBoxRequirementTip())
		} else {
			c <- fmt.Sprintf(`planted%s1_%c%c,blocks,Farming,"%s (Seed, %s, %s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].Name, traits[1].Name)
			c <- fmt.Sprintf(`planted%s2_%c%c,blocks,Farming,"%s (Growing, %s, %s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].Name, traits[1].Name)
			c <- fmt.Sprintf(`planted%s3_%c%c,blocks,Farming,"%s (%s, %s)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, plant.GetDisplayName(), traits[0].Name, traits[1].Name)
			c <- fmt.Sprintf(`planted%s1_%c%cDesc,blocks,Farming,"%s\n\n%s\n\n%s\n\n%s\n\n%s"`,
				plant.GetName(),
				traits[0].Code,
				traits[1].Code,
				plant.GetDescription(),
				getEnhancedSeedEffectDescription(),
				traits[0].GetTraitDescription(plant.GetPreferredConsumer()),
				traits[1].GetTraitDescription(plant.GetPreferredConsumer()),
				getHotBoxRequirementTip())
		}
	}
}