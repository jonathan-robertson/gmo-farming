package gen

// CrystalHellUIDisplay is responsible for producing content for ui_display.xml
type CrystalHellUIDisplay struct{}

// GetPath returns file path for this producer
func (*CrystalHellUIDisplay) GetPath() string {
	return "Config-CrystalHell"
}

// GetFilename returns filename for this producer
func (*CrystalHellUIDisplay) GetFilename() string {
	return "ui_display.xml"
}

// Produce xml data to the provided channel
func (*CrystalHellUIDisplay) Produce(c chan string) {
	defer close(c)
	c <- `<config><append xpath="/ui_display_info/crafting_category_display">
        <crafting_category_list display_type="hotbox">
		<crafting_category name="Tier1Seeds" icon="ui_game_symbol_crops" display_name="lblCategoryTier1Seeds" />
		<crafting_category name="Tier2Seeds" icon="ui_game_symbol_add" display_name="lblCategoryTier2Seeds" />
		<crafting_category name="Tier3Seeds" icon="ui_game_symbol_healing_factor" display_name="lblCategoryTier3Seeds" />
    </crafting_category_list>
</append></config>`
}
