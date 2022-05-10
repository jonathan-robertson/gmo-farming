package data

import "fmt"

type Cotton struct {
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

func CreateCotton() *Cotton {
	return &Cotton{
		Name:               "Cotton",
		DisplayName:        "Cotton",
		PreferredConsumer:  "",
		CropYield:          2,
		BonusYield:         1,
		CraftTime:          2,
	}
}

func (p *Cotton) GetCraftTime() int {
	return p.CraftTime
}

func (p *Cotton) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

func (p *Cotton) GetDisplayName() string {
	return p.DisplayName
}

func (p *Cotton) GetName() string {
	return p.Name
}

func (p *Cotton) GetPreferredConsumer() string {
	return p.PreferredConsumer
}

func (p *Cotton) GetSchematicName(traits string) string {
	return fmt.Sprintf("plantedCotton1_%sschematic", traits)
}

func (p *Cotton) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

func (p *Cotton) WriteBlockStages(c chan string, target, traits string) {
	p.WriteStage1(c, target, traits)
	p.WriteStage2(c, traits)
	p.WriteStage3(c, traits)
}

func (p *Cotton) WriteStage1(c chan string, target, traits string) {
	c <- fmt.Sprintf(`<block name="plantedCotton1_%s" stage="1" traits="%s">
	<drop event="Destroy" name="plantedCotton1_%s" count="1"/>
	<property name="CreativeMode" value="Player"/>
	<property name="CustomIcon" value="plantedCotton1"/>
	<property name="DescriptionKey" value="plantedCotton1_%sDesc"/>
	<property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
	<property name="Group" value="%s"/>
	<property name="PlaceAsRandomRotation" value="true"/>
	<property name="PlantGrowing.Next" value="plantedCotton2_%s"/>
	<property name="Texture" value="392"/>
	%s
</block>`, traits, traits, traits, traits, getCraftingGroup(traits), traits, optionallyAddUnlock(p, target, traits))
}

func (*Cotton) WriteStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedCotton2_%s" stage="2" traits="%s">
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIconTint" value="00ff80"/>
	<property name="Extends" value="plantedCotton1_%s"/>
	<property name="PlantGrowing.Next" value="plantedCotton3_%s"/>
	<property name="Texture" value="20"/>
</block>`, traits, traits, traits, traits)
}

func (p *Cotton) WriteStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedCotton3_%s" stage="3" traits="%s" tags="T%dPlant">
	<drop event="Destroy" name="plantedCotton1_%s" count="1" prob="0.5"/>
	<drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
	<drop event="Harvest" name="resourceCropCottonPlant" count="%d" tag="cropHarvest"/>
	<drop event="Harvest" name="resourceCropCottonPlant" prob="0.5" count="%d" tag="bonusCropHarvest"/>
	<property name="Collide" value="melee"/>
	<property name="CreativeMode" value="Dev"/>
	<property name="CustomIcon" value="plantedCotton1"/>
	<property name="CustomIconTint" value="ff8000"/>
	<property name="DescriptionKey" value="plantedCotton3_%s"/>
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
	<property name="Texture" value="363"/>
	%s
</block>`,
		traits,
		traits,
		calculatePlantTier(traits),
		traits,
		calculateCropYield(p.CropYield, traits),
		calculateBonusYield(p.BonusYield, traits),
		traits,
		optionallyAddRenewable(traits, p))
}
