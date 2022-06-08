package data

import (
	"fmt"
)

// Mushroom is a type of plant
type Mushroom struct {
	Name               string
	DisplayName        string
	Description        string
	CropYield          int
	BonusYield         int
	CraftTime          int
	incompatibleTraits []rune
}

// GetCraftTime returns the time required to craft this seed
func (p *Mushroom) GetCraftTime() int {
	return p.CraftTime
}

// GetDescription returns the seed description for this plant
func (p *Mushroom) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

// GetDisplayName returns the display name
func (p *Mushroom) GetDisplayName() string {
	return p.DisplayName
}

// GetName returns the name of this plant
func (p *Mushroom) GetName() string {
	return p.Name
}

// GetSchematicName returns the schematic name for this plant, given the provided traits
func (p *Mushroom) GetSchematicName(traits string) string {
	return fmt.Sprintf("plantedMushroom1_%sschematic", traits)
}

// IsCompatibleWith checks for trait compatibility with this plant
func (p *Mushroom) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

// WriteBlockStages produces each of the 3 block stages for this plant
func (p *Mushroom) WriteBlockStages(c chan string, target, traits string) {
	p.writeStage1(c, target, traits)
	p.writeStage2(c, traits)
	p.writeStage3(c, traits)
}

// TODO: return to mushroom... seems like overkill - why not extend naturally?
func (p *Mushroom) writeStage1(c chan string, target, traits string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom1_%s" stage="1" traits="%s">
    <drop event="Destroy" name="plantedMushroom1_%s" count="1"/>
    <property name="Collide" value="melee"/>
    <property name="CreativeMode" value="Player"/>
    <property name="CustomIcon" value="plantedMushroom1"/>
    <property name="ItemTypeIcon" value="%s"/>
    <property name="DescriptionKey" value="plantedMushroom1_%sDesc"/>
    <property name="DisplayInfo" value="Name"/>
    <property name="EconomicBundleSize" value="5"/>
    <property name="EconomicValue" value="12"/>
    <property name="Extends" value="cropsGrowingMaster" param1="CustomIcon,DescriptionKey,MultiBlockDim,OnlySimpleRotations"/>
    <property name="Group" value="%s"/>
    <property name="HandleFace" value="Bottom"/>
    <property name="HarvestOverdamage" value="false"/>
    <property name="LightOpacity" value="0"/>
    <property name="Material" value="Mmushrooms"/>
    <property name="Mesh" value="models"/>
    <property name="Model" value="OutdoorDecor/mushroom_sprout" param1="main_mesh"/>
    <property name="PickupJournalEntry" value="farmingTip"/>
    <property name="PlantGrowing.FertileLevel" value="0"/>
    <property name="PlantGrowing.LightLevelGrow" value="0"/>
    <property name="PlantGrowing.LightLevelStay" value="0"/>
    <property name="PlantGrowing.Next" value="plantedMushroom2_%s"/>
    <property name="Shape" value="Ext3dModel"/>
    <property name="Texture" value="293"/>
    <property name="UnlockedBy" value="%s"/>
</block>`, traits, traits, traits, getItemTypeIcon(traits), traits, getCraftingGroup(traits), traits, getUnlock(p, target, traits))
}

func (*Mushroom) writeStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom2_%s" stage="2" traits="%s">
    <property name="Collide" value="melee"/>
    <property name="CreativeMode" value="Dev"/>
    <property name="CustomIconTint" value="00ff80"/>
    <property name="Extends" value="plantedMushroom1_%s"/>
    <property name="DescriptionKey" value="plantedMushroom2"/>
    <property name="DisplayInfo" value="Description"/>
    <property name="Model" value="OutdoorDecor/mushroom_growth" param1="main_mesh"/>
    <property name="PlantGrowing.Next" value="plantedMushroom3_%s"/>
</block>`, traits, traits, traits, traits)
}

func (p *Mushroom) writeStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom3_%s" stage="3" traits="%s" tags="T%dPlant">
    <drop event="Destroy" count="0" />
    <drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
    <drop event="Harvest" name="foodCropMushrooms" count="%d" tag="cropHarvest"/>
    <drop event="Harvest" name="foodCropMushrooms" prob="0.5" count="%d" tag="bonusCropHarvest"/>
    <property name="Collide" value="melee"/>
    <property name="CreativeMode" value="Dev"/>
    <property name="CropsGrown.BonusHarvestDivisor" value="16"/>
    <property name="CustomIcon" value="plantedMushroom1"/>
    <property name="ItemTypeIcon" value="%s"/>
    <property name="CustomIconTint" value="ff8000"/>
    <property name="DescriptionKey" value="plantedMushroom3HarvestDesc"/>
    <property name="DisplayInfo" value="Description"/>
    <property name="DisplayType" value="blockMulti"/>
    <property name="FilterTags" value="MC_outdoor,SC_crops"/>
    <property name="HarvestOverdamage" value="false"/>
    <property name="IsTerrainDecoration" value="true"/>
    <property name="Material" value="Mmushrooms"/>
    <property name="Mesh" value="models"/>
    <property name="Model" value="OutdoorDecor/mushroom_harvest" param1="main_mesh"/>
    <property name="PickupJournalEntry" value="farmingTip"/>
    <property name="PlantGrowing.FertileLevel" value="0"/>
    <property name="Shape" value="Ext3dModel"/>
    <property name="Texture" value="293"/>
    <property name="VehicleHitScale" value=".1"/>
    %s
</block>`,
		traits,
		traits,
		calculatePlantTier(traits),
		calculateCropYield(p.CropYield, traits),
		calculateBonusYield(p.BonusYield, traits),
		getItemTypeIcon(traits),
		optionallyAddRenewable(traits, p))
}
