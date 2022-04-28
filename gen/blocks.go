package gen

import (
	"fmt"
)

func WritePlantBlocks() error {
	file, err := getFile("Config/blocks.xml")
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go producePlantBlocks(c)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func producePlantBlocks(c chan string) {
	defer close(c)
	c <- `<config>`
	c <- `<append xpath="/blocks">`
	produceWorkstationHotBox(c)
	for _, plant := range Plants {
		for _, tier := range []int{2, 3} {
			// produce T2, T3 with no traits
			plant.WriteBlockStages(c, tier, "")
			for i1 := 0; i1 < len(Traits); i1++ {
				switch tier {
				case 2:
					traits := fmt.Sprintf("%c", Traits[i1].code)
					if plant.IsCompatibleWith(traits) {
						plant.WriteBlockStages(c, tier, traits)
					}
				case 3:
					for i2 := i1; i2 < len(Traits); i2++ {
						if Traits[i1].isCompatibleWith(Traits[i2]) {
							traits := fmt.Sprintf("%c%c", Traits[i1].code, Traits[i2].code)
							if plant.IsCompatibleWith(traits) {
								plant.WriteBlockStages(c, tier, traits)
							}
						}
					}
				}
			}
		}
	}
	c <- `</append>`
	produceBlockModifications(c)
	c <- `</config>`
}

// TODO: add hinge sound? Maybe similar to secure_storage or drawer?
// <property name="CustomIcon" value="farmPlotBlockRaised"/>
/*
	<property name="Extends" value="workbench"/>
	<property name="CustomIcon" value="shapeFarmPlotRaised"/>
	<property name="Shape" value="New"/>
	<property name="Model" value="farm_plot_raised"/>
	<property name="ModelOffset" value="1,0,1"/>
	<property name="Texture" value="194,194,194,194,194,194"/>
	<property name="UseGlobalUV" value="G,L,L,L,L,L"/>
	<property name="DisplayType" value="blockMulti"/>
	<property name="MultiBlockDim" value="1,1,1"/>
	<property class="Workstation">
		<property name="Modules" value="output"/>
		<property name="CraftingAreaRecipes" value="hotbox"/>
	</property>

	<property name="OpenSound" value="campfire_open"/>
	<property name="CloseSound" value="campfire_close"/>
	<property name="CraftSound" value="campfire_cook_click"/>

	<property class="RepairItems">
		<property name="resourceRockSmall" value="5"/>
	</property>


	<!--<property name="Group" value="Food/Cooking"/>-->
	<property name="FilterTags" value="MC_playerBlocks,SC_decor"/>


	<property name="StabilitySupport" value="true"/>
	<property name="DescriptionKey" value="campfireDesc"/>
*/
func produceWorkstationHotBox(c chan string) {
	/*
			c <- `<block name="hotbox">
			<property name="Extends" value="workbench"/>
			<property name="CustomIcon" value="shapeFarmPlotRaised"/>
			<property name="CreativeMode" value="Player"/>
			<property name="Material" value="Mmetal"/>
			<property name="Texture" value="194,194,194,194,194,194"/>
			<property name="UseGlobalUV" value="G,L,L,L,L,L"/>
			<property name="Shape" value="New"/>
			<property name="Model" value="farm_plot_raised"/>
			<property name="FilterTags" value="MC_playerBlocks,SC_decor"/>

			<drop event="Destroy" name="resourceClayLump" count="0" tag="oreWoodHarvest"/>
			<drop event="Destroy" name="resourcePotassiumNitratePowder" count="0" tag="oreWoodHarvest"/>
			<property name="ModelOffset" value="1,0,1"/>
			<property name="DisplayType" value="blockMulti"/>
			<property name="MultiBlockDim" value="1,1,1"/>
			<property class="Workstation">
				<property name="Modules" value="output"/>
				<property name="CraftingAreaRecipes" value="hotbox"/>
			</property>
			<property name="OpenSound" value="campfire_open"/>
			<property name="CloseSound" value="campfire_close"/>
			<property name="CraftSound" value="campfire_cook_click"/>
			<property name="StabilitySupport" value="true"/>
		</block>`
	*/
	// Shavick's imported version
	c <- `<block name="hotbox">
			<property name="Extends" value="workbench"/>
			<property name="CustomIcon" value="shapeFarmPlotRaised"/>
			<property name="Shape" value="New"/>
			<property name="Model" value="farm_plot_raised"/>
			<property name="ModelOffset" value="1,0,1"/>
			<property name="Texture" value="194,194,194,194,194,194"/>
			<property name="UseGlobalUV" value="G,L,L,L,L,L"/>
			<property name="DisplayType" value="blockMulti"/>
			<property name="MultiBlockDim" value="1,1,1"/>
			<property class="Workstation">
				<property name="Modules" value="output"/>
				<property name="CraftingAreaRecipes" value="hotbox"/>
			</property>
        </block>`
}

func produceBlockModifications(c chan string) {
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

	// TODO: [T] Thorny
	// TODO: [S] Sweet
}
