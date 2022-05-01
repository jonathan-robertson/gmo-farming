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
		produceRecipeStub(c, plant)
		for i1 := 0; i1 < len(Traits); i1++ {
			if plant.IsCompatibleWith(Traits[i1].code) {
				produceRecipeStub(c, plant, Traits[i1])
			}
			for i2 := i1; i2 < len(Traits); i2++ {
				if Traits[i1].isCompatibleWith(Traits[i2]) {
					if plant.IsCompatibleWith(Traits[i1].code) && plant.IsCompatibleWith(Traits[i2].code) {
						produceRecipeStub(c, plant, Traits[i1], Traits[i2])
					}
				}
			}
		}
	}
	c <- `</append>`
	produceRecipeModifications(c)
	c <- `</config>`
}

func produceRecipeStub(c chan string, plant Plant, traits ...Trait) {
	switch len(traits) {
	case 0:
		// TODO: tags="learnable"
		c <- fmt.Sprintf(`<recipe name="planted%s1_" count="1" craft_time="%d" traits="">
    <ingredient name="planted%s1" count="1"/>
</recipe>`,
			plant.GetName(),
			plant.GetCraftTime()*450,
			plant.GetName())
	case 1:
		// TODO: tags="learnable"
		c <- fmt.Sprintf(`<recipe name="planted%s1_%c" count="1" craft_time="%d" traits="%c" craft_area="hotbox">
    <ingredient name="planted%s1_" count="1"/>
</recipe>`,
			plant.GetName(),
			traits[0].code,
			plant.GetCraftTime(),
			traits[0].code,
			plant.GetName())
	case 2: // support bi-directional recipes
		// TODO: tags="learnable"
		signature := fmt.Sprintf(`<recipe name="planted%s1_%c%c" count="1" craft_time="%d" traits="%c%c" craft_area="hotbox">
    <ingredient name="planted%s1_`,
			plant.GetName(),
			traits[0].code, traits[1].code,
			plant.GetCraftTime(),
			traits[0].code, traits[1].code,
			plant.GetName())
		c <- fmt.Sprintf(`%s%c" count="1"/>
</recipe>`,
			signature,
			traits[0].code)
		c <- fmt.Sprintf(`%s%c" count="1"/>
</recipe>`,
			signature,
			traits[1].code)
	}
}

func produceRecipeModifications(c chan string) {
	// Initial Enhancement
	c <- `    <append xpath="/recipes/recipe[@traits='']">
        <ingredient name="resourceCloth" count="1"/>
        <ingredient name="resourceYuccaFibers" count="2"/>
        <ingredient name="foodRottingFlesh" count="1"/>
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
