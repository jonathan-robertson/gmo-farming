package data

import "fmt"

type Coffee struct {
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

func CreateCoffee() *Coffee {
	return &Coffee{
		Name:               "Coffee",
		DisplayName:        "Coffee",
		PreferredConsumer:  "",
		CropYield:          2,
		BonusYield:         1,
		CraftTime:          2,
		incompatibleTraits: []rune{'S'},
	}
}

func (p *Coffee) GetCraftTime() int {
	return p.CraftTime
}

func (p *Coffee) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

func (p *Coffee) GetDisplayName() string {
	return p.DisplayName
}

func (p *Coffee) GetName() string {
	return p.Name
}

func (p *Coffee) GetPreferredConsumer() string {
	return p.PreferredConsumer
}

func (p *Coffee) GetSchematicName(traits string) string {
	return fmt.Sprintf("plantedCoffee1_%sschematic", traits)
}

func (p *Coffee) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

func (p *Coffee) WriteBlockStages(c chan string, target, traits string) {
	p.WriteStage1(c, target, traits)
	p.WriteStage2(c, traits)
	p.WriteStage3(c, traits)
}

func (p *Coffee) WriteStage1(c chan string, target, traits string) {
	c <- fmt.Sprintf(`<block name="plantedCoffee1_%s" stage="1" traits="%s">
	<drop event="Destroy" name="plantedCoffee1_%s" count="1"/>
	<property name="CraftingIngredientTime" value="5"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="plantedCoffee1"/>
	<property name="DescriptionKey" value="plantedCoffee1_%sDesc"/>
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
	<property name="Group" value="%s"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="PlantGrowing.Next" value="plantedCoffee2_%s"/>
	<property name="Texture" value="393"/>
	%s
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits, optionallyAddUnlock(p, target, traits))
}

func (*Coffee) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedCoffee2_%s" stage="2" traits="%s">
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIconTint" value="00ff80"/>
	<property name="Extends" value="plantedCoffee1_%s"/>
	<property name="PlantGrowing.Next" value="plantedCoffee3_%s"/>
	<property name="Texture" value="394"/>
</block>`, traits, traits, traits, traits)
}

func (p *Coffee) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedCoffee3_%s" stage="3" traits="%s">
	<drop event="Destroy" name="plantedCoffee1_%s" count="1" prob="0.5"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="resourceCropCoffeeBeans" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="resourceCropCoffeeBeans" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIcon" value="plantedCoffee3Harvest"/>
	<property name="CustomIconTint" value="ff8000"/>
	<property name="DescriptionKey" value="plantedCoffee3_%s"/>
	<property name="DisplayInfo" value="Description"/> <!-- also valid: "Name" -->
	<property name="DisplayType" value="blockMulti"/>
	<property name="FilterTags" value="MC_outdoor,SC_crops"/>
	<property name="HarvestOverdamage" value="false"/>
	<property name="ImposterDontBlock" value="true"/>
	<property name="IsDecoration" value="true"/>
	<property name="IsTerrainDecoration" value="true"/>
	<property name="LightOpacity" value="0"/>
	<property name="Material" value="Mplants"/>
	<property name="Mesh" value="grass"/>
	<property name="MultiBlockDim" value="1,2,1"/>
	<property name="PlantGrowing.FertileLevel" value="1"/>
	<property name="Shape" value="BillboardPlant"/>
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
