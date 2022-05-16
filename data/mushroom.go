package data

import (
	"fmt"
)

type Mushroom struct {
	Name               string
	DisplayName        string
	PreferredConsumer  string
	Description        string
	CropYield          int
	BonusYield         int
	CraftTime          int
	incompatibleTraits []rune
}

func CreateMushroom() *Mushroom {
	return &Mushroom{
		Name:               "Mushroom",
		DisplayName:        "Mushroom Spores",
		Description:        `Mushroom spores can be planted on all surfaces and walls and will grow without sunlight.`,
		PreferredConsumer:  "Boars",
		CropYield:          2,
		BonusYield:         1,
		CraftTime:          2,
		incompatibleTraits: []rune{'U'},
	}
}

func (p *Mushroom) GetCraftTime() int {
	return p.CraftTime
}

func (p *Mushroom) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

func (p *Mushroom) GetDisplayName() string {
	return p.DisplayName
}

func (p *Mushroom) GetName() string {
	return p.Name
}

func (p *Mushroom) GetPreferredConsumer() string {
	return p.PreferredConsumer
}

func (p *Mushroom) GetSchematicName(traits string) string {
	return fmt.Sprintf("plantedMushroom1_%sschematic", traits)
}

func (p *Mushroom) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

func (p *Mushroom) WriteBlockStages(c chan string, target, traits string) {
	p.WriteStage1(c, target, traits)
	p.WriteStage2(c, traits)
	p.WriteStage3(c, traits)
}

// TODO: return to mushroom... seems like overkill - why not extend naturally?
func (p *Mushroom) WriteStage1(c chan string, target, traits string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom1_%s" stage="1" traits="%s">
	<drop event="Destroy" name="plantedMushroom1_%s" count="1"/>
	<property name="Collide" value="melee"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="plantedMushroom1"/>
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
	%s
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits, optionallyAddUnlock(p, target, traits))
}

func (*Mushroom) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom2_%s" stage="2" traits="%s">
	<property name="Collide" value="melee"/>
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIconTint" value="00ff80"/>
	<property name="Extends" value="plantedMushroom1_%s"/>
	<property name="Model" value="OutdoorDecor/mushroom_growth" param1="main_mesh"/>
	<property name="PlantGrowing.Next" value="plantedMushroom3_%s"/>
</block>`, traits, traits, traits, traits)
}

func (p *Mushroom) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom3_%s" stage="3" traits="%s" tags="T%dPlant">
	<drop event="Destroy" name="plantedMushroom1_%s" count="1" prob="0.5"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="foodCropMushrooms" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="foodCropMushrooms" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CreativeMode" value="Dev"/>
	<property name="CropsGrown.BonusHarvestDivisor" value="16"/>
	<property name="CustomIcon" value="plantedMushroom1"/>
	<property name="CustomIconTint" value="ff8000"/>
	<property name="DescriptionKey" value="plantedMushroom3_%s"/>
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
		traits,
		calculateCropYield(p.CropYield, traits),
		calculateBonusYield(p.BonusYield, traits),
		traits,
		optionallyAddRenewable(traits, p))
}
