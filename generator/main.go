package main

import (
	"fmt"
	"os"
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
	singleEffect string
	doubleEffect string
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
		{code: 'U', name: "Underground", incompatible: []rune{'U', 'S'},
			singleEffect: `<property name="PlantGrowing.LightLevelGrow" value="0" /><property name="PlantGrowing.LightLevelStay" value="0" />`},
		{code: 'F', name: "Fast"},
		{code: 'E', name: "Explosive", incompatible: []rune{'B'}},
		{code: 'R', name: "Renewable", incompatible: []rune{'R'}},
		{code: 'T', name: "Thorny"},
		{code: 'S', name: "Sweet", incompatible: []rune{'U'}},
	}
}

func printVanillaStages() {
	fmt.Println("Vanilla Stages")
	count := 0
	for _, plant := range plants {
		for _, stage := range vanillaStages {
			count++
			fmt.Printf("|%02d| %s%s\n", count, plant.name, stage)
		}
	}
}

func printVariantStages() {
	fmt.Println("Variant Stages")
	count := 0
	for _, p := range plants {
		for sNum, s := range variantStages {
			count++
			fmt.Printf("|%03d|. %s%s\n", count, p.name, vanillaStages[sNum])
			for _, v1 := range variants {
				count++
				fmt.Printf("  |%03d|. %s%s%s%c\n", count, p.name, s, variantTiers[1], v1.code)
				for _, v2 := range variants {
					if v1.isCompatibleWith(v2) {
						count++
						fmt.Printf("    |%03d|. %s%s%s%c%c\n", count, p.name, s, variantTiers[2], v1.code, v2.code)
					}
				}
			}
		}
	}
}

// Write all 3 stages to file
func produceBlock(c chan string, p plant, tier string, gmos ...rune) (err error) {
	for n, stage := range variantStages {
		var gmoSuffix, gmoTag string
		switch len(gmos) {
		case 0:
			gmoSuffix = ""
			gmoTag = ""
		case 1:
			gmoSuffix = fmt.Sprintf("%c", gmos[0])
			gmoTag = fmt.Sprintf("%c", gmos[0])
		case 2:
			gmoSuffix = fmt.Sprintf("%c%c", gmos[0], gmos[1])
			gmoTag = fmt.Sprintf("%c,%c", gmos[0], gmos[1])
		default:
			return fmt.Errorf("received too many GMOs")
		}
		c <- fmt.Sprintf(`        <block name="%s%s%s%s" stage="%s" gmo="%s">`, p.name, stage, tier, gmoSuffix, stage, gmoTag)
		c <- fmt.Sprintf(`            <property name="Extends" value="%s%s" />`, p.name, vanillaStages[n])
		if stage == "1" {
			c <- fmt.Sprintf(`            <property name="CustomIcon" value="%s%s" />`, p.name, vanillaStages[n])
		}
		if stage != "3" {
			c <- fmt.Sprintf(`            <property name="PlantGrowing.Next" value="%s%s%s%s" />`, p.name, variantStages[n+1], tier, gmoSuffix)
			c <- fmt.Sprintf(`            <drop event="Destroy" name="%s%s%s%s" count="1" />`, p.name, stage, tier, gmoSuffix)
		}
		if stage == "3" {
			c <- fmt.Sprintf(`            <drop event="Harvest" name="%s" count="4" tag="cropHarvest" />`, p.crop)
			c <- fmt.Sprintf(`            <drop event="Harvest" name="%s" prob="0.5" count="2" tag="bonusCropHarvest" />`, p.crop)
			c <- fmt.Sprintf(`            <drop event="Destroy" name="%s1%s%s" count="1" prob="0.5" />`, p.name, tier, gmoSuffix)
		}
		c <- "        </block>"
	}
	return
}

func produceBlocks(c chan string) {
	defer close(c)
	c <- "<config>\n    <append xpath=\"/blocks\">"
	for _, plant := range plants {
		for _, tier := range variantTiers {
			produceBlock(c, plant, tier)
			for _, variant1 := range variants {
				produceBlock(c, plant, tier, variant1.code)
				for _, variant2 := range variants {
					produceBlock(c, plant, tier, variant1.code, variant2.code)
				}
			}
		}
	}

	c <- "    </append>\n</config>"
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
