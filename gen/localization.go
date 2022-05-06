package gen

import (
	"data"
	"fmt"
)

func WritePlantLocalization() error {
	file, err := getFile("Config/Localization.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go producePlantLocalization(c)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func producePlantLocalization(c chan string) {
	defer close(c)
	c <- "Key,File,Type,english"
	ProduceHotBoxLocalization(c)
	ProduceThornyBuffLocalization(c)
	for _, plant := range data.Plants {
		ProduceLocalization(c, plant)
		for i1 := 0; i1 < len(data.Traits); i1++ {
			if plant.IsCompatibleWith(data.Traits[i1]) {
				ProduceLocalization(c, plant, data.Traits[i1])
			}
			for i2 := i1; i2 < len(data.Traits); i2++ {
				if data.Traits[i1].IsCompatibleWith(data.Traits[i2]) {
					if plant.IsCompatibleWith(data.Traits[i1]) && plant.IsCompatibleWith(data.Traits[i2]) {
						ProduceLocalization(c, plant, data.Traits[i1], data.Traits[i2])
					}
				}
			}
		}
	}
}

func ProduceLocalization(c chan string, plant data.Plant, traits ...data.Trait) {
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

func ProduceHotBoxLocalization(c chan string) {
	c <- `hotbox,blocks,Workstation,Hot Box`
	c <- `hotboxDesc,blocks,Workstation,The Hot Box is a simple workstation that allows enhanced seeds to absorb various materials and take on new traits.`
	c <- `hotboxTip,Journal Tip,,"The Hot Box is a simple workstation that allows enhanced seeds to absorb various materials and take on new traits."`
	c <- `hotboxTip_title,Journal Tip,,Hot Box`
	c <- `perkLivingOffTheLandRank3Desc,progression,perk For,Farmer`
	c <- `perkLivingOffTheLandRank3LongDesc,progression,perk For,Triple the harvest of wild or planted crops. Craft Hot Boxes and Enhanced Seeds that you'll be able to research and add special traits to.`
	c <- `lblCategorySeedEnhancement,UI,Tooltip,Seed Enhancement`
}

func ProduceThornyBuffLocalization(c chan string) {
	c <- `buffInjuryThornsName,buffs,Buff,Thorns`
	c <- `buffInjuryCriticalThornsName,buffs,Buff,Critical Thorns`
	c <- `buffInjuryThornsDesc,buffs,Buff,"Your skin is pierced by the thorny barbs of an aggressively engineered plant.\n\nStep away from the plant to avoid further injury."`
	c <- `buffInjuryThornsTooltip,buffs,Buff,The thorns on this plant are cutting into your skin.`
}
