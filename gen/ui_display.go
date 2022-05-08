package gen

import (
	"fmt"
)

func WriteUiDisplay(target string) error {

	file, err := getFile(fmt.Sprintf("Config-%s/ui_display.xml", target))
	if err != nil {
		return err
	}
	defer file.Close()
	c := make(chan string, 10)
	go produceUiDisplay(c, target)
	for line := range c {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func produceUiDisplay(c chan string, target string) {
	defer close(c)
	c <- `<config><append xpath="/ui_display_info/crafting_category_display">
    <crafting_category_list display_type="hotbox">`

	spacing := "        "
	if target == "Vanilla" {
		c <- spacing + `<crafting_category name="SeedExperiments" icon="ui_game_symbol_infection" display_name="lblCategorySeedExperiments" />`
	}

	c <- spacing + `<crafting_category name="SeedEnhancements" icon="ui_game_symbol_crops" display_name="lblCategorySeedEnhancements" />`

	if target == "Vanilla" {
		c <- spacing + `<crafting_category name="SeedTraitResearch" icon="ui_game_symbol_book" display_name="lblCategorySeedTraitResearch" />`
	}

	c <- spacing + `<crafting_category name="SeedTraits" icon="ui_game_symbol_add" display_name="lblCategorySeedTraits" />
    </crafting_category_list>
</append></config>`
}
