package data

import "fmt"

// GraceCorn (Super Corn) is a type of plant
type GraceCorn struct {
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
func (p *GraceCorn) GetCraftTime() int {
	return p.CraftTime
}

// GetDescription returns the seed description for this plant
func (p *GraceCorn) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

// GetDisplayName returns the display name
func (p *GraceCorn) GetDisplayName() string {
	return p.DisplayName
}

// GetName returns the name of this plant
func (p *GraceCorn) GetName() string {
	return p.Name
}

// GetSchematicName returns the schematic name for this plant, given the provided traits
func (p *GraceCorn) GetSchematicName(traits string) string {
	return fmt.Sprintf("plantedGraceCorn1_%sschematic", traits)
}

// IsCompatibleWith checks for trait compatibility with this plant
func (p *GraceCorn) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

// WriteBlockStages produces each of the 3 block stages for this plant
func (p *GraceCorn) WriteBlockStages(c chan string, target, traits string) {
	p.writeStage1(c, target, traits)
	p.writeStage2(c, traits)
	p.writeStage3(c, traits)
}

func (p *GraceCorn) writeStage1(c chan string, target, traits string) {
	c <- fmt.Sprintf(`<block name="plantedGraceCorn1_%s" stage="1" traits="%s">
    <drop event="Destroy" name="plantedGraceCorn1_%s" count="1"/>
    <property name="CreativeMode" value="Player"/>
    <property name="CustomIcon" value="plantedCorn1"/>
    <property name="ItemTypeIcon" value="%s"/>
    <property name="CustomIconTint" value="ff9f9f"/>
    <property name="DescriptionKey" value="plantedGraceCorn1_%sDesc"/>
    <property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
    <property name="Group" value="%s"/>
    <property name="Material" value="Mcorn"/> <!-- mostly for the particle effect -->
    <property name="Mesh" value="cutoutmoveable"/>
    <property name="Model" value="corn_sprout_shape"/>
    <property name="MultiBlockDim" value="1,3,1"/>
    <property name="Place" value="Door"/>
    <property name="PlaceAsRandomRotation" value="true"/>
    <property name="PlantGrowing.Next" value="plantedGraceCorn2_%s"/>
    <property name="Shape" value="New"/>
    <property name="Texture" value="529"/>%s
</block>`, traits, traits, traits, getItemTypeIcon(traits), traits, getCraftingGroup(traits), traits, getUnlock(p, target, traits))
}

func (*GraceCorn) writeStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedGraceCorn2_%s" stage="2" traits="%s">
    <property name="CreativeMode" value="Dev"/>
    <property name="CustomIconTint" value="00ff9f"/>
    <property name="Extends" value="plantedGraceCorn1_%s"/>
    <property name="DescriptionKey" value="plantedGraceCorn2"/>
    <property name="DisplayInfo" value="Description"/>
    <property name="Model" value="corn_growth_shape"/>
    <property name="PlantGrowing.Next" value="plantedGraceCorn3_%s"/>
    <property name="Texture" value="529"/>
</block>`, traits, traits, traits, traits)
}

func (p *GraceCorn) writeStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedGraceCorn3_%s" stage="3" traits="%s" tags="T%dPlant">
    <drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
    <drop event="Harvest" name="foodCropGraceCorn" count="%d" tag="cropHarvest"/>
    <drop event="Harvest" name="foodCropGraceCorn" prob="0.5" count="%d" tag="bonusCropHarvest"/>
    <property name="Collide" value="melee"/>
    <property name="CreativeMode" value="Dev"/>
    <property name="CustomIcon" value="plantedCorn1"/>
    <property name="ItemTypeIcon" value="%s"/>
    <property name="CustomIconTint" value="ff8f9f"/>
    <property name="DescriptionKey" value="plantedGraceCorn3HarvestDesc"/>
    <property name="DisplayInfo" value="Description"/>
    <property name="DisplayType" value="blockMulti"/>
    <property name="FilterTags" value="MC_outdoor,SC_crops"/>
    <property name="HarvestOverdamage" value="false"/>
    <property name="ImposterDontBlock" value="true"/>
    <property name="IsDecoration" value="true"/>
    <property name="IsTerrainDecoration" value="true"/>
    <property name="LightOpacity" value="0"/>
    <property name="Material" value="Mcorn"/>
    <property name="Mesh" value="cutoutmoveable"/>
    <property name="Model" value="corn_harvest_shape"/>
    <property name="MultiBlockDim" value="1,3,1"/>
    <property name="PlantGrowing.FertileLevel" value="1"/>
    <property name="Shape" value="New"/>
    <property name="SortOrder1" value="a090"/>
    <property name="SortOrder2" value="0002"/>
    <property name="Texture" value="529"/>%s
</block>`,
		traits,
		traits,
		calculatePlantTier(traits),
		calculateCropYield(p.CropYield, traits),
		calculateBonusYield(p.BonusYield, traits),
		getItemTypeIcon(traits),
		getRenewableAndDropTags(traits, p))
}
