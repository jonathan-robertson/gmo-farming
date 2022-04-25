package gen

import (
	"fmt"
	"strings"
)

type Corn struct{}

func (p *Corn) IsCompatibleWith(traits string) bool {
	return true
}

func (p *Corn) WriteStages(c chan string, tier int, traits string) {
	p.WriteStage1(c, tier, traits)
	p.WriteStage2(c, tier, traits)
	p.WriteStage3(c, tier, traits)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedCorn1Schematic"/>
func (p *Corn) WriteStage1(c chan string, tier int, traits string) {
	suffix := fmt.Sprintf("T%d%s", tier, traits)
	c <- fmt.Sprintf(`<block name="plantedCorn1%s" stage="1" traits="%s">
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
	<property name="CreativeMode" value="Player"/>
	<property name="UnlockedBy" value="perkLivingOffTheLand"/>
	<property name="Material" value="Mcorn"/> <!-- mostly for the particle effect -->
	<property name="Shape" value="New"/>
	<property name="Model" value="corn_sprout_shape"/>
	<property name="Place" value="Door"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="Mesh" value="cutoutmoveable"/>
	<property name="MultiBlockDim" value="1,3,1"/>
	<property name="Texture" value="529"/>
	<property name="PlantGrowing.Next" value="plantedCorn2%s"/>
	<property name="Group" value="Food/Cooking"/>
	<drop event="Destroy" name="plantedCorn1%s" count="1"/>
	
	<property name="CustomIcon" value="plantedCorn1"/>
</block>`, suffix, traits, suffix, suffix)
}

func (p *Corn) WriteStage2(c chan string, tier int, traits string) {
	suffix := fmt.Sprintf("T%d%s", tier, traits)
	c <- fmt.Sprintf(`<block name="plantedCorn2%s" stage="2" traits="%s">
	<property name="Extends" value="plantedCorn1%s"/>
	
	<property name="Shape" value="New"/>
	<property name="Model" value="corn_growth_shape"/>
	<property name="Mesh" value="cutoutmoveable"/>
	<property name="MultiBlockDim" value="1,3,1"/>
	<property name="Texture" value="529"/>
	<property name="PlantGrowing.Next" value="plantedCorn3HarvestPlayer"/>
	<drop event="Destroy" name="plantedCorn1%s" count="1"/>
</block>`, suffix, traits, suffix, suffix)
	// TODO: <property name="CreativeMode" value="None"/>
}

func (p *Corn) WriteStage3(c chan string, tier int, traits string) {
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
		renewableLine = fmt.Sprintf(`<property name="DowngradeBlock" value="plantedCorn1%s" />`, suffix)
	}

	c <- fmt.Sprintf(`<block name="plantedCorn3%s" stage="3" traits="%s">
	<property name="DisplayInfo" value="Description"/> <!-- also valid: "Name" -->
	<property name="DisplayType" value="blockMulti"/>
	<property name="LightOpacity" value="0"/>
	<property name="ImposterDontBlock" value="true"/>
	<property name="Collide" value="melee"/>
	<property name="IsTerrainDecoration" value="true"/>
	<property name="IsDecoration" value="true"/>
	
	<property name="PlantGrowing.FertileLevel" value="1"/>
	<property name="HarvestOverdamage" value="false"/>
	<drop event="Destroy" count="0"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<property name="FilterTags" value="MC_outdoor,SC_crops"/>
	<property name="SortOrder1" value="a090"/>
	<property name="SortOrder2" value="0002"/>

	<property name="Material" value="Mcorn"/>
	<property name="Shape" value="New"/>
	<property name="Model" value="corn_harvest_shape"/>
	<property name="Mesh" value="cutoutmoveable"/>
	<property name="MultiBlockDim" value="1,3,1"/>
	<property name="Texture" value="529"/>

	<property name="DescriptionKey" value="plantedCorn3%s"/>
	<property name="CustomIcon" value="plantedCorn3Harvest"/>

	<drop event="Harvest" name="foodCropCorn" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="foodCropCorn" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<drop event="Destroy" name="plantedCorn1%s" count="1" prob="0.5"/>
	%s
</block>`, suffix, traits, suffix, yield, bonusYield, suffix, renewableLine)
	// TODO: <property name="CreativeMode" value="None"/>
}
