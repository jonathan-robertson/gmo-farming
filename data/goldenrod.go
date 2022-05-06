package data

import "fmt"

type Goldenrod struct {
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

func CreateGoldenrod() *Goldenrod {
	return &Goldenrod{
		Name:               "Goldenrod",
		DisplayName:        "Goldenrod",
		PreferredConsumer:  "Stags",
		CropYield:          2,
		BonusYield:         1,
		CraftTime:          2,
		incompatibleTraits: []rune{},
	}
}

func (p *Goldenrod) GetCraftTime() int {
	return p.CraftTime
}

func (p *Goldenrod) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

func (p *Goldenrod) GetDisplayName() string {
	return p.DisplayName
}

func (p *Goldenrod) GetName() string {
	return p.Name
}

func (p *Goldenrod) GetPreferredConsumer() string {
	return p.PreferredConsumer
}

func (p *Goldenrod) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

func (p *Goldenrod) WriteBlockStages(c chan string, traits string) {
	p.WriteStage1(c, traits)
	p.WriteStage2(c, traits)
	p.WriteStage3(c, traits)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedGoldenrod1Schematic"/>
func (*Goldenrod) WriteStage1(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedGoldenrod1_%s" stage="1" traits="%s">
	<drop event="Destroy" name="plantedGoldenrod1_%s" count="1"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="plantedGoldenrod1"/>
	<property name="DescriptionKey" value="plantedGoldenrod1_%sDesc"/>
	<property name="Extends" value="cropsGrowingMaster"/>
	<property name="Group" value="%s"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="PlantGrowing.Next" value="plantedGoldenrod2_%s"/>
	<property name="Texture" value="401"/>
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits)
}

func (*Goldenrod) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedGoldenrod2_%s" stage="2" traits="%s">
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIconTint" value="00ff80"/>
	<property name="Extends" value="plantedGoldenrod1_%s"/>
	<property name="PlantGrowing.Next" value="plantedGoldenrod3_%s"/>
	<property name="Texture" value="402"/>
</block>`, traits, traits, traits, traits)
}

func (p *Goldenrod) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedGoldenrod3_%s" stage="3" traits="%s">
	<drop event="Destroy" name="plantedGoldenrod1_%s" count="1" prob="0.5"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="resourceCropGoldenrodPlant" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="resourceCropGoldenrodPlant" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIcon" value="plantedGoldenrod3HarvestPlayer"/>
	<property name="CustomIconTint" value="ff8000"/>
	<property name="DescriptionKey" value="plantedGoldenrod3_%s"/>
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
	<property name="Texture" value="362"/>
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
