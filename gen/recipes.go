package gen

import (
	"fmt"
)

func WritePlantRecipes() error {
	file, err := getFile("Config/recipes.xml")
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go producePlantRecipes(c)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func producePlantRecipes(c chan string) {
	defer close(c)
	c <- `<config>`
	c <- `<append xpath="/recipes">`
	for _, plant := range Plants {
		for _, tier := range []int{2, 3} {
			// produce T2, T3 with no traits
			produceRecipeStub(c, plant.GetName(), tier, "", plant.GetCraftTime())
			for i1 := 0; i1 < len(Traits); i1++ {
				switch tier {
				case 2:
					traits := fmt.Sprintf("%c", Traits[i1].code)
					if plant.IsCompatibleWith(traits) {
						produceRecipeStub(c, plant.GetName(), tier, traits, plant.GetCraftTime())
					}
				case 3:
					for i2 := i1; i2 < len(Traits); i2++ {
						if Traits[i1].isCompatibleWith(Traits[i2]) {
							traits := fmt.Sprintf("%c%c", Traits[i1].code, Traits[i2].code)
							if plant.IsCompatibleWith(traits) {
								produceRecipeStub(c, plant.GetName(), tier, traits, plant.GetCraftTime())
							}
						}
					}
				}
			}
		}
	}
	c <- `</append>`
	produceRecipeModifications(c)
	c <- `</config>`
}

func produceRecipeStub(c chan string, name string, tier int, traits string, craftTime int) {
	var ingredientTier, optionalCraftingArea string
	switch true {
	case tier == 2 && traits == "":
		ingredientTier = ""
		optionalCraftingArea = ` craft_area="hotbox"`
	case tier == 3 && traits == "":
		ingredientTier = "T2"
		optionalCraftingArea = ` craft_area="hotbox"`
	default:
		ingredientTier = fmt.Sprintf("T%d", tier)
		optionalCraftingArea = ``
	}
	// TODO: tags="learnable"
	c <- fmt.Sprintf(`<recipe name="planted%s1T%d%s" count="1" craft_time="%d" tier="%d" traits="%s"%s>
    <ingredient name="planted%s1%s" count="1"/>
</recipe>`,
		name,
		tier,
		traits,
		calculateCraftTime(craftTime, tier, traits),
		tier,
		traits,
		optionalCraftingArea,
		name,
		ingredientTier)
}

func produceRecipeModifications(c chan string) {
	// Tier 2 Upgrade
	c <- `    <append xpath="/recipes/recipe[@tier='2' and @traits='']">
        <ingredient name="foodRottingFlesh" count="10"/>
    </append>`
	// Tier 3 Upgrade
	c <- `    <append xpath="/recipes/recipe[@tier='3' and @traits='']">
        <ingredient name="foodRottingFlesh" count="20"/>
    </append>`
	// [B] Bonus
	c <- `    <append xpath="/recipes/recipe[contains(@traits, 'B') and not (@traits='BB')]">
        <ingredient name="foodRottingFlesh" count="5"/>
        <ingredient name="medicalBloodBag" count="2"/>
    </append>`
	c <- `    <append xpath="/recipes/recipe[@traits='BB']">
        <ingredient name="foodRottingFlesh" count="7"/>
        <ingredient name="medicalBloodBag" count="3"/>
    </append>`
	// [U] Underground
	c <- `    <append xpath="/recipes/recipe[contains(@traits, 'U')]">
        <ingredient name="plantedMushroom1" count="1"/>
    </append>`
	// [F] Fast
	c <- `    <append xpath="/recipes/recipe[contains(@traits, 'F') and not (@traits='FF')]">
        <ingredient name="drinkCanMegaCrush" count="2"/>
    </append>`
	c <- `    <append xpath="/recipes/recipe[@traits='FF']">
        <ingredient name="drinkCanMegaCrush" count="3"/>
    </append>`
	// [E] Explosive
	c <- `    <append xpath="/recipes/recipe[contains(@traits, 'E') and not (@traits='EE')]">
		<ingredient name="resourceScrapIron" count="4"/>
		<ingredient name="resourceGunPowder" count="4"/>
		<ingredient name="resourceNail" count="1"/>
		<ingredient name="resourceDuctTape" count="1"/>
    </append>`
	c <- `    <append xpath="/recipes/recipe[@traits='EE']">
        <ingredient name="resourceForgedIron" count="1" />
		<ingredient name="resourceGunPowder" count="12"/>
		<ingredient name="resourceNail" count="1"/>
		<ingredient name="resourceDuctTape" count="1"/>
    </append>`
	// [R] Renewable
	c <- `    <append xpath="/recipes/recipe[contains(@traits, 'R')]">
        <ingredient name="drinkJarPureMineralWater" count="10"/>
    </append>`
	// [T] Thorny
	c <- `    <append xpath="/recipes/recipe[contains(@traits, 'T') and not (@traits='TT')]">
        <ingredient name="resourceScrapIron" count="10" />
        <ingredient name="resourceNail" count="10" />
    </append>`
	c <- `    <append xpath="/recipes/recipe[@traits='TT']">
        <ingredient name="resourceScrapIron" count="15" />
        <ingredient name="resourceNail" count="15" />
    </append>`
	// [S] Sweet
	c <- `    <append xpath="/recipes/recipe[contains(@traits, 'S') and not (@traits='SS')]">
        <ingredient name="resourceTestosteroneExtract" count="2"/>
    </append>`
	c <- `    <append xpath="/recipes/recipe[@traits='SS']">
        <ingredient name="resourceTestosteroneExtract" count="3"/>
    </append>`
}
