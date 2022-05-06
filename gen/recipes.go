package gen

import (
	"data"
	"fmt"
)

func WritePlantRecipes(target string) error {
	file, err := getFile(fmt.Sprintf("Config-%s/recipes.xml", target))
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go producePlantRecipes(c, target)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func producePlantRecipes(c chan string, target string) {
	optionalTags := ""
	switch target {
	case "Vanilla":
		optionalTags = ` tags="learnable"`
	}

	defer close(c)
	c <- `<config><append xpath="/recipes">`
	produceHotBoxRecipe(c)
	for _, plant := range data.Plants {
		producePlantRecipe(c, plant, optionalTags)
		for i1 := 0; i1 < len(data.Traits); i1++ {
			if plant.IsCompatibleWith(data.Traits[i1]) {
				producePlantRecipe(c, plant, optionalTags, data.Traits[i1])
			}
			for i2 := i1; i2 < len(data.Traits); i2++ {
				if data.Traits[i1].IsCompatibleWith(data.Traits[i2]) {
					if plant.IsCompatibleWith(data.Traits[i1]) && plant.IsCompatibleWith(data.Traits[i2]) {
						producePlantRecipe(c, plant, optionalTags, data.Traits[i1], data.Traits[i2])
					}
				}
			}
		}
	}
	c <- `</append></config>`
}

func produceHotBoxRecipe(c chan string) {
	c <- `<recipe name="hotbox" count="1" craft_area="workbench" tags="learnable,workbenchCrafting">
	<ingredient name="resourceForgedIron" count="50"/>
	<ingredient name="resourceMechanicalParts" count="8"/>
	<ingredient name="resourceWood" count="25"/>
</recipe>`
}

func producePlantRecipe(c chan string, plant data.Plant, optionalTags string, traits ...data.Trait) {
	switch len(traits) {
	case 0:
		c <- fmt.Sprintf(`<recipe name="planted%s1_" count="1" craft_time="%d" traits=""%s>
    <ingredient name="planted%s1" count="1"/>
    <ingredient name="foodRottingFlesh" count="1"/>
    <ingredient name="resourceCloth" count="1"/>
    <ingredient name="resourceYuccaFibers" count="2"/>
</recipe>`,
			plant.GetName(),
			plant.GetCraftTime()*450,
			optionalTags,
			plant.GetName())
	case 1:
		c <- fmt.Sprintf(`<recipe name="planted%s1_%c" count="1" craft_time="%d" traits="%c" craft_area="hotbox"%s>
    <ingredient name="planted%s1_" count="1"/>`,
			plant.GetName(),
			traits[0].Code,
			plant.GetCraftTime(),
			traits[0].Code,
			optionalTags,
			plant.GetName())
		produceIngredients(c, traits[0])
		c <- `</recipe>`
	case 2: // support bi-directional recipes
		signature := fmt.Sprintf(`<recipe name="planted%s1_%c%c" count="1" craft_time="%d" traits="%c%c" craft_area="hotbox"%s>
    `,
			plant.GetName(),
			traits[0].Code, traits[1].Code,
			plant.GetCraftTime(),
			traits[0].Code, traits[1].Code,
			optionalTags)
		c <- fmt.Sprintf(`%s<ingredient name="planted%s1_%c" count="1"/>`,
			signature,
			plant.GetName(),
			traits[0].Code)
		produceIngredients(c, traits[1])
		c <- `</recipe>`
		if traits[0].Code != traits[1].Code {
			c <- fmt.Sprintf(`%s<ingredient name="planted%s1_%c" count="1"/>`,
				signature,
				plant.GetName(),
				traits[1].Code)
			produceIngredients(c, traits[0])
			c <- `</recipe>`
		}
	}
}

func produceIngredients(c chan string, trait data.Trait) {
	for _, ingredient := range trait.Ingredients {
		c <- fmt.Sprintf(`    <ingredient name="%s" count="%d"/>`,
			ingredient.Name,
			ingredient.Count)
	}
}
