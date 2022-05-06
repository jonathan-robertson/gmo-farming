package data

import "fmt"

type Pumpkin struct {
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

func CreatePumpkin() *Pumpkin {
	return &Pumpkin{
		Name:               "Pumpkin",
		DisplayName:        "Pumpkin",
		PreferredConsumer:  "",
		CropYield:          2,
		BonusYield:         1,
		CraftTime:          2,
		incompatibleTraits: []rune{'S'},
	}
}

func (p *Pumpkin) GetCraftTime() int {
	return p.CraftTime
}

func (p *Pumpkin) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

func (p *Pumpkin) GetDisplayName() string {
	return p.DisplayName
}

func (p *Pumpkin) GetName() string {
	return p.Name
}

func (p *Pumpkin) GetPreferredConsumer() string {
	return p.PreferredConsumer
}

func (p *Pumpkin) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

func (p *Pumpkin) WriteBlockStages(c chan string, traits string) {
	p.WriteStage1(c, traits)
	p.WriteStage2(c, traits)
	p.WriteStage3(c, traits)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedPumpkin1Schematic"/>
func (*Pumpkin) WriteStage1(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedPumpkin1_%s" stage="1" traits="%s">
	<drop event="Destroy" name="plantedPumpkin1_%s" count="1"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="plantedPumpkin1"/>
	<property name="DescriptionKey" value="plantedPumpkin1_%s"/>
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
	<property name="Group" value="%s"/>
	<property name="Material" value="Mcorn"/>
	<property name="Model" value="Entities/Plants/pumpkinSproutPrefab"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="PlantGrowing.Next" value="plantedPumpkin2_%s"/>
	<property name="Shape" value="ModelEntity"/>
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits)
}

func (*Pumpkin) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedPumpkin2_%s" stage="2" traits="%s">
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIconTint" value="00ff80"/>
	<property name="Extends" value="plantedPumpkin1_%s"/>
	<property name="Model" value="Entities/Plants/pumpkinGrowthPrefab"/>
	<property name="PlantGrowing.Next" value="plantedPumpkin3_%s"/>
</block>`, traits, traits, traits, traits)
}

func (p *Pumpkin) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedPumpkin3_%s" stage="3" traits="%s">
	<drop event="Destroy" name="plantedPumpkin1_%s" count="1" prob="0.5"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="foodCropPumpkin" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="foodCropPumpkin" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIcon" value="plantedPumpkin3HarvestPlayer"/>
	<property name="CustomIconTint" value="ff8000"/>
	<property name="DescriptionKey" value="plantedPumpkin3_%s"/>
	<property name="DisplayInfo" value="Description"/> <!-- also valid: "Name" -->
	<property name="DisplayType" value="blockMulti"/>
	<property name="Extends" value="cropsHarvestableMaster"/>
	<property name="FilterTags" value="MC_outdoor,SC_crops"/>
	<property name="HarvestOverdamage" value="false"/>
	<property name="ImposterDontBlock" value="true"/>
	<property name="IsDecoration" value="true"/>
	<property name="IsTerrainDecoration" value="true"/>
	<property name="LightOpacity" value="0"/>
	<property name="Material" value="Mcorn"/>
	<property name="Material" value="Mplants"/>
	<property name="Mesh" value="grass"/>
	<property name="Model" value="Entities/Plants/pumpkinHarvestPrefab"/>
	<property name="MultiBlockDim" value="1,2,1"/>
	<property name="PlantGrowing.FertileLevel" value="1"/>
	<property name="Shape" value="BillboardPlant"/>
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
