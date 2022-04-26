package main

import (
	"fmt"
	"gen"
	"os"
)

type variant struct {
	code         rune
	name         string
	incompatible []rune
}

/*
var vanillaStages []string = []string{
	"1",
	"2",
	"3HarvestPlayer",
}

var variantStages []string = []string{
	"1",
	"2",
	"3",
}
*/
var variants []variant = []variant{
	{code: 'B', name: "Bonus", incompatible: []rune{'E'}},
	{code: 'U', name: "Underground", incompatible: []rune{'U', 'S'}},
	{code: 'F', name: "Fast"},
	{code: 'E', name: "Explosive", incompatible: []rune{'B'}},
	{code: 'R', name: "Renewable", incompatible: []rune{'R'}},
	{code: 'T', name: "Thorny"},
	{code: 'S', name: "Sweet", incompatible: []rune{'U'}},
}

func (v *variant) isCompatibleWith(v2 variant) bool {
	for _, r := range v2.incompatible {
		if v.code == r {
			return false
		}
	}
	return true
}

/*
// Write all 3 stages to file
func produceBlocks(c chan string, b gen.Block, tier int, traits string) (err error) {
	for n, stage := range variantStages {
		c <- fmt.Sprintf(`        <block name="%s%sT%d%s" stage="%s" traits="%s">`, b.name, stage, tier, traits, stage, traits)
		switch stage {
		case "1":
			c <- fmt.Sprintf(`            <property name="Extends" value="%s%s" />`, b.name, vanillaStages[n])
			c <- fmt.Sprintf(`            <property name="CustomIcon" value="%s%s" />`, b.name, vanillaStages[n])
			c <- fmt.Sprintf(`            <property name="PlantGrowing.Next" value="%s%sT%d%s" />`, b.name, variantStages[n+1], tier, traits)
			c <- fmt.Sprintf(`            <drop event="Destroy" name="%s%sT%d%s" count="1" />`, b.name, stage, tier, traits)
		case "2":
			c <- fmt.Sprintf(`            <property name="Extends" value="%s1T%d%s" />`, b.name, tier, traits)
			c <- fmt.Sprintf(`            <property name="PlantGrowing.Next" value="%s%sT%d%s" />`, b.name, variantStages[n+1], tier, traits)
			c <- fmt.Sprintf(`            <drop event="Destroy" name="%s1T%d%s" count="1" />`, b.name, tier, traits)
		case "3":
			c <- fmt.Sprintf(`            <property name="Extends" value="%s%s" />`, b.name, vanillaStages[n])
			c <- fmt.Sprintf(`            <drop event="Harvest" name="%s" count="4" tag="cropHarvest" />`, b.crop)
			c <- fmt.Sprintf(`            <drop event="Harvest" name="%s" prob="0.5" count="2" tag="bonusCropHarvest" />`, b.crop)
			c <- fmt.Sprintf(`            <drop event="Destroy" name="%s1T%d%s" count="1" prob="0.5" />`, b.name, tier, traits)
			if strings.ContainsRune(traits, 'R') {
				c <- fmt.Sprintf(`            <property name="DowngradeBlock" value="%s1T%d%s" />`, b.name, tier, traits)
			}
		}
		c <- "        </block>"
	}
	return
}

func produceVariant(c chan string, p plant, tier int, traits string) (err error) {
	if err := produceBlocks(c, p, tier, traits); err != nil {
		return err
	}
	// TODO: produce recipe
	// TODO: produce localization
	return nil
}

func produceVariants(c chan string) {
	defer close(c)
	c <- "<config>\n    <append xpath=\"/blocks\">"
	for _, plant := range plants {
		for _, tier := range []int{2, 3} {
			produceVariant(c, plant, tier, "")
			for i1 := 0; i1 < len(variants); i1++ {
				traits := fmt.Sprintf("%c", variants[i1].code)
				produceVariant(c, plant, tier, traits)
				if tier == 3 {
					for i2 := i1; i2 < len(variants); i2++ {
						if variants[i1].isCompatibleWith(variants[i2]) {
							traits = fmt.Sprintf("%s%c", traits, variants[i2].code)
							produceVariant(c, plant, tier, traits)
						}
					}
				}
			}
		}
	}
	c <- "    </append>"
	produceBlockModifications(c)
	c <- "</config>"
}

func writeBlocks() error {
	file, err := getFile("Config/blocks.xml")
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go produceVariants(c)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}
*/

/*
func producePlantLocalization(c chan string) {

}
func writePlantLocalization() error {
	file, err := getFile("Config/localization.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go producePlantRecipes(c)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func producePlantRecipes(c chan string) {

}
func writePlantRecipes() error {
	file, err := getFile("Config/recipes.xml")
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go producePlantRecipes(c)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}
*/

