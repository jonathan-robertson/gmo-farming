package gen

import (
	"fmt"
)

type Corn struct {
	Name              string
	NamePlural        string
	DisplayName       string
	Description       string
	PreferredConsumer string
	CropYield         int
	BonusYield        int
	CraftTime         int
}

func CreateCorn() *Corn {
	return &Corn{
		Name:              "Corn",     // TODO: confirm
		NamePlural:        "Corn",     // TODO: confirm
		DisplayName:       "Corn",     // TODO: confirm
		PreferredConsumer: "Chickens", // TODO: confirm
		CropYield:         2,
		BonusYield:        1,
		CraftTime:         2,
	}
}

func (c *Corn) GetName() string {
	return c.Name
}

func (c *Corn) GetNamePlural() string {
	return c.NamePlural
}

func (c *Corn) GetDisplayName() string {
	return c.DisplayName
}

func (c *Corn) GetDescription() string {
	if c.Description == "" {
		return `Plant these seeds on a craftable Farm Plot block to grow plants for you to harvest.\n\nWhen harvested, there is a 50% chance to get a seed back for replanting.`
	}
	return c.Description
}

func (c *Corn) GetPreferredConsumer() string {
	return c.PreferredConsumer
}

func (c *Corn) GetCraftTime() int {
	return c.CraftTime
}

func (*Corn) IsCompatibleWith(trait rune) bool {
	return true
}

func (corn *Corn) WriteBlockStages(c chan string, traits string) {
	corn.WriteStage1(c, traits)
	corn.WriteStage2(c, traits)
	corn.WriteStage3(c, traits)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedCorn1Schematic"/>
// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand"/>
func (*Corn) WriteStage1(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedCorn1_%s" stage="1" traits="%s">
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
	<property name="CreativeMode" value="Player"/>
	
	<property name="Material" value="Mcorn"/> <!-- mostly for the particle effect -->
	<property name="Shape" value="New"/>
	<property name="Model" value="corn_sprout_shape"/>
	<property name="Place" value="Door"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="Mesh" value="cutoutmoveable"/>
	<property name="MultiBlockDim" value="1,3,1"/>
	<property name="Texture" value="529"/>
	<property name="PlantGrowing.Next" value="plantedCorn2_%s"/>
	<property name="Group" value="Food/Cooking"/>
	<drop event="Destroy" name="plantedCorn1_%s" count="1"/>
	
	<property name="CustomIcon" value="plantedCorn1"/>
	<property name="DescriptionKey" value="plantedCorn1_%sDesc"/>
</block>`,
		traits,
		traits,
		traits,
		traits,
		traits)
}

func (*Corn) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedCorn2_%s" stage="2" traits="%s">
	<property name="Extends" value="plantedCorn1_%s"/>
	
	<property name="Shape" value="New"/>
	<property name="Model" value="corn_growth_shape"/>
	<property name="Mesh" value="cutoutmoveable"/>
	<property name="MultiBlockDim" value="1,3,1"/>
	<property name="Texture" value="529"/>
	<property name="PlantGrowing.Next" value="plantedCorn3HarvestPlayer"/>
	<drop event="Destroy" name="plantedCorn1_%s" count="1"/>
</block>`,
		traits,
		traits,
		traits,
		traits)
	// TODO: <property name="CreativeMode" value="None"/>
}

func (corn *Corn) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedCorn3_%s" stage="3" traits="%s">
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

	<property name="DescriptionKey" value="plantedCorn3_%s"/>
	<property name="CustomIcon" value="plantedCorn3Harvest"/>

	<drop event="Harvest" name="foodCropCorn" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="foodCropCorn" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<drop event="Destroy" name="plantedCorn1_%s" count="1" prob="0.5"/>
	%s
</block>`,
		traits,
		traits,
		traits,
		calculateCropYield(corn.CropYield, traits),
		calculateBonusYield(corn.BonusYield, traits),
		traits,
		optionallyAddRenewable(traits, corn))
	// TODO: <property name="CreativeMode" value="None"/>
}
