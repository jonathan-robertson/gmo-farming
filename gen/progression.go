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
	c <- `<config><append xpath="/progression/perks/perk[@name='perkLivingOffTheLand']/effect_group">
    <passive_effect name="RecipeTagUnlocked" operation="base_set" level="3" value="1" tags="hotbox" />`

	if target != "Vanilla" {
		c <- `</append></config>`
		return
	}

	var tags []string
	for _, plant := range data.Plants {
		tags = append(tags, plant.GetSchematicName(""))
	}
	c <- fmt.Sprintf(`    <passive_effect name="RecipeTagUnlocked" operation="base_set" level="3" value="1" tags="%s" />`,
		strings.Join(tags, ","))
	c <- `</append></config>`
}
