package main

import (
	"fmt"
	"os"
	"strings"
)

type plant struct {
	name string
}
type stage struct {
	text string
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

var vanillaPlants []plant
var vanillaStages []stage

var variantStages []stage
var variantTiers []tier
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
	vanillaPlants = []plant{
		{name: "plantedMushroom"},
		{name: "plantedYucca"},
		{name: "plantedCotton"},
		{name: "plantedCoffee"},
		{name: "plantedGoldenrod"},
		{name: "plantedAloe"},
		{name: "plantedBlueberry"},
		{name: "plantedPotato"},
		{name: "plantedChrysanthemum"},
		{name: "plantedCorn"},
		{name: "plantedGraceCorn"},
		{name: "plantedHop"},
		{name: "plantedPumpkin"},
	}
	vanillaStages = []stage{
		{text: "1"},
		{text: "2"},
		{text: "3HarvestPlayer"},
	}

	variantStages = []stage{
		{text: "1"},
		{text: "2"},
		{text: "3"},
	}
	variantTiers = []tier{
		{text: "T1"},
		{text: "T2"},
		{text: "T3"},
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
	for _, plant := range vanillaPlants {
		for _, stage := range vanillaStages {
			count++
			fmt.Printf("|%02d| %s%s\n", count, plant.name, stage.text)
		}
	}
}

func printVariantStages() {
	fmt.Println("Variant Stages")
	count := 0
	for _, p := range vanillaPlants {
		for sNum, s := range variantStages {
			count++
			fmt.Printf("|%03d|. %s%s\n", count, p.name, vanillaStages[sNum].text)
			for _, v1 := range variants {
				count++
				fmt.Printf("  |%03d|. %s%s%s%c\n", count, p.name, s.text, variantTiers[1].text, v1.code)
				for _, v2 := range variants {
					if v1.isCompatibleWith(v2) {
						count++
						fmt.Printf("    |%03d|. %s%s%s%c%c\n", count, p.name, s.text, variantTiers[2].text, v1.code, v2.code)
					}
				}
			}
		}
	}
}

func writeVariantBlocks(file *os.File) {
	fmt.Println("Variant Stages")
	count := 0
	for _, p := range vanillaPlants {
		for stageNum, s := range variantStages {
			count++
			fmt.Printf("|%03d|. %s%s\n", count, p.name, vanillaStages[stageNum].text)
			for _, v1 := range variants {
				count++
				fmt.Printf("  |%03d|. %s%s%s%c\n", count, p.name, s.text, variantTiers[1].text, v1.code)

				var lines []string
				lines = append(lines, fmt.Sprintf(`<block name="%s%s%s%c" stage="%s" gmo="%c">`, p.name, s.text, variantTiers[1].text, v1.code, s.text, v1.code))
				vanillaName := p.name + vanillaStages[stageNum].text
				lines = append(lines, fmt.Sprintf(`<property name="Extends" value="%s" />`, vanillaName))
				if "1" == s.text {
					lines = append(lines, fmt.Sprintf(`<property name="CustomIcon" value="%s" />`, vanillaName))
				}
				if "3" != s.text {
					lines = append(lines, fmt.Sprintf(`<property name="PlantGrowing.Next" value="%s%s%s%c" />`, p.name, variantStages[stageNum+1].text, variantTiers[1].text, v1.code))
					lines = append(lines, fmt.Sprintf(`<drop event="Destroy" name="%s%s%s%c" count="1" />`, p.name, s.text, variantTiers[1].text, v1.code))
				}
				if "3" == s.text {
					lines = append(lines, fmt.Sprintf(`<drop event="Destroy" name="%s1%s%c" count="1" prob="0.5" />`, p.name, variantTiers[1].text, v1.code))
				}
				lines = append(lines, `</block>`)

				if _, err := file.WriteString(strings.Join(lines, "\n")); err != nil {
					panic(err)
				}
				/*
					for _, v2 := range variants {
						if v1.isCompatibleWith(v2) {
							count++
							fmt.Printf("    |%03d|. %s%s%s%c%c\n", count, p.name, s.text, variantTiers[2].text, v1.code, v2.code)
						}
					}
				*/
			}
		}
	}
}

func getFile(filename string) (*os.File, error) {
	os.Remove(filename)
	return os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
}

func writeBlocksFile() {
	file, err := getFile("blocks-example.xml")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	const header string = "<config>\n<append xpath=\"/blocks\">\n"
	const footer string = "</append>\n</config>\n"

	if _, err = file.WriteString(header); err != nil {
		panic(err)
	}
	writeVariantBlocks(file)
	if _, err = file.WriteString(footer); err != nil {
		panic(err)
	}
}

func main() {
	//printVanillaStages()
	//printVariantStages()
	writeBlocksFile()
}
