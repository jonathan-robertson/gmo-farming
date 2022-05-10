package gen

import (
	"data"
	"fmt"
)

func WritePlantSchematics(target string) error {
	if target != "Vanilla" {
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
		produceSchematic(c, plant)
		for i1 := 0; i1 < len(data.Traits); i1++ {
			if plant.IsCompatibleWith(data.Traits[i1]) {
				produceSchematic(c, plant, data.Traits[i1])
			}
			for i2 := i1; i2 < len(data.Traits); i2++ {
				if data.Traits[i1].IsCompatibleWith(data.Traits[i2]) {
					if plant.IsCompatibleWith(data.Traits[i1]) && plant.IsCompatibleWith(data.Traits[i2]) {
						produceSchematic(c, plant, data.Traits[i1], data.Traits[i2])
					}
				}
			}
		}
	}
	c <- `</append></config>`
}

func produceSchematic(c chan string, p data.Plant, traits ...data.Trait) {
	var iconName string
	if p.GetName() == "GraceCorn" {
		iconName = `plantedCorn1"/><property name="CustomIconTint" value="ff9f9f`
	} else {
		iconName = fmt.Sprintf(`planted%s1`, p.GetName())
	}

	var traitsStr, group, unlockedBy string
	var unlocks []string
	switch len(traits) {
	case 0:
		traitsStr = ""
		group = "SeedExperiments"
		unlockedBy = "perkLivingOffTheLand"
		unlocks = append(unlocks, fmt.Sprintf(`planted%s1_`, p.GetName()))
		for _, trait := range data.Traits {
			if p.IsCompatibleWith(trait) {
				unlocks = append(unlocks, fmt.Sprintf(`planted%s1_%c`, p.GetName(), trait.Code))
			}
		}
	case 1:
		traitsStr = string(traits[0].Code)
		group = "SeedTraitResearch"
		unlockedBy = fmt.Sprintf(`planted%s1_schematic`, p.GetName())
		unlocks = append(unlocks, fmt.Sprintf(`planted%s1_%c`, p.GetName(), traits[0].Code))
		for _, trait := range data.Traits {
			if p.IsCompatibleWith(trait) && traits[0].IsCompatibleWith(trait) {
				unlocks = append(unlocks, fmt.Sprintf(`planted%s1_%c%c`, p.GetName(), traits[0].Code, trait.Code))
			}
		}
	case 2:
		traitsStr = string(traits[0].Code) + string(traits[1].Code)
		group = "SeedTraitResearch"
		unlockedBy = fmt.Sprintf(`planted%s1_%cschematic`, p.GetName(), traits[0].Code)
		unlocks = append(unlocks, fmt.Sprintf(`planted%s1_%c%c`, p.GetName(), traits[0].Code, traits[1].Code))
	}

	c <- fmt.Sprintf(`<item name="%s">
    <property name="Extends" value="schematicNoQualityRecipeMaster"/>
    <property name="CreativeMode" value="Player"/>
    <property name="CustomIcon" value="%s"/>
    <property name="Group" value="%s"/>
    <property name="UnlockedBy" value="%s"/>`, p.GetSchematicName(traitsStr), iconName, group, unlockedBy)
	for _, unlock := range unlocks {
		c <- fmt.Sprintf(`    <property name="Unlocks" value="%s"/>`, unlock)
	}
	c <- `    <effect_group tiered="false">`
	for _, unlock := range unlocks {
		c <- fmt.Sprintf(`    <triggered_effect trigger="onSelfPrimaryActionEnd" action="ModifyCVar" cvar="%s" operation="set" value="1"/>`, unlock)
	}
	c <- `		<triggered_effect trigger="onSelfPrimaryActionEnd" action="GiveExp" exp="50"/>
	</effect_group>
</item>`
}
