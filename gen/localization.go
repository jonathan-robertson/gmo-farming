package gen

import (
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
	for _, plant := range Plants {
		ProduceLocalization(c, plant)
		for i1 := 0; i1 < len(Traits); i1++ {
			if plant.IsCompatibleWith(Traits[i1].code) {
				ProduceLocalization(c, plant, Traits[i1])
			}
			for i2 := i1; i2 < len(Traits); i2++ {
				if Traits[i1].isCompatibleWith(Traits[i2]) {
					if plant.IsCompatibleWith(Traits[i1].code) && plant.IsCompatibleWith(Traits[i2].code) {
						ProduceLocalization(c, plant, Traits[i1], Traits[i2])
					}
				}
			}
		}
	}
}

// plantedAloe1T2U,blocks,Farming,[00FF00]T2 Aloe Vera (Seed)
// plantedAloe2T2U,blocks,Farming,[00FF00]T2 Aloe Vera (Growing)
// plantedAloe3T2U,blocks,Farming,[00FF00]T2 Aloe Vera Plant
// plantedAloe1T2U,blocks,Farming,[FF0000]T3 Aloe Vera (Seed)
// plantedAloe2T2U,blocks,Farming,[FF0000]T3 Aloe Vera (Growing)
// plantedAloe3T2U,blocks,Farming,[FF0000]T3 Aloe Vera Plant
// TODO: maybe include traits as well until fully grown? 'T3 Renewable Underground Aloe Vera Plant (Growing)'
func ProduceLocalization(c chan string, plant Plant, traits ...Trait) {
	switch len(traits) {
	case 0:
		c <- fmt.Sprintf(`planted%s1_,blocks,Farming,Upgraded %s (Seed)`,
			plant.GetName(), plant.GetDisplayName())
		c <- fmt.Sprintf(`planted%s2_,blocks,Farming,Upgraded %s (Growing)`,
			plant.GetName(), plant.GetNamePlural())
		// TODO: rename stage 3 to something more generic if players can see the name
		c <- fmt.Sprintf(`planted%s3_,blocks,Farming,Upgraded %s`,
			plant.GetName(), plant.GetNamePlural())
		c <- fmt.Sprintf(`planted%s1_Desc,blocks,Farming,"%s"`,
			plant.GetName(), plant.GetDescription())
	case 1:
		c <- fmt.Sprintf(`planted%s1_%c,blocks,Farming,%s %s (Seed)`,
			plant.GetName(), traits[0].code, traits[0].name, plant.GetDisplayName())
		c <- fmt.Sprintf(`planted%s2_%c,blocks,Farming,%s %s (Growing)`,
			plant.GetName(), traits[0].code, traits[0].name, plant.GetNamePlural())
		// TODO: rename stage 3 to something more generic if players can see the name
		c <- fmt.Sprintf(`planted%s3_%c,blocks,Farming,%s %s`,
			plant.GetName(), traits[0].code, traits[0].name, plant.GetNamePlural())
		c <- fmt.Sprintf(`planted%s1_%cDesc,blocks,Farming,"%s%s"`,
			plant.GetName(), traits[0].code, plant.GetDescription(),
			traits[0].getTraitDescription(plant.GetPreferredConsumer()))
	case 2:
		if traits[0].code == traits[1].code {
			c <- fmt.Sprintf(`planted%s1_%c%c,blocks,Farming,%s %s (Seed)`,
				plant.GetName(), traits[0].code, traits[1].code, traits[0].doubleName, plant.GetDisplayName())
			c <- fmt.Sprintf(`planted%s2_%c%c,blocks,Farming,%s %s (Growing)`,
				plant.GetName(), traits[0].code, traits[1].code, traits[0].doubleName, plant.GetNamePlural())
			// TODO: rename stage 3 to something more generic if players can see the name
			c <- fmt.Sprintf(`planted%s3_%c%c,blocks,Farming,%s %s`,
				plant.GetName(), traits[0].code, traits[1].code, traits[0].doubleName, plant.GetNamePlural())
			c <- fmt.Sprintf(`planted%s1_%c%cDesc,blocks,Farming,"%s%s"`,
				plant.GetName(), traits[0].code, traits[1].code, plant.GetDescription(),
				traits[0].getDoubleTraitDescription(plant.GetPreferredConsumer()))
		} else {
			c <- fmt.Sprintf(`planted%s1_%c%c,blocks,Farming,"%s, %s %s (Seed)"`,
				plant.GetName(), traits[0].code, traits[1].code, traits[0].name, traits[1].name, plant.GetDisplayName())
			c <- fmt.Sprintf(`planted%s2_%c%c,blocks,Farming,"%s, %s %s (Growing)"`,
				plant.GetName(), traits[0].code, traits[1].code, traits[0].name, traits[1].name, plant.GetNamePlural())
			// TODO: rename stage 3 to something more generic if players can see the name
			c <- fmt.Sprintf(`planted%s3_%c%c,blocks,Farming,"%s, %s %s"`,
				plant.GetName(), traits[0].code, traits[1].code, traits[0].name, traits[1].name, plant.GetNamePlural())
			c <- fmt.Sprintf(`planted%s1_%c%cDesc,blocks,Farming,"%s%s%s"`,
				plant.GetName(), traits[0].code, traits[1].code, plant.GetDescription(),
				traits[0].getTraitDescription(plant.GetPreferredConsumer()),
				traits[1].getTraitDescription(plant.GetPreferredConsumer()))
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
