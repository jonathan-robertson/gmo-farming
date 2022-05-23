package gen

import (
	"data"
	"fmt"
)

// ResearcherRecipes is responsible for producing content for recipes.xml
type ResearcherRecipes struct{}

// GetPath returns file path for this producer
func (*ResearcherRecipes) GetPath() string {
	return "Config-Researcher"
}

// GetFilename returns filename for this producer
func (*ResearcherRecipes) GetFilename() string {
	return "recipes.xml"
}

// Produce xml data to the provided channel
func (p *ResearcherRecipes) Produce(c chan string) {
	defer close(c)
	c <- `<config><append xpath="/recipes">`
	p.produceHotBoxRecipe(c)
	for _, plant := range data.Plants {
		p.produceSchematicsRecipe(c, plant)
		p.producePlantRecipe(c, plant)
		for _, trait1 := range data.Traits {
			if plant.IsCompatibleWith(trait1) {
				p.produceSchematicsRecipe(c, plant, trait1)
				p.producePlantRecipe(c, plant, trait1)
				for _, trait2 := range data.Traits {
					if trait1.IsCompatibleWith(trait2) && plant.IsCompatibleWith(trait2) {
						p.produceSchematicsRecipe(c, plant, trait1, trait2)
						p.producePlantRecipe(c, plant, trait1, trait2)
					}
				}
			}
		}
	}
	c <- `</append></config>`
}

func (*ResearcherRecipes) produceHotBoxRecipe(c chan string) {
	c <- `<recipe name="hotbox" count="1" craft_time="240" craft_area="workbench" tags="learnable,workbenchCrafting">
    <ingredient name="resourceForgedIron" count="50"/>
    <ingredient name="resourceMechanicalParts" count="8"/>
    <ingredient name="resourceWood" count="25"/>
    <ingredient name="resourcePotassiumNitratePowder" count="10"/>
    <ingredient name="resourceClayLump" count="50"/>
</recipe>`
}

func (p *ResearcherRecipes) producePlantRecipe(c chan string, plant data.Plant, traits ...data.Trait) {
	switch len(traits) {
	case 0:
		enhancedSeedCraftTime := plant.GetCraftTime() * 450
		c <- fmt.Sprintf(`<recipe name="planted%s1_" count="1" craft_time="%d" traits="" tags="learnable">
    <ingredient name="planted%s1" count="1"/>
    <ingredient name="foodRottingFlesh" count="1"/>
    <ingredient name="resourceCloth" count="1"/>
    <ingredient name="resourceYuccaFibers" count="2"/>
</recipe>`,
			plant.GetName(),
			enhancedSeedCraftTime,
			plant.GetName())
		c <- fmt.Sprintf(`<recipe name="planted%s1_" count="1" craft_time="%d" traits="" tags="learnable" craft_area="hotbox">
    <ingredient name="planted%s1" count="1"/>
    <ingredient name="foodRottingFlesh" count="1"/>
</recipe>`,
			plant.GetName(),
			enhancedSeedCraftTime,
			plant.GetName())
	case 1:
		c <- fmt.Sprintf(`<recipe name="planted%s1_%c" count="1" craft_time="%d" traits="%c" tags="learnable" craft_area="hotbox">
    <ingredient name="planted%s1_" count="1"/>`,
			plant.GetName(),
			traits[0].Code,
			plant.GetCraftTime(),
			traits[0].Code,
			plant.GetName())
		p.producePlantIngredients(c, traits[0])
		c <- `</recipe>`
	case 2: // support bi-directional recipes
		signature := fmt.Sprintf(`<recipe name="planted%s1_%c%c" count="1" craft_time="%d" traits="%c%c" tags="learnable" craft_area="hotbox">`,
			plant.GetName(),
			traits[0].Code, traits[1].Code,
			plant.GetCraftTime(),
			traits[0].Code, traits[1].Code)
		c <- fmt.Sprintf(`%s
    <ingredient name="planted%s1_%c" count="1"/>`,
			signature,
			plant.GetName(),
			traits[0].Code)
		p.producePlantIngredients(c, traits[1])
		c <- `</recipe>`
		if traits[0].Code == traits[1].Code {
			return
		}
		c <- fmt.Sprintf(`%s
    <ingredient name="planted%s1_%c" count="1"/>`,
			signature,
			plant.GetName(),
			traits[1].Code)
		p.producePlantIngredients(c, traits[0])
		c <- `</recipe>`
	}
}

func (*ResearcherRecipes) producePlantIngredients(c chan string, trait data.Trait) {
	for _, ingredient := range trait.Ingredients {
		c <- fmt.Sprintf(`    <ingredient name="%s" count="%d"/>`,
			ingredient.Name,
			ingredient.Count)
	}
}

func (p *ResearcherRecipes) produceSchematicsRecipe(c chan string, plant data.Plant, traits ...data.Trait) {
	switch len(traits) {
	case 0:
		c <- fmt.Sprintf(`<recipe name="%s" count="1" craft_time="10" traits="" craft_area="hotbox" tags="learnable">
    <ingredient name="resourcePaper" count="10"/>
    <ingredient name="planted%s1" count="100"/>
    <ingredient name="foodRottingFlesh" count="100"/>
</recipe>`,
			plant.GetSchematicName(""),
			plant.GetName())
	case 1:
		c <- fmt.Sprintf(`<recipe name="%s" count="1" craft_time="10" traits="%c" craft_area="hotbox" tags="learnable">
    <ingredient name="resourcePaper" count="10"/>
    <ingredient name="planted%s1_" count="10"/>`,
			plant.GetSchematicName(string(traits[0].Code)),
			traits[0].Code,
			plant.GetName())
		p.produceSchematicIngredients(c, traits[0])
		c <- `</recipe>`
	case 2: // support bi-directional recipes
		signature := fmt.Sprintf(`<recipe name="%s" count="1" craft_time="10" traits="%c%c" craft_area="hotbox" tags="learnable">`,
			plant.GetSchematicName(string(traits[0].Code)+string(traits[1].Code)),
			traits[0].Code, traits[1].Code)
		c <- fmt.Sprintf(`%s
    <ingredient name="resourcePaper" count="10"/>
    <ingredient name="planted%s1_%c" count="10"/>`,
			signature, plant.GetName(), traits[0].Code)
		p.produceSchematicIngredients(c, traits[1])
		c <- `</recipe>`
		if traits[0].Code == traits[1].Code {
			return
		}
		c <- fmt.Sprintf(`%s
    <ingredient name="resourcePaper" count="10"/>
    <ingredient name="planted%s1_%c" count="10"/>`,
			signature, plant.GetName(), traits[1].Code)
		p.produceSchematicIngredients(c, traits[0])
		c <- `</recipe>`
	}
}

func (*ResearcherRecipes) produceSchematicIngredients(c chan string, trait data.Trait) {
	for _, ingredient := range trait.Ingredients {
		c <- fmt.Sprintf(`    <ingredient name="%s" count="%d"/>`,
			ingredient.Name,
			ingredient.Count*10)
	}
}
