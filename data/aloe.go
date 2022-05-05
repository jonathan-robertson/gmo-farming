package data

import "fmt"

type Aloe struct {
	Name               string
	NamePlural         string
	DisplayName        string
	Description        string
	PreferredConsumer  string
	CropYield          int
	BonusYield         int
	CraftTime          int
	incompatibleTraits []rune
}

func CreateAloe() *Aloe {
	return &Aloe{
		Name:               "Aloe",
		DisplayName:        "Aloe",
		PreferredConsumer:  "",
		CropYield:          2,
		BonusYield:         1,
		CraftTime:          2,
		incompatibleTraits: []rune{'S'},
	}
}

func (p *Aloe) GetCraftTime() int {
	return p.CraftTime
}

func (p *Aloe) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

func (p *Aloe) GetDisplayName() string {
	return p.DisplayName
}

func (p *Aloe) GetName() string {
	return p.Name
}

func (p *Aloe) GetPreferredConsumer() string {
	return p.PreferredConsumer
}

func (p *Aloe) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

func (p *Aloe) WriteBlockStages(c chan string, traits string) {
	p.WriteStage1(c, traits)
	p.WriteStage2(c, traits)
	p.WriteStage3(c, traits)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedAloe1Schematic"/>
func (*Aloe) WriteStage1(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedAloe1_%s" stage="1" traits="%s">
	<drop event="Destroy" name="plantedAloe1_%s" count="1"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="plantedAloe1"/>
	<property name="DescriptionKey" value="plantedAloe1_%sDesc"/>
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
	<property name="Group" value="%s"/>
	<property name="Material" value="Mcorn"/>
	<property name="Model" value="Entities/Plants/plant_aloe1_Prefab"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="PlantGrowing.Next" value="plantedAloe2_%s"/>
	<property name="Shape" value="ModelEntity"/>
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits)
}

// TODO: <property name="CreativeMode" value="None"/>
func (*Aloe) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedAloe2_%s" stage="2" traits="%s">
	<property name="Extends" value="plantedAloe1_%s"/>
	<property name="Model" value="Entities/Plants/plant_aloe2_Prefab"/>
	<property name="PlantGrowing.Next" value="plantedAloe3_%s"/>
</block>`, traits, traits, traits, traits)
}

// TODO: <property name="CreativeMode" value="None"/>
func (p *Aloe) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedAloe3_%s" stage="3" traits="%s">
	<drop event="Destroy" name="plantedCorn1_%s" count="1" prob="0.5"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="foodCropCorn" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="foodCropCorn" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CustomIcon" value="plantedCorn1"/>
	<property name="DescriptionKey" value="plantedCorn3_%s"/>
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
	<property name="Texture" value="529"/>
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
