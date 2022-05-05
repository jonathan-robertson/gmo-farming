package data

import "fmt"

type GraceCorn struct {
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

func CreateGraceCorn() *GraceCorn {
	return &GraceCorn{
		Name:               "GraceCorn",
		DisplayName:        "Super Corn",
		PreferredConsumer:  "",
		CropYield:          2,
		BonusYield:         1,
		CraftTime:          2,
		incompatibleTraits: []rune{'S'},
	}
}

func (p *GraceCorn) GetCraftTime() int {
	return p.CraftTime
}

func (p *GraceCorn) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

func (p *GraceCorn) GetDisplayName() string {
	return p.DisplayName
}

func (p *GraceCorn) GetName() string {
	return p.Name
}

func (p *GraceCorn) GetPreferredConsumer() string {
	return p.PreferredConsumer
}

func (p *GraceCorn) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

func (p *GraceCorn) WriteBlockStages(c chan string, traits string) {
	p.WriteStage1(c, traits)
	p.WriteStage2(c, traits)
	p.WriteStage3(c, traits)
}

// TODO: <property name="UnlockedBy" value="perkLivingOffTheLand,plantedGraceCorn1Schematic"/>
func (*GraceCorn) WriteStage1(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedGraceCorn1_%s" stage="1" traits="%s">
	<drop event="Destroy" name="plantedGraceCorn1_%s" count="1"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="plantedCorn1"/>
	<property name="CustomIconTint" value="ff9f9f"/>
	<property name="DescriptionKey" value="plantedGraceCorn1_%sDesc"/>
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
	<property name="Group" value="%s"/>
	<property name="Material" value="Mcorn"/> <!-- mostly for the particle effect -->
	<property name="Mesh" value="cutoutmoveable"/>
	<property name="Model" value="corn_sprout_shape"/>
	<property name="MultiBlockDim" value="1,3,1"/>
	<property name="Place" value="Door"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="PlantGrowing.Next" value="plantedGraceCorn2_%s"/>
	<property name="Shape" value="New"/>
	<property name="Texture" value="529"/>
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits)
}

func (*GraceCorn) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedGraceCorn2_%s" stage="2" traits="%s">
	<property name="CreativeMode" value="Dev"/>
	<property name="Extends" value="plantedGraceCorn1_%s"/>
	<property name="PlantGrowing.Next" value="plantedGraceCorn3_%s"/>
</block>`, traits, traits, traits, traits)
}

func (p *GraceCorn) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedGraceCorn3_%s" stage="3" traits="%s">
	<drop event="Destroy" name="plantedGraceCorn1_%s" count="1" prob="0.5"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="foodCropGraceCorn" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="foodCropGraceCorn" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIcon" value="plantedCorn1"/>
	<property name="CustomIconTint" value="ff9f9f"/>
	<property name="DescriptionKey" value="plantedGraceCorn3_%s"/>
	<property name="DisplayInfo" value="Description"/> <!-- also valid: "Name" -->
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
