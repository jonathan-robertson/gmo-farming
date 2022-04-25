package main

import (
	"fmt"
	"os"
	"strings"
)

type plant struct {
	name string
	crop string
}
type tier struct {
	text string
}
type variant struct {
	code         rune
	name         string
	incompatible []rune
}

var plants []plant
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
var variantTiers []string = []string{
	"T1",
	"T2",
	"T3",
}
var variants []variant

func (v *variant) isCompatibleWith(v2 variant) bool {
	for _, r := range v2.incompatible {
		if v.code == r {
			return false
		}
	}
	return true
}

func init() {
	names := []string{
		"Mushroom",
		"Yucca",
		"Cotton",
		"Coffee",
		"Goldenrod",
		"Aloe",
		"Blueberry",
		"Potato",
		"Chrysanthemum",
		"Corn",
		"GraceCorn",
		"Hop",
		"Pumpkin",
	}
	for _, name := range names {
		plants = append(plants, plant{name: "planted" + name, crop: "foodCrop" + name})
	}

	variants = []variant{
		{code: 'B', name: "Bonus", incompatible: []rune{'E'}},
		{code: 'U', name: "Underground", incompatible: []rune{'U', 'S'}},
		{code: 'F', name: "Fast"},
		{code: 'E', name: "Explosive", incompatible: []rune{'B'}},
		{code: 'R', name: "Renewable", incompatible: []rune{'R'}},
		{code: 'T', name: "Thorny"},
		{code: 'S', name: "Sweet", incompatible: []rune{'U'}},
	}
}

// Write all 3 stages to file
func produceBlock(c chan string, p plant, tier int, traits ...rune) (err error) {
	for n, stage := range variantStages {
		var traitSuffix, traitTag string
		switch len(traits) {
		case 0:
			traitSuffix = ""
			traitTag = ""
		case 1:
			traitSuffix = fmt.Sprintf("%c", traits[0])
			traitTag = fmt.Sprintf("%c", traits[0])
		case 2:
			traitSuffix = fmt.Sprintf("%c%c", traits[0], traits[1])
			traitTag = fmt.Sprintf("%c,%c", traits[0], traits[1])
		default:
			return fmt.Errorf("received too many traits")
		}
		c <- fmt.Sprintf(`        <block name="%s%sT%d%s" stage="%s" trait="%s">`, p.name, stage, tier, traitSuffix, stage, traitTag)
		switch stage {
		case "1":
			c <- fmt.Sprintf(`            <property name="Extends" value="%s%s" />`, p.name, vanillaStages[n])
			c <- fmt.Sprintf(`            <property name="CustomIcon" value="%s%s" />`, p.name, vanillaStages[n])
			c <- fmt.Sprintf(`            <property name="PlantGrowing.Next" value="%s%sT%d%s" />`, p.name, variantStages[n+1], tier, traitSuffix)
			c <- fmt.Sprintf(`            <drop event="Destroy" name="%s%sT%d%s" count="1" />`, p.name, stage, tier, traitSuffix)
		case "2":
			c <- fmt.Sprintf(`            <property name="Extends" value="%s%sT%d%s" />`, p.name, stage, tier, traitSuffix)
			c <- fmt.Sprintf(`            <property name="PlantGrowing.Next" value="%s%sT%d%s" />`, p.name, variantStages[n+1], tier, traitSuffix)
			c <- fmt.Sprintf(`            <drop event="Destroy" name="%s%sT%d%s" count="1" />`, p.name, stage, tier, traitSuffix)
		case "3":
			c <- fmt.Sprintf(`            <property name="Extends" value="%s%s" />`, p.name, vanillaStages[n])
			c <- fmt.Sprintf(`            <drop event="Harvest" name="%s" count="4" tag="cropHarvest" />`, p.crop)
			c <- fmt.Sprintf(`            <drop event="Harvest" name="%s" prob="0.5" count="2" tag="bonusCropHarvest" />`, p.crop)
			c <- fmt.Sprintf(`            <drop event="Destroy" name="%s1T%d%s" count="1" prob="0.5" />`, p.name, tier, traitSuffix)
			if strings.ContainsRune(traitSuffix, 'R') {
				c <- fmt.Sprintf(`            <property name="DowngradeBlock" value="%s1T%d%s" />`, p.name, tier, traitSuffix)
			}
		}
		c <- "        </block>"
	}
	return
}

