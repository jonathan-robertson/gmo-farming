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
	for _, plant := range data.Plants {
		ProduceLocalization(c, plant)
		for i1 := 0; i1 < len(data.Traits); i1++ {
			if plant.IsCompatibleWith(data.Traits[i1].Code) {
				ProduceLocalization(c, plant, data.Traits[i1])
			}
			for i2 := i1; i2 < len(data.Traits); i2++ {
				if data.Traits[i1].IsCompatibleWith(data.Traits[i2]) {
					if plant.IsCompatibleWith(data.Traits[i1].Code) && plant.IsCompatibleWith(data.Traits[i2].Code) {
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
		c <- fmt.Sprintf(`planted%s1_,blocks,Farming,Enhanced %s (Seed)`,
			plant.GetName(), plant.GetDisplayName())
		c <- fmt.Sprintf(`planted%s2_,blocks,Farming,Enhanced %s (Growing)`,
			plant.GetName(), plant.GetNamePlural())
		c <- fmt.Sprintf(`planted%s3_,blocks,Farming,Enhanced %s`,
			plant.GetName(), plant.GetNamePlural())
		c <- fmt.Sprintf(`planted%s1_Desc,blocks,Farming,"%s\n\n%s\n\n%s"`,
			plant.GetName(),
			plant.GetDescription(),
			getEnhancedSeedEffectDescription(),
			getInitialEnhancementCraftingTip())
	case 1:
		c <- fmt.Sprintf(`planted%s1_%c,blocks,Farming,%s %s (Seed)`,
			plant.GetName(), traits[0].Code, traits[0].Name, plant.GetDisplayName())
		c <- fmt.Sprintf(`planted%s2_%c,blocks,Farming,%s %s (Growing)`,
			plant.GetName(), traits[0].Code, traits[0].Name, plant.GetNamePlural())
		c <- fmt.Sprintf(`planted%s3_%c,blocks,Farming,%s %s`,
			plant.GetName(), traits[0].Code, traits[0].Name, plant.GetNamePlural())
		c <- fmt.Sprintf(`planted%s1_%cDesc,blocks,Farming,"%s\n\n%s\n\n%s\n\n%s"`,
			plant.GetName(),
			traits[0].Code,
			plant.GetDescription(),
			getEnhancedSeedEffectDescription(),
			traits[0].GetTraitDescription(plant.GetPreferredConsumer()),
			getHotBoxRequirementTip())
	case 2:
		if traits[0].Code == traits[1].Code {
			c <- fmt.Sprintf(`planted%s1_%c%c,blocks,Farming,%s %s (Seed)`,
				plant.GetName(), traits[0].Code, traits[1].Code, traits[0].DoubleName, plant.GetDisplayName())
			c <- fmt.Sprintf(`planted%s2_%c%c,blocks,Farming,%s %s (Growing)`,
				plant.GetName(), traits[0].Code, traits[1].Code, traits[0].DoubleName, plant.GetNamePlural())
			c <- fmt.Sprintf(`planted%s3_%c%c,blocks,Farming,%s %s`,
				plant.GetName(), traits[0].Code, traits[1].Code, traits[0].DoubleName, plant.GetNamePlural())
			c <- fmt.Sprintf(`planted%s1_%c%cDesc,blocks,Farming,"%s\n\n%s\n\n%s\n\n%s"`,
				plant.GetName(),
				traits[0].Code,
				traits[1].Code,
				plant.GetDescription(),
				getEnhancedSeedEffectDescription(),
				traits[0].GetDoubleTraitDescription(plant.GetPreferredConsumer()),
				getHotBoxRequirementTip())
		} else {
			c <- fmt.Sprintf(`planted%s1_%c%c,blocks,Farming,"%s, %s %s (Seed)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, traits[0].Name, traits[1].Name, plant.GetDisplayName())
			c <- fmt.Sprintf(`planted%s2_%c%c,blocks,Farming,"%s, %s %s (Growing)"`,
				plant.GetName(), traits[0].Code, traits[1].Code, traits[0].Name, traits[1].Name, plant.GetNamePlural())
			c <- fmt.Sprintf(`planted%s3_%c%c,blocks,Farming,"%s, %s %s"`,
				plant.GetName(), traits[0].Code, traits[1].Code, traits[0].Name, traits[1].Name, plant.GetNamePlural())
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
	c <- `hotboxDesc,blocks,Workstation,The Hot Box is a simple workstation that allows seeds to slowly absorb the viral zombie mutagens.`
	c <- `hotboxTip,Journal Tip,,"The Hot Box is a simple workstation that allows seeds to slowly absorb the viral zombie mutagens.\n\nLeaving seeds and meat in this box will slowly attract zombies at the same rate a workbench would."`
	c <- `hotboxTip_title,Journal Tip,,Hot Box`
	c <- `perkLivingOffTheLandRank3Desc,progression,perk For,Farmer`
	c <- `perkLivingOffTheLandRank3LongDesc,progression,perk For,Triple the harvest of wild or planted crops. Craft Hot Boxes and GMO Seeds that you'll be able to research and add special traits to.`
}
