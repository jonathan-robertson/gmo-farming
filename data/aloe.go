package data

import "fmt"

// Aloe is a type of plant
type Aloe struct {
	Name               string
	NamePlural         string
	DisplayName        string
	Description        string
	CropYield          int
	BonusYield         int
	CraftTime          int
	incompatibleTraits []rune
}

// GetCraftTime returns the time required to craft this seed
func (p *Aloe) GetCraftTime() int {
	return p.CraftTime
}

// GetDescription returns the seed description for this plant
func (p *Aloe) GetDescription() string {
	if p.Description == "" {
		return getDefaultSeedDescription()
	}
	return p.Description
}

// GetDisplayName returns the display name
func (p *Aloe) GetDisplayName() string {
	return p.DisplayName
}

// GetName returns the name of this plant
func (p *Aloe) GetName() string {
	return p.Name
}

// GetSchematicName returns the schematic name for this plant, given the provided traits
func (p *Aloe) GetSchematicName(traits string) string {
	return fmt.Sprintf("plantedAloe1_%sschematic", traits)
}

// IsCompatibleWith checks for trait compatibility with this plant
func (p *Aloe) IsCompatibleWith(t Trait) bool {
	for _, incompatibleTrait := range p.incompatibleTraits {
		if incompatibleTrait == t.Code {
			return false
		}
	}
	return true
}

// WriteBlockStages produces each of the 3 block stages for this plant
func (p *Aloe) WriteBlockStages(c chan string, target, traits string) {
	p.writeStage1(c, target, traits)
	p.writeStage2(c, traits)
	p.writeStage3(c, traits)
}

func (p *Aloe) writeStage1(c chan string, target, traits string) {
	c <- fmt.Sprintf(`<block name="plantedAloe1_%s" stage="1" traits="%s">
    <drop event="Destroy" name="plantedAloe1_%s" count="1"/>
    <property name="CreativeMode" value="Player"/>
    <property name="CustomIcon" value="plantedAloe1"/>
    <property name="ItemTypeIcon" value="%s"/>
    <property name="DescriptionKey" value="plantedAloe1_%sDesc"/>
    <property name="Extends" value="cropsGrowingMaster" param1="CustomIcon"/>
    <property name="Group" value="%s"/>
    <property name="Material" value="Mcorn"/>
    <property name="Model" value="Entities/Plants/plant_aloe1_Prefab"/>
    <property name="PlaceAsRandomRotation" value="true"/>
    <property name="PlantGrowing.Next" value="plantedAloe2_%s"/>
    <property name="Shape" value="ModelEntity"/>
    <property name="UnlockedBy" value="%s"/>
</block>`, traits, traits, traits, getItemTypeIcon(traits), traits, getCraftingGroup(traits), traits, getUnlock(p, target, traits))
}

func (*Aloe) writeStage2(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedAloe2_%s" stage="2" traits="%s">
    <property name="CreativeMode" value="Dev"/>
    <property name="CustomIconTint" value="00ff80"/>
    <property name="Extends" value="plantedAloe1_%s"/>
    <property name="DescriptionKey" value="plantedAloe2"/>
    <property name="DisplayInfo" value="Description"/>
    <property name="Model" value="Entities/Plants/plant_aloe2_Prefab"/>
    <property name="PlantGrowing.Next" value="plantedAloe3_%s"/>
</block>`, traits, traits, traits, traits)
}

func (p *Aloe) writeStage3(c chan string, traits string) {
	c <- fmt.Sprintf(`<block name="plantedAloe3_%s" stage="3" traits="%s" tags="T%dPlant">
    <drop event="Destroy" count="plantedAloe1_%s" count="1" prob="0.5"/>
    <drop event="Fall" name="resourceYuccaFibers" count="0" prob="1" stick_chance="0"/>
    <drop event="Harvest" name="resourceCropAloeLeaf" count="%d" tag="cropHarvest"/>
    <drop event="Harvest" name="resourceCropAloeLeaf" prob="0.5" count="%d" tag="bonusCropHarvest"/>
    <property name="Collide" value="melee"/>
    <property name="CreativeMode" value="Dev"/>
    <property name="CustomIcon" value="plantedAloe1"/>
    <property name="ItemTypeIcon" value="%s"/>
    <property name="CustomIconTint" value="ff8000"/>
    <property name="DescriptionKey" value="plantedAloe3HarvestDesc"/>
    <property name="DisplayInfo" value="Description"/>
    <property name="DisplayType" value="blockMulti"/>
    <property name="FilterTags" value="MC_outdoor,SC_crops"/>
    <property name="Group" value="Food/Cooking,Science"/>
    <property name="HarvestOverdamage" value="false"/>
    <property name="ImposterDontBlock" value="true"/>
    <property name="IsDecoration" value="true"/>
    <property name="IsTerrainDecoration" value="true"/>
    <property name="LightOpacity" value="0"/>
    <property name="Material" value="Mcorn"/>
    <property name="Mesh" value="grass"/>
    <property name="Model" value="Entities/Plants/plant_aloe_harvestPrefab"/>
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
		calculatePlantTier(traits),
		traits,
		calculateCropYield(p.CropYield, traits),
		calculateBonusYield(p.BonusYield, traits),
		getItemTypeIcon(traits),
		optionallyAddRenewable(traits, p))
}
