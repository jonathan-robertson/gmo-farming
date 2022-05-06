package data

import "fmt"

type Chrysanthemum struct {
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

func CreateChrysanthemum() *Chrysanthemum {
	return &Chrysanthemum{
		Name:               "Chrysanthemum",
		DisplayName:        "Chrysanthemum",
		PreferredConsumer:  "Does",
		CropYield:          2,
		BonusYield:         1,
		CraftTime:          2,
		incompatibleTraits: []rune{},
	}
}

func (p *Chrysanthemum) GetCraftTime() int {
	return p.CraftTime
}

func (p *Chrysanthemum) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

func (p *Chrysanthemum) GetDisplayName() string {
	return p.DisplayName
}

func (p *Chrysanthemum) GetName() string {
	return p.Name
}

func (p *Chrysanthemum) GetPreferredConsumer() string {
	return p.PreferredConsumer
}

func (p *Chrysanthemum) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

func (p *Chrysanthemum) WriteBlockStages(c chan string, traits string) {
	p.WriteStage1(c, traits)
	p.WriteStage2(c, traits)
	p.WriteStage3(c, traits)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedChrysanthemum1Schematic"/>
func (*Chrysanthemum) WriteStage1(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedChrysanthemum1_%s" stage="1" traits="%s">
	<drop event="Destroy" name="plantedChrysanthemum1_%s" count="1"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="plantedChrysanthemum1"/>
	<property name="DescriptionKey" value="plantedChrysanthemum1_%sDesc"/>
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
	<property name="Group" value="%s"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="PlantGrowing.Next" value="plantedChrysanthemum2_%s"/>
	<property name="Texture" value="550"/>
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits)
}

func (*Chrysanthemum) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedChrysanthemum2_%s" stage="2" traits="%s">
	<property name="Extends" value="plantedChrysanthemum1_%s"/>
	<property name="Texture" value="551"/>
	<property name="CreativeMode" value="Dev"/>
	<property name="PlantGrowing.Next" value="plantedChrysanthemum3_%s"/>
</block>`, traits, traits, traits, traits)
}

func (p *Chrysanthemum) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedChrysanthemum3_%s" stage="3" traits="%s">
	<drop event="Destroy" name="plantedChrysanthemum1_%s" count="1" prob="0.5"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="resourceCropChrysanthemumPlant" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="resourceCropChrysanthemumPlant" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIcon" value="plantedChrysanthemum1"/>
	<property name="DescriptionKey" value="plantedChrysanthemum3_%s"/>
	<property name="DisplayInfo" value="Name"/>
	<property name="DisplayType" value="blockMulti"/>
	<property name="FilterTags" value="MC_helpers,SC_helperOutdoor"/>
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
	<property name="Texture" value="244"/>
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
