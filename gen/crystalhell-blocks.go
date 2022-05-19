package gen

import (
	"data"
	"fmt"
)

// CrystalHellBlocks is responsible for producing content for blocks.xml
type CrystalHellBlocks struct{}

// GetPath returns file path for this producer
func (*CrystalHellBlocks) GetPath() string {
	return "Config-CrystalHell"
}

// GetFilename returns filename for this producer
func (*CrystalHellBlocks) GetFilename() string {
	return "blocks.xml"
}

// Produce xml data to the provided channel
func (p *CrystalHellBlocks) Produce(c chan string) {
	defer close(c)
	c <- `<config>`
	c <- `<append xpath="/blocks">`
	p.produceWorkstationHotBox(c)
	for _, plant := range data.Plants {
		plant.WriteBlockStages(c, p.getTarget(), "")
		for _, trait1 := range data.Traits {
			if plant.IsCompatibleWith(trait1) {
				plant.WriteBlockStages(c, p.getTarget(), fmt.Sprintf("%c", trait1.Code))
				for _, trait2 := range data.Traits {
					if trait1.IsCompatibleWith(trait2) && plant.IsCompatibleWith(trait2) {
						plant.WriteBlockStages(c, p.getTarget(), fmt.Sprintf("%c%c", trait1.Code, trait2.Code))
					}
				}
			}
		}
	}
	c <- `</append>`
	p.produceBlockModifications(c)
	c <- `</config>`
}

func (*CrystalHellBlocks) produceWorkstationHotBox(c chan string) {
	c <- `<block name="hotbox">
	<property name="Extends" value="workbench"/>
	<property class="Workstation">
		<property name="Modules" value="output"/>
		<property name="CraftingAreaRecipes" value="hotbox"/>
	</property>
	
	<property name="ModelOffset" value="1,0,1"/>
	<property name="MultiBlockDim" value="1,1,1"/>
	<property name="StabilitySupport" value="true"/>
	<property name="IsTerrainDecoration" value="false"/>
	<property name="DisplayType" value="blockMulti"/>
	
	<!-- from farmPlotRaised shape -->
	<property name="Path" value="solid"/>  <!-- This is a hint for the AI; see XML.txt -->
	<property name="Shape" value="New"/>
	<property name="Model" value="farm_plot_raised"/>
	<property name="CustomIcon" value="shapeFarmPlotRaised"/>
	<property name="ImposterExchange" value="imposterBlock"/>
	<property name="UseGlobalUV" value="G,L,L,L,L,L"/>

	<!-- from corrugatedMetalShapes -->
	<property name="Texture" value="194"/>
	<property name="UiBackgroundTexture" value="194"/>

	<!-- audio -->
	<property name="OpenSound" value="drone_storage_open"/>
	<property name="CloseSound" value="drone_storage_close"/>
	
	<!-- Localization -->
	<property name="WorkstationJournalTip" value="hotboxTip"/>
	<property name="DescriptionKey" value="hotboxDesc"/>

	<!-- recipe/unlock -->
	<property name="UnlockedBy" value="perkAdvancedEngineering,workbenchSchematic"/>
	<!-- TODO: use these ingredients for the recipe since they'll drop (thanks, workbench)
		<drop event="Harvest" name="resourceScrapIron" count="200" tag="allHarvest"/>
		<drop event="Harvest" name="resourceWood" count="20" tag="allHarvest"/>
		<drop event="Harvest" name="terrStone" count="0" tool_category="Disassemble"/>
		<drop event="Harvest" name="resourceForgedIron" count="10" tag="salvageHarvest"/>
		<drop event="Harvest" name="resourceMechanicalParts" count="8" tag="salvageHarvest"/>
		<drop event="Harvest" name="resourceWood" count="20" tag="salvageHarvest"/>
	-->
	<property name="EconomicValue" value="2000"/>

	<!-- TODO: Heat -->
	<property name="HeatMapStrength" value="2"/>
	<property name="HeatMapTime" value="5000"/>
	<property name="HeatMapFrequency" value="1000"/>

	<!-- TODO: Other -->
	<property name="TakeDelay" value="5"/>
	<property name="WorkstationIcon" value="ui_game_symbol_workbench"/>
</block>`
}

func (*CrystalHellBlocks) produceBlockModifications(c chan string) {
	// [U] Underground
	c <- `    <append xpath="/blocks/block[contains(@traits, 'U') and @stage='1']">
        <property name="PlantGrowing.LightLevelGrow" value="0" />
        <property name="PlantGrowing.LightLevelStay" value="0" />
    </append>`

	// [F] Fast
	c <- `    <append xpath="/blocks/block[contains(@traits, 'F') and @stage='1' and not (@traits='FF')]">
        <property name="PlantGrowing.GrowthRate" value="31.5" />
    </append>`
	c <- `    <append xpath="/blocks/block[@traits='FF' and @stage='1']">
        <property name="PlantGrowing.GrowthRate" value="15.75" />
    </append>`

	// [E] Explosive: based off of mineCookingPot with reduced trigger delay from .5 to .1
	c <- `    <append xpath="/blocks/block[contains(@traits, 'E') and @stage='3' and not (@traits='EE')]">
        <property name="Class" value="Mine" /> <!-- a mine destroyed by an *explosion* only has a 33 percent chance to detonate -->
        <property name="Tags" value="Mine" />
        <property name="Collide" value="movement,melee,arrow" />
		<property name="MaxDamage" value="1" /> <!-- reduced from 4 -->
        <property name="TriggerDelay" value="0.1" /> <!-- reduced from 0.5 -->
        <property name="TriggerSound" value="landmine_trigger" />
        <property name="Explosion.ParticleIndex" value="11" />
        <property name="Explosion.RadiusEntities" value="3" />
        <property name="Explosion.EntityDamage" value="300" /> <!-- damage for entities in the center of the explosion. -->
        <property name="CanPickup" value="false" />
    </append>`
	// [EE] Explosive: based off of mineHubcap with reduced trigger delay from .5 to .1
	c <- `    <append xpath="/blocks/block[contains(@traits, 'EE') and @stage='3']">
        <property name="Class" value="Mine" /> <!-- a mine destroyed by an *explosion* only has a 33 percent chance to detonate -->
        <property name="Tags" value="Mine" />
        <property name="Collide" value="movement,melee,arrow" />
		<property name="MaxDamage" value="1" /> <!-- reduced from 4 -->
        <property name="TriggerDelay" value="0.1" /> <!-- reduced from 0.5 -->
        <property name="TriggerSound" value="landmine_trigger" />
        <property name="Explosion.ParticleIndex" value="11" />
        <property name="Explosion.RadiusEntities" value="5" />
        <property name="Explosion.EntityDamage" value="450" /> <!-- damage for entities in the center of the explosion. -->
        <property name="CanPickup" value="false" />
    </append>`
	// [T] Thorny
	c <- `    <append xpath="/blocks/block[contains(@traits, 'T') and @stage='3' and not (@traits='TT')]">
        <property name="BuffsWhenWalkedOn" value="triggerInjuryThorns"/>
    </append>`
	// [TT] Extra Thorny
	c <- `    <append xpath="/blocks/block[contains(@traits, 'TT') and @stage='3']">
        <property name="BuffsWhenWalkedOn" value="triggerInjuryCriticalThorns"/>
    </append>`
}

func (p *CrystalHellBlocks) getTarget() string {
	return "CrystalHell"
}
