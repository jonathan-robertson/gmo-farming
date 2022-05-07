package gen

import (
	"data"
	"fmt"
)

func WriteItems(target string) error {
	if target != "Vanilla" { // only run this for vanilla target
		return nil
	}

	file, err := getFile(fmt.Sprintf("Config-%s/items.xml", target))
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go produceItems(c)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func produceItems(c chan string) {
	defer close(c)
	c <- `<config><append xpath="/items">`
	for _, plant := range data.Plants {
		produceSchematic(c, plant, "")
		for i1 := 0; i1 < len(data.Traits); i1++ {
			if plant.IsCompatibleWith(data.Traits[i1]) {
				produceSchematic(c, plant, fmt.Sprintf("%c", data.Traits[i1].Code))
			}
			for i2 := i1; i2 < len(data.Traits); i2++ {
				if data.Traits[i1].IsCompatibleWith(data.Traits[i2]) {
					if plant.IsCompatibleWith(data.Traits[i1]) && plant.IsCompatibleWith(data.Traits[i2]) {
						produceSchematic(c, plant, fmt.Sprintf("%c%c", data.Traits[i1].Code, data.Traits[i2].Code))
					}
				}
			}
		}
	}
	c <- `</append></config>`
}

func produceSchematic(c chan string, p data.Plant, traits string) {
	optionalGraceCornIconTint := ""
	if p.GetName() == "GraceCorn" {
		optionalGraceCornIconTint = `<property name="CustomIconTint" value="ff9f9f"/>`
	}
	c <- fmt.Sprintf(`<item name="%s">
	<property name="Extends" value="schematicNoQualityRecipeMaster"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="planted%s1"/>%s
	<property name="Unlocks" value="planted%s1_%s"/>
	<effect_group tiered="false">
		<triggered_effect trigger="onSelfPrimaryActionEnd" action="ModifyCVar" cvar="planted%s1_%s" operation="set" value="1"/>
		<triggered_effect trigger="onSelfPrimaryActionEnd" action="GiveExp" exp="50"/>
	</effect_group>
</item>`, p.GetSchematicName(traits), p.GetName(), optionalGraceCornIconTint, p.GetName(), traits, p.GetName(), traits)
}
