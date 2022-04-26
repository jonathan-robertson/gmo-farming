package gen

import (
	"fmt"
	"strings"
)

type Mushroom struct {
	CropYield  int
	BonusYield int
	CraftTime  int
}

func CreateMushroom() *Mushroom {
	return &Mushroom{
		CropYield:  2,
		BonusYield: 1,
		CraftTime:  2,
	}
}

func (*Mushroom) IsCompatibleWith(traits string) bool {
	return !strings.ContainsRune(traits, 'U')
}

func (mushroom *Mushroom) WriteRecipe(c chan string, tier int, traits string) {
	suffix := fmt.Sprintf("T%d%s", tier, traits)

	c <- fmt.Sprintf(`<recipe name="plantedMushroom1%s" count="1" craft_time="%d" tags="learnable">`,
		suffix,
		calculateCraftTime(mushroom.CraftTime, tier, traits))
	c <- `<ingredient name="plantedMushroom1" count="1"/>`
	// TODO:
	// <ingredient name="foodCropMushrooms" count="5"/>
	// <ingredient name="resourceClayLump" count="2"/>

	c <- `</recipe>`
}

func (mushroom *Mushroom) WriteStages(c chan string, tier int, traits string) {
	suffix := fmt.Sprintf("T%d%s", tier, traits)
	mushroom.WriteStage1(c, tier, traits, suffix)
	mushroom.WriteStage2(c, tier, traits, suffix)
	mushroom.WriteStage3(c, tier, traits, suffix)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedMushroom1Schematic"/>
func (*Mushroom) WriteStage1(c chan string, tier int, traits, suffix string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom1%s" stage="1" traits="%s">
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
	<property name="PlantGrowing.Next" value="plantedMushroom2%s"/>
	<property name="HarvestOverdamage" value="false"/>
	<drop event="Destroy" name="plantedMushroom1%s" count="1"/>
	<property name="EconomicValue" value="12"/>
	<property name="EconomicBundleSize" value="5"/>
	<property name="Group" value="Food/Cooking"/>
	<property name="PickupJournalEntry" value="farmingTip"/>

	<property name="CustomIcon" value="plantedMushroom1"/>
</block>`,
		suffix,
		traits,
		suffix,
		suffix)
}

func (*Mushroom) WriteStage2(c chan string, tier int, traits, suffix string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom2%s" stage="2" traits="%s">
	<property name="Extends" value="plantedMushroom1%s"/>
	
	<property name="Model" value="OutdoorDecor/mushroom_growth" param1="main_mesh"/>
	<property name="PlantGrowing.Next" value="plantedMushroom3%s"/>
	<property name="Collide" value="melee"/>
</block>`, suffix, traits, suffix, suffix)
	// TODO: <property name="CreativeMode" value="None"/>
}

func (mushroom *Mushroom) WriteStage3(c chan string, tier int, traits, suffix string) {
	c <- fmt.Sprintf(`<block name="plantedMushroom3%s" stage="3" traits="%s">
	<property name="Material" value="Mmushrooms"/>
	<property name="DisplayType" value="blockMulti"/>
	<property name="DisplayInfo" value="Description"/>
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

	<property name="DescriptionKey" value="plantedMushroom3%s"/>
	<property name="CustomIcon" value="plantedMushroom3Harvest"/>
	
	<drop event="Harvest" name="foodCropMushrooms" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="foodCropMushrooms" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<drop event="Destroy" name="plantedMushroom1%s" count="1" prob="0.5"/>
	%s
</block>`,
		suffix,
		traits,
		suffix,
		calculateCropYield(mushroom.CropYield, tier, traits),
		calculateBonusYield(mushroom.BonusYield, tier, traits),
		suffix,
		optionallyAddRenewable(traits, "plantedMushroom1", suffix))
	// TODO: <property name="CreativeMode" value="None"/>
}
