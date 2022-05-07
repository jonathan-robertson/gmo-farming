package data

import "fmt"

type Potato struct {
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

func CreatePotato() *Potato {
	return &Potato{
		Name:               "Potato",
		DisplayName:        "Potato",
		PreferredConsumer:  "",
		CropYield:          2,
		BonusYield:         1,
		CraftTime:          2,
	}
}

func (p *Potato) GetCraftTime() int {
	return p.CraftTime
}

func (p *Potato) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

func (p *Potato) GetDisplayName() string {
	return p.DisplayName
}

func (p *Potato) GetName() string {
	return p.Name
}

func (p *Potato) GetPreferredConsumer() string {
	return p.PreferredConsumer
}

func (p *Potato) GetSchematicName(traits string) string {
	return fmt.Sprintf("plantedPotato1_%sschematic", traits)
}

func (p *Potato) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

func (p *Potato) WriteBlockStages(c chan string, target, traits string) {
	p.WriteStage1(c, target, traits)
	p.WriteStage2(c, traits)
	p.WriteStage3(c, traits)
}

func (p *Potato) WriteStage1(c chan string, target, traits string) {
	c <- fmt.Sprintf(`<block name="plantedPotato1_%s" stage="1" traits="%s">
	<drop event="Destroy" name="plantedPotato1_%s" count="1"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="plantedPotato1"/>
	<property name="DescriptionKey" value="plantedPotato1_%sDesc"/>
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
	<property name="Group" value="%s"/>
	<property name="Model" value="Entities/Plants/potato_plant_sproutPrefab"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="PlantGrowing.Next" value="plantedPotato2_%s"/>
	<property name="Shape" value="ModelEntity"/>
	%s
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits, optionallyAddUnlock(p, target, traits))
}

func (*Potato) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedPotato2_%s" stage="2" traits="%s">
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIconTint" value="00ff80"/>
	<property name="Extends" value="plantedPotato1_%s"/>
	<property name="Model" value="Entities/Plants/potato_plant_growthPrefab"/>
	<property name="PlantGrowing.Next" value="plantedPotato3_%s"/>
</block>`, traits, traits, traits, traits)
}

func (p *Potato) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedPotato3_%s" stage="3" traits="%s">
	<drop event="Destroy" name="plantedPotato1_%s" count="1" prob="0.5"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="foodCropPotato" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="foodCropPotato" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIcon" value="plantedPotato1"/>
	<property name="CustomIconTint" value="ff8000"/>
	<property name="DescriptionKey" value="plantedPotato3_%sDesc"/>
	<property name="DisplayInfo" value="Description"/> <!-- also valid: "Name" -->
	<property name="DisplayType" value="blockMulti"/>
	<property name="Extends" value="cropsHarvestableMaster"/>
	<property name="FilterTags" value="MC_outdoor,SC_crops"/>
	<property name="HarvestOverdamage" value="false"/>
	<property name="ImposterDontBlock" value="true"/>
	<property name="IsDecoration" value="true"/>
	<property name="IsTerrainDecoration" value="true"/>
	<property name="LightOpacity" value="0"/>
	<property name="Material" value="Mplants"/>
	<property name="Mesh" value="grass"/>
	<property name="Model" value="Entities/Plants/potato_plant_harvestPrefab"/>
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
