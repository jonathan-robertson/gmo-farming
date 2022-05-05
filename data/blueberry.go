package data

import "fmt"

type Blueberry struct {
	Name              string
	NamePlural        string
	DisplayName       string
	Description       string
	PreferredConsumer string
	CropYield         int
	BonusYield        int
	CraftTime         int
}

func CreateBlueberry() *Blueberry {
	return &Blueberry{
		Name:              "Blueberry",
		DisplayName:       "Blueberry",
		PreferredConsumer: "Rabbits",
		CropYield:         2,
		BonusYield:        1,
		CraftTime:         2,
	}
}

func (p *Blueberry) GetCraftTime() int {
	return p.CraftTime
}

func (p *Blueberry) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

func (p *Blueberry) GetDisplayName() string {
	return p.DisplayName
}

func (p *Blueberry) GetName() string {
	return p.Name
}

func (p *Blueberry) GetPreferredConsumer() string {
	return p.PreferredConsumer
}

func (p *Blueberry) WriteBlockStages(c chan string, traits string) {
	p.WriteStage1(c, traits)
	p.WriteStage2(c, traits)
	p.WriteStage3(c, traits)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedBlueberry1Schematic"/>
func (*Blueberry) WriteStage1(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedBlueberry1_%s" stage="1" traits="%s">
	<drop event="Destroy" name="plantedBlueberry1_%s" count="1"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="plantedBlueberry1"/>
	<property name="DescriptionKey" value="plantedBlueberry1_%sDesc"/>
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
	<property name="Group" value="%s"/>
	<property name="Model" value="Entities/Plants/blueberry_plant_sproutPrefab"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="PlantGrowing.Next" value="plantedBlueberry2_%s"/>
	<property name="Shape" value="ModelEntity"/>
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits)
}

// TODO: <property name="CreativeMode" value="None"/>
func (*Blueberry) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedBlueberry2_%s" stage="2" traits="%s">
	<property name="Extends" value="plantedBlueberry1_%s"/>
	<property name="Model" value="Entities/Plants/blueberry_plant_growthPrefab"/>
	<property name="PlantGrowing.Next" value="plantedBlueberry3_%s"/>
</block>`, traits, traits, traits, traits)
}

// TODO: <property name="CreativeMode" value="None"/>
func (p *Blueberry) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedBlueberry3_%s" stage="3" traits="%s">
	<drop event="Destroy" name="plantedBlueberry1_%s" count="1" prob="0.5"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="foodCropBlueberries" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="foodCropBlueberries" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CustomIcon" value="plantedBlueberry1"/>
	<property name="CustomIconTint" value="ff8080"/>
	<property name="DescriptionKey" value="plantedBlueberry3_%s"/>
	<property name="DisplayInfo" value="Name"/>
	<property name="DisplayType" value="blockMulti"/>
	<property name="FilterTags" value="MC_helpers,SC_helperOutdoor"/>
	<property name="HarvestOverdamage" value="false"/>
	<property name="ImposterDontBlock" value="true"/>
	<property name="IsDecoration" value="true"/>
	<property name="IsTerrainDecoration" value="true"/>
	<property name="LightOpacity" value="0"/>
	<property name="Material" value="Mplants"/>
	<property name="Mesh" value="grass"/>
	<property name="Model" value="Entities/Plants/blueberry_plant_harvestPrefab"/>
	<property name="MultiBlockDim" value="1,2,1"/>
	<property name="PlantGrowing.FertileLevel" value="1"/>
	<property name="Shape" value="ModelEntity"/>
	<property name="SortOrder1" value="a090"/>
	<property name="SortOrder2" value="0002"/>
	<property name="Texture" value="395"/>
	%s
</block>`,
		traits,
		traits,
		traits,
		calculateCropYield(p.CropYield, traits),
		calculateBonusYield(p.BonusYield, traits),
		traits,
		optionallyAddRenewable(traits, p))
}