func produceModifications(c chan string) {
	// {code: 'U', name: "Underground", incompatible: []rune{'U', 'S'}},
	c <- `    <append xpath="/blocks/block[contains(@trait, 'U') and @stage='1']">
        <property name="PlantGrowing.LightLevelGrow" value="0" />
        <property name="PlantGrowing.LightLevelStay" value="0" />
    </append>`

	// {code: 'F', name: "Fast"},
	c <- `    <append xpath="/blocks/block[contains(@trait, 'F') and @stage='1' and not @trait='F,F']">
        <property name="PlantGrowing.GrowthRate" value="31.5" />
    </append>`
	c <- `    <append xpath="/blocks/block[@trait='F,F' and @stage='1']">
        <property name="PlantGrowing.GrowthRate" value="15.75" />
    </append>`

	// {code: 'E', name: "Explosive", incompatible: []rune{'E'}},
	// based off of mineCookingPot
	c <- `    <append xpath="/blocks/block[contains(@trait, 'E') and @stage='3' and not @trait='E,E']">
        <property name="Class" value="Mine" /> <!-- a mine destroyed by an *explosion* only has a 33 percent chance to detonate -->
        <property name="Tags" value="Mine" />
        <property name="Material" value="MLandMine" />
        <property name="Collide" value="movement,melee,arrow" />
        <property name="MaxDamage" value="4" />
        <property name="TriggerDelay" value="0.5" />
        <property name="TriggerSound" value="landmine_trigger" />
        <property name="Explosion.ParticleIndex" value="11" />
        <property name="Explosion.RadiusEntities" value="3" />
        <property name="Explosion.EntityDamage" value="300" /> <!-- damage for entities in the center of the explosion. -->
        <property name="CanPickup" value="false" />
    </append>`
	// based off of mineHubcap
	c <- `    <append xpath="/blocks/block[contains(@trait='E,E') and @stage='3']">
        <property name="Class" value="Mine" /> <!-- a mine destroyed by an *explosion* only has a 33 percent chance to detonate -->
        <property name="Tags" value="Mine" />
        <property name="Material" value="MLandMine" />
        <property name="Collide" value="movement,melee,arrow" />
        <property name="MaxDamage" value="4" />
        <property name="TriggerDelay" value="0.5" />
        <property name="TriggerSound" value="landmine_trigger" />
        <property name="Explosion.ParticleIndex" value="11" />
        <property name="Explosion.RadiusEntities" value="5" />
        <property name="Explosion.EntityDamage" value="450" /> <!-- damage for entities in the center of the explosion. -->
        <property name="CanPickup" value="false" />
    </append>`

	// {code: 'B', name: "Bonus", incompatible: []rune{'B'}},
	// {code: 'T', name: "Thorny"},
	// {code: 'S', name: "Sweet", incompatible: []rune{'U'}},

}

func produceBlocks(c chan string) {
	defer close(c)
	c <- "<config>\n    <append xpath=\"/blocks\">"
	for _, plant := range plants {
		for _, tier := range []int{2, 3} {
			produceBlock(c, plant, tier)
			for _, variant1 := range variants {
				produceBlock(c, plant, tier, variant1.code)
				for _, variant2 := range variants {
					if variant1.isCompatibleWith(variant2) {
						produceBlock(c, plant, tier, variant1.code, variant2.code)
					}
				}
			}
		}
	}
	c <- "    </append>"
	produceModifications(c)
	c <- "</config>"
}

func getFile(filename string) (*os.File, error) {
	os.Remove(filename)
	return os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
}

func writeBlocks() error {
	file, err := getFile("blocks-example.xml")
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go produceBlocks(c)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// printVanillaStages()
	// printVariantStages()
	// writeBlocksFile()

	writeBlocks()

}
