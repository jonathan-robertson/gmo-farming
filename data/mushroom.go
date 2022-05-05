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

func (m *Mushroom) GetCraftTime() int {
	return m.CraftTime
}

func (m *Mushroom) GetDescription() string {
	return m.Description
}

func (m *Mushroom) GetDisplayName() string {
	return m.DisplayName
}

func (m *Mushroom) GetName() string {
	return m.Name
}

func (m *Mushroom) GetPreferredConsumer() string {
	return m.PreferredConsumer
}

func (mushroom *Mushroom) WriteBlockStages(c chan string, traits string) {
	mushroom.WriteStage1(c, traits)
	mushroom.WriteStage2(c, traits)
	mushroom.WriteStage3(c, traits)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedMushroom1Schematic"/>
func (*Mushroom) WriteStage1(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom1_%s" stage="1" traits="%s">
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon,DescriptionKey,MultiBlockDim,OnlySimpleRotations"/>
	<property name="CreativeMode" value="Player"/>
	<property name="DisplayInfo" value="Name"/>
	<property name="UnlockedBy" value="perkLivingOffTheLand"/>
	<property name="Material" value="Mmushrooms"/>
	<property name="LightOpacity" value="0"/>
	<property name="Shape" value="Ext3dModel"/>
	<property name="Texture" value="293"/>
	<property name="Mesh" value="models"/>
	<property name="Model" value="OutdoorDecor/mushroom_sprout" param1="main_mesh"/>
	<property name="Collide" value="melee"/>
	<property name="HandleFace" value="Bottom"/>
	<property name="PlantGrowing.LightLevelGrow" value="0"/>
	<property name="PlantGrowing.LightLevelStay" value="0"/>
	<property name="PlantGrowing.FertileLevel" value="0"/>
	<property name="PlantGrowing.Next" value="plantedMushroom2_%s"/>
	<property name="HarvestOverdamage" value="false"/>
	<drop event="Destroy" name="plantedMushroom1_%s" count="1"/>
	<property name="EconomicValue" value="12"/>
	<property name="EconomicBundleSize" value="5"/>
	<property name="PickupJournalEntry" value="farmingTip"/>
	
	<property name="CustomIcon" value="plantedMushroom1"/>
	<property name="DescriptionKey" value="plantedMushroom1_%sDesc"/>
	<property name="Group" value="%s"/>
</block>`,
		traits,
		traits,
		traits,
		traits,
		traits,
		getCraftingGroup(traits))
}

func (*Mushroom) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom2_%s" stage="2" traits="%s">
	<property name="Extends" value="plantedMushroom1_%s"/>
	
	<property name="Model" value="OutdoorDecor/mushroom_growth" param1="main_mesh"/>
	<property name="PlantGrowing.Next" value="plantedMushroom3_%s"/>
	<property name="Collide" value="melee"/>
</block>`, traits, traits, traits, traits)
	// TODO: <property name="CreativeMode" value="None"/>
}

func (mushroom *Mushroom) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom3_%s" stage="3" traits="%s">
	<property name="Material" value="Mmushrooms"/>
	<property name="DisplayType" value="blockMulti"/>
	<property name="Shape" value="Ext3dModel"/>
	<property name="Texture" value="293"/>
	<property name="Mesh" value="models"/>
	<property name="Model" value="OutdoorDecor/mushroom_harvest" param1="main_mesh"/>
	<property name="Collide" value="melee"/>
	<property name="VehicleHitScale" value=".1"/>
	<property name="IsTerrainDecoration" value="true"/>
	<property name="PlantGrowing.FertileLevel" value="0"/>
	<property name="CropsGrown.BonusHarvestDivisor" value="16"/>
	<property name="HarvestOverdamage" value="false"/>
	
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<property name="PickupJournalEntry" value="farmingTip"/>
	<property name="FilterTags" value="MC_outdoor,SC_crops"/>

	<property name="DescriptionKey" value="plantedMushroom3_%s"/>
	<property name="CustomIcon" value="plantedMushroom3Harvest"/>
	
	<drop event="Harvest" name="foodCropMushrooms" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="foodCropMushrooms" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<drop event="Destroy" name="plantedMushroom1_%s" count="1" prob="0.5"/>
	%s
</block>`,
		traits,
		traits,
		traits,
		calculateCropYield(mushroom.CropYield, traits),
		calculateBonusYield(mushroom.BonusYield, traits),
		traits,
		optionallyAddRenewable(traits, mushroom))
	// TODO: <property name="CreativeMode" value="None"/>
}
