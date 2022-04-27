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
	for _, plant := range Plants {
		for _, tier := range []int{2, 3} {
			// produce T2, T3 with no traits
			ProduceLocalizationNameEntries(c, plant, tier, "")
			ProduceLocalizationDescription(c, plant, tier)
			for i1 := 0; i1 < len(Traits); i1++ {
				switch tier {
				case 2:
					traits := fmt.Sprintf("%c", Traits[i1].code)
					if plant.IsCompatibleWith(traits) {
						ProduceLocalizationNameEntries(c, plant, tier, traits)
						ProduceLocalizationDescription(c, plant, tier, Traits[i1])
					}
				case 3:
					for i2 := i1; i2 < len(Traits); i2++ {
						if Traits[i1].isCompatibleWith(Traits[i2]) {
							traits := fmt.Sprintf("%c%c", Traits[i1].code, Traits[i2].code)
							if plant.IsCompatibleWith(traits) {
								ProduceLocalizationNameEntries(c, plant, tier, traits)
								ProduceLocalizationDescription(c, plant, tier, Traits[i1], Traits[i2])
							}
						}
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
func ProduceLocalizationNameEntries(c chan string, plant Plant, tier int, traits string) {
	switch tier {
	case 1:
		c <- fmt.Sprintf(`planted%s1,blocks,Farming,%s (T1 Seed)`, plant.GetName(), plant.GetDisplayName())
		c <- fmt.Sprintf(`planted%s2,blocks,Farming,%s (Growing)`, plant.GetName(), plant.GetNamePlural())
		c <- fmt.Sprintf(`planted%s3HarvestDesc,blocks,Farming,%s`, plant.GetName(), plant.GetNamePlural())
	default:
		c <- fmt.Sprintf(`planted%s1T%d%s,blocks,Farming,%s (T%d%s Seed)`,
			plant.GetName(), tier, traits, plant.GetDisplayName(), tier, traits)
		c <- fmt.Sprintf(`planted%s2T%d%s,blocks,Farming,%s (Growing)`,
			plant.GetName(), tier, traits, plant.GetNamePlural())
		c <- fmt.Sprintf(`planted%s3T%d%s,blocks,Farming,%s`,
			plant.GetName(), tier, traits, plant.GetNamePlural())
	}
}

func ProduceLocalizationDescription(c chan string, plant Plant, tier int, traits ...Trait) {
	switch len(traits) {
	case 0:
		c <- fmt.Sprintf(`planted%s1T%dDesc,blocks,Farming,"%s"`,
			plant.GetName(), tier, plant.GetDescription())
	case 1:
		c <- fmt.Sprintf(`planted%s1T%d%cDesc,blocks,Farming,"%s%s"`,
			plant.GetName(), tier, traits[0].code, plant.GetDescription(),
			traits[0].getTraitDescription(plant.GetPreferredConsumer()))
	case 2:
		if traits[0].code == traits[1].code {
			c <- fmt.Sprintf(`planted%s1T%d%c%cDesc,blocks,Farming,"%s%s"`,
				plant.GetName(), tier, traits[0].code, traits[1].code, plant.GetDescription(),
				traits[0].getDoubleTraitDescription(plant.GetPreferredConsumer()))
		} else {
			c <- fmt.Sprintf(`planted%s1T%d%c%cDesc,blocks,Farming,"%s%s%s"`,
				plant.GetName(), tier, traits[0].code, traits[1].code, plant.GetDescription(),
				traits[0].getTraitDescription(plant.GetPreferredConsumer()),
				traits[1].getTraitDescription(plant.GetPreferredConsumer()))
		}
	}
}
