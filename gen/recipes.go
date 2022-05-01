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
	c <- `<config><append xpath="/recipes">`
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
	c <- `</append></config>`
}

func produceRecipeStub(c chan string, plant Plant, traits ...Trait) {
	switch len(traits) {
	case 0:
		// TODO: tags="learnable"
		c <- fmt.Sprintf(`<recipe name="planted%s1_" count="1" craft_time="%d" traits="">
    <ingredient name="planted%s1" count="1"/>
    <ingredient name="foodRottingFlesh" count="1"/>
    <ingredient name="resourceCloth" count="1"/>
    <ingredient name="resourceYuccaFibers" count="2"/>
</recipe>`,
			plant.GetName(),
			plant.GetCraftTime()*450,
			plant.GetName())
	case 1:
		// TODO: tags="learnable"
		c <- fmt.Sprintf(`<recipe name="planted%s1_%c" count="1" craft_time="%d" traits="%c" craft_area="hotbox">
    <ingredient name="planted%s1_" count="1"/>`,
			plant.GetName(),
			traits[0].code,
			plant.GetCraftTime(),
			traits[0].code,
			plant.GetName())
		produceIngredients(c, traits[0])
		c <- `</recipe>`
	case 2: // support bi-directional recipes
		// TODO: tags="learnable"
		signature := fmt.Sprintf(`<recipe name="planted%s1_%c%c" count="1" craft_time="%d" traits="%c%c" craft_area="hotbox">
    `,
			plant.GetName(),
			traits[0].code, traits[1].code,
			plant.GetCraftTime(),
			traits[0].code, traits[1].code)
		c <- fmt.Sprintf(`%s<ingredient name="planted%s1_%c" count="1"/>`,
			signature,
			plant.GetName(),
			traits[0].code)
		produceIngredients(c, traits[1])
		c <- `</recipe>`
		if traits[0].code != traits[1].code {
			c <- fmt.Sprintf(`%s<ingredient name="planted%s1_%c" count="1"/>`,
				signature,
				plant.GetName(),
				traits[1].code)
			produceIngredients(c, traits[0])
			c <- `</recipe>`
		}
	}
}

func produceIngredients(c chan string, trait Trait) {
	for _, ingredient := range trait.ingredients {
		c <- fmt.Sprintf(`    <ingredient name="%s" count="%d"/>`,
			ingredient.name,
			ingredient.count)
	}
}
