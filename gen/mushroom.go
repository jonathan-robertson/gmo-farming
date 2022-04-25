package gen

import (
	"fmt"
	"strings"
)

type Mushroom struct{}

func (p *Mushroom) IsCompatibleWith(traits string) bool {
	return !strings.ContainsRune(traits, 'U')
}

func (p *Mushroom) WriteStages(c chan string, tier int, traits string) {
	p.WriteStage1(c, tier, traits)
	p.WriteStage2(c, tier, traits)
	p.WriteStage3(c, tier, traits)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedMushroom1Schematic"/>
func (m *Mushroom) WriteStage1(c chan string, tier int, traits string) {
	suffix := fmt.Sprintf("T%d%s", tier, traits)
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
</block>`, suffix, traits, suffix, suffix)
}

func (m *Mushroom) WriteStage2(c chan string, tier int, traits string) {
	suffix := fmt.Sprintf("T%d%s", tier, traits)

	c <- fmt.Sprintf(`<block name="plantedMushroom2%s" stage="2" traits="%s">
	<property name="Extends" value="plantedMushroom1%s"/>
	<property name="CustomIcon" value="plantedMushroom1"/>
	
	<property name="Model" value="OutdoorDecor/mushroom_growth" param1="main_mesh"/>
	<property name="PlantGrowing.Next" value="plantedMushroom3%s"/>
	<property name="Collide" value="melee"/>
</block>`, suffix, traits, suffix, suffix)
	// TODO: <property name="CreativeMode" value="None"/>
}

func (m *Mushroom) WriteStage3(c chan string, tier int, traits string) {
	suffix := fmt.Sprintf("T%d%s", tier, traits)

	// Apply tier bonus to yield
	yield := 2
	bonusYield := 1
	switch tier {
	case 2:
		yield *= 2
		bonusYield *= 2
	case 3:
		yield *= 4
		bonusYield *= 4
	}

	// Apply Bonus trait to yield
	if strings.Contains(traits, "BB") {
		yield *= 2
		bonusYield *= 2
	} else if strings.Contains(traits, "B") {
		yield = int(float64(yield) * 1.5)
		bonusYield = int(float64(bonusYield) * 1.5)
	}

	// Apply Renewable Trait
	var renewableLine string
	if strings.Contains(traits, "R") {
		renewableLine = fmt.Sprintf(`<property name="DowngradeBlock" value="plantedMushroom1%s" />`, suffix)
	}

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
</block>`, suffix, traits, suffix, yield, bonusYield, suffix, renewableLine)
	// TODO: <property name="CreativeMode" value="None"/>
}
