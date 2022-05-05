package data

import (
	"fmt"
)

type Mushroom struct {
	Name              string
	DisplayName       string
	PreferredConsumer string
	Description       string
	CropYield         int
	BonusYield        int
	CraftTime         int
}

func CreateMushroom() *Mushroom {
	return &Mushroom{
		Name:              "Mushroom",
		DisplayName:       "Mushroom Spores",
		Description:       `Mushroom spores can be planted on all surfaces and walls and will grow without sunlight.`,
		PreferredConsumer: "Boars",
		CropYield:         2,
		BonusYield:        1,
		CraftTime:         2,
	}
}

func (p *Mushroom) GetCraftTime() int {
	return p.CraftTime
}

func (p *Mushroom) GetDescription() string {
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

func (p *Mushroom) WriteBlockStages(c chan string, traits string) {
	p.WriteStage1(c, traits)
	p.WriteStage2(c, traits)
	p.WriteStage3(c, traits)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedMushroom1Schematic"/>
func (*Mushroom) WriteStage1(c chan string, traits string) {
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
	<property name="UnlockedBy" value="perkLivingOffTheLand"/>
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits)
}

func (*Mushroom) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom2_%s" stage="2" traits="%s">
	<property name="Collide" value="melee"/>
	<property name="Extends" value="plantedMushroom1_%s"/>
	<property name="Model" value="OutdoorDecor/mushroom_growth" param1="main_mesh"/>
	<property name="PlantGrowing.Next" value="plantedMushroom3_%s"/>
</block>`, traits, traits, traits, traits)
	// TODO: <property name="CreativeMode" value="None"/>
}

func (p *Mushroom) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom3_%s" stage="3" traits="%s">
	<drop event="Destroy" name="plantedMushroom1_%s" count="1" prob="0.5"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="foodCropMushrooms" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="foodCropMushrooms" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CropsGrown.BonusHarvestDivisor" value="16"/>
	<property name="CustomIcon" value="plantedMushroom3Harvest"/>
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
		traits,
		calculateCropYield(p.CropYield, traits),
		calculateBonusYield(p.BonusYield, traits),
		traits,
		optionallyAddRenewable(traits, p))
	// TODO: <property name="CreativeMode" value="None"/>
}
