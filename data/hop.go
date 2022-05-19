package data

import "fmt"

// Hop is a type of plant
type Hop struct {
	Name               string
	NamePlural         string
	DisplayName        string
	Description        string
	CropYield          int
	BonusYield         int
	CraftTime          int
	incompatibleTraits []rune
}

// GetCraftTime returns the time required to craft this seed
func (p *Hop) GetCraftTime() int {
	return p.CraftTime
}

// GetDescription returns the seed description for this plant
func (p *Hop) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

// GetDisplayName returns the display name
func (p *Hop) GetDisplayName() string {
	return p.DisplayName
}

// GetName returns the name of this plant
func (p *Hop) GetName() string {
	return p.Name
}

// GetSchematicName returns the schematic name for this plant, given the provided traits
func (p *Hop) GetSchematicName(traits string) string {
	return fmt.Sprintf("plantedHop1_%sschematic", traits)
}

// IsCompatibleWith checks for trait compatibility with this plant
func (p *Hop) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

// WriteBlockStages produces each of the 3 block stages for this plant
func (p *Hop) WriteBlockStages(c chan string, target, traits string) {
	p.writeStage1(c, target, traits)
	p.writeStage2(c, traits)
	p.writeStage3(c, traits)
}

func (p *Hop) writeStage1(c chan string, target, traits string) {
	c <- fmt.Sprintf(`<block name="plantedHop1_%s" stage="1" traits="%s">
	<drop event="Destroy" name="plantedHop1_%s" count="1"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="plantedHop1"/>
	<property name="DescriptionKey" value="plantedHop1_%sDesc"/>
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
	<property name="Group" value="%s"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="PlantGrowing.Next" value="plantedHop2_%s"/>
	<property name="Texture" value="447"/>
	%s
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits, optionallyAddUnlock(p, target, traits))
}

func (*Hop) writeStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedHop2_%s" stage="2" traits="%s">
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIconTint" value="00ff80"/>
	<property name="Extends" value="plantedHop1_%s"/>
	<property name="PlantGrowing.Next" value="plantedHop3_%s"/>
	<property name="Texture" value="448"/>
</block>`, traits, traits, traits, traits)
}

func (p *Hop) writeStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedHop3_%s" stage="3" traits="%s" tags="T%dPlant">
	<drop event="Destroy" count="0" />
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="resourceCropHopsFlower" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="resourceCropHopsFlower" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIcon" value="plantedHop1"/>
	<property name="CustomIconTint" value="ff8000"/>
	<property name="DescriptionKey" value="plantedHop3_%s"/>
	<property name="DisplayInfo" value="Description"/> <!-- also valid: "Name" -->
	<property name="DisplayType" value="blockMulti"/>
	<property name="FilterTags" value="MC_outdoor,SC_crops"/>
	<property name="HarvestOverdamage" value="false"/>
	<property name="ImposterDontBlock" value="true"/>
	<property name="IsDecoration" value="true"/>
	<property name="IsTerrainDecoration" value="true"/>
	<property name="LightOpacity" value="0"/>
	<property name="Material" value="Mplants"/>
	<property name="Mesh" value="grass"/>
	<property name="MultiBlockDim" value="1,2,1"/>
	<property name="PlantGrowing.FertileLevel" value="1"/>
	<property name="Shape" value="BillboardPlant"/>
	<property name="SortOrder1" value="a090"/>
	<property name="SortOrder2" value="0002"/>
	<property name="Texture" value="395"/>
	<property name="Texture" value="449"/>
	%s
</block>`,
		traits,
		traits,
		calculatePlantTier(traits),
		calculateCropYield(p.CropYield, traits),
		calculateBonusYield(p.BonusYield, traits),
		traits,
		optionallyAddRenewable(traits, p))
}
