package gen

import (
	"data"
	"fmt"
	"strings"
)

func WriteProgression(target string) error {

	file, err := getFile(fmt.Sprintf("Config-%s/progression.xml", target))
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go produceProgression(c, target)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func produceProgression(c chan string, target string) {
	defer close(c)
	c <- `<config>`
	c <- `<set xpath="/progression/perks/perk[@name='perkLivingOffTheLand']/@max_level">5</set>`
	c <- `<append xpath="/progression/perks/perk[@name='perkLivingOffTheLand']/effect_group">
    <passive_effect name="RecipeTagUnlocked" operation="base_set" level="3" value="1" tags="hotbox" />`
	c <- fmt.Sprintf(`<passive_effect name="RecipeTagUnlocked" operation="base_set" level="3" value="1" tags="%s" />`,
		strings.Join(getTraitTagsEnhanced(target), ","))
	c <- fmt.Sprintf(`<passive_effect name="RecipeTagUnlocked" operation="base_set" level="4" value="1" tags="%s">`,
		strings.Join(getTraitTagsSingles(target), ","))
	c <- fmt.Sprintf(`<passive_effect name="RecipeTagUnlocked" operation="base_set" level="5" value="1" tags="%s">`,
		strings.Join(getTraitTagsDoubles(target), ","))
	c <- `</append>`
	c <- `</config>`
}

func getTraitTagsEnhanced(target string) (tags []string) {
	switch target {
	case "Vanilla":
		for _, plant := range data.Plants {
			tags = append(tags, plant.GetSchematicName(""))
		}
	case "CrystalHell":
		for _, plant := range data.Plants {
			tags = append(tags, fmt.Sprintf("planted%s1_", plant.GetName()))
		}
	}
	return
}

func getTraitTagsSingles(target string) (tags []string) {
	switch target {
	case "Vanilla":
		for _, plant := range data.Plants {
			for _, trait := range data.Traits {
				if plant.IsCompatibleWith(trait) {
					tags = append(tags, plant.GetSchematicName(string(trait.Code)))
				}
			}
		}
	case "CrystalHell":
		for _, plant := range data.Plants {
			for _, trait := range data.Traits {
				if plant.IsCompatibleWith(trait) {
					tags = append(tags, fmt.Sprintf("planted%s1_%c", plant.GetName(), trait.Code))
				}
			}
		}
	}
	return
}

func getTraitTagsDoubles(target string) (tags []string) {
	switch target {
	case "Vanilla":
		for _, plant := range data.Plants {
			for _, first := range data.Traits {
				if plant.IsCompatibleWith(first) {
					for _, second := range data.Traits {
						if first.IsCompatibleWith(second) && plant.IsCompatibleWith(second) {
							tags = append(tags, plant.GetSchematicName(string(first.Code)))
						}
					}
				}
			}
		}
	case "CrystalHell":
		for _, plant := range data.Plants {
			for _, first := range data.Traits {
				if plant.IsCompatibleWith(first) {
					for _, second := range data.Traits {
						if first.IsCompatibleWith(second) && plant.IsCompatibleWith(second) {
							tags = append(tags, fmt.Sprintf("planted%s1_%c%c", plant.GetName(), first.Code, second.Code))
						}
					}
				}
			}
		}
	}
	return
}