func produceBlockModifications(c chan string) {
	// {code: 'U', name: "Underground", incompatible: []rune{'U', 'S'}},
	c <- `    <append xpath="/blocks/block[contains(@traits, 'U') and @stage='1']">
        <property name="PlantGrowing.LightLevelGrow" value="0" />
        <property name="PlantGrowing.LightLevelStay" value="0" />
    </append>`

	// {code: 'F', name: "Fast"},
	c <- `    <append xpath="/blocks/block[contains(@traits, 'F') and @stage='1' and not (@traits='FF')]">
        <property name="PlantGrowing.GrowthRate" value="31.5" />
    </append>`
	c <- `    <append xpath="/blocks/block[@traits='FF' and @stage='1']">
        <property name="PlantGrowing.GrowthRate" value="15.75" />
    </append>`

	// {code: 'E', name: "Explosive", incompatible: []rune{'E'}},
	// based off of mineCookingPot with reduced trigger delay from .5 to .1
	// TODO: <property name="Material" value="MLandMine" />
	c <- `    <append xpath="/blocks/block[contains(@traits, 'E') and @stage='3' and not (@traits='EE')]">
        <property name="Class" value="Mine" /> <!-- a mine destroyed by an *explosion* only has a 33 percent chance to detonate -->
        <property name="Tags" value="Mine" />
        <property name="Collide" value="movement,melee,arrow" />
		<property name="MaxDamage" value="1" /> <!-- reduced from 4 -->
        <property name="TriggerDelay" value="0.1" /> <!-- reduced from 0.5 -->
        <property name="TriggerSound" value="landmine_trigger" />
        <property name="Explosion.ParticleIndex" value="11" />
        <property name="Explosion.RadiusEntities" value="3" />
        <property name="Explosion.EntityDamage" value="300" /> <!-- damage for entities in the center of the explosion. -->
        <property name="CanPickup" value="false" />
    </append>`
	// based off of mineHubcap with reduced trigger delay from .5 to .1
	// TODO: <property name="Material" value="MLandMine" />
	c <- `    <append xpath="/blocks/block[contains(@traits, 'EE') and @stage='3']">
        <property name="Class" value="Mine" /> <!-- a mine destroyed by an *explosion* only has a 33 percent chance to detonate -->
        <property name="Tags" value="Mine" />
        <property name="Collide" value="movement,melee,arrow" />
		<property name="MaxDamage" value="1" /> <!-- reduced from 4 -->
        <property name="TriggerDelay" value="0.1" /> <!-- reduced from 0.5 -->
        <property name="TriggerSound" value="landmine_trigger" />
        <property name="Explosion.ParticleIndex" value="11" />
        <property name="Explosion.RadiusEntities" value="5" />
        <property name="Explosion.EntityDamage" value="450" /> <!-- damage for entities in the center of the explosion. -->
        <property name="CanPickup" value="false" />
    </append>`

	// {code: 'T', name: "Thorny"},
	// {code: 'S', name: "Sweet", incompatible: []rune{'U'}},

	// LOCALIZED: {code: 'B', name: "Bonus", incompatible: []rune{'B'}},
	// LOCALIZED: {code: 'R', name: "Renewable", incompatible: []rune{'R'}},
}

func producePlantBlocks(c chan string) {
	defer close(c)
	c <- `<config>`
	c <- `<append xpath="/blocks">`
	for _, plant := range gen.Plants {
		for _, tier := range []int{2, 3} {
			// produce T2, T3 with no traits
			plant.WriteStages(c, tier, "")
			for i1 := 0; i1 < len(variants); i1++ {
				switch tier {
				case 2:
					traits := fmt.Sprintf("%c", variants[i1].code)
					if plant.IsCompatibleWith(traits) {
						plant.WriteStages(c, tier, traits)
					}
				case 3:
					for i2 := i1; i2 < len(variants); i2++ {
						if variants[i1].isCompatibleWith(variants[i2]) {
							traits := fmt.Sprintf("%c%c", variants[i1].code, variants[i2].code)
							if plant.IsCompatibleWith(traits) {
								plant.WriteStages(c, tier, traits)
							}
						}
					}
				}
			}
		}
	}
	c <- `</append>`
	produceBlockModifications(c)
	c <- `</config>`
}

func writePlantBlocks() error {
	file, err := getFile("Config/blocks.xml")
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go producePlantBlocks(c)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func getFile(filename string) (*os.File, error) {
	os.Remove(filename)
	return os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
}

func main() {
	// printVanillaStages()
	// printVariantStages()
	// writeBlocksFile()

	if err := writePlantBlocks(); err != nil {
		panic(err)
	}
}
