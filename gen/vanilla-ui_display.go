package gen

// VanillaUIDisplay is responsible for producing content for ui_display.xml
type VanillaUIDisplay struct{}

// GetPath returns file path for this producer
func (*VanillaUIDisplay) GetPath() string {
	return "Config-Vanilla"
}

// GetFilename returns filename for this producer
func (*VanillaUIDisplay) GetFilename() string {
	return "ui_display.xml"
}

// Produce xml data to the provided channel
func (*VanillaUIDisplay) Produce(c chan string) {
	defer close(c)
	c <- `<config><append xpath="/ui_display_info/crafting_category_display">
        <crafting_category_list display_type="hotbox">
        <crafting_category name="Tier1SeedResearch" icon="ui_game_symbol_book" display_name="lblCategoryTier1SeedResearch" />
        <crafting_category name="Tier1Seeds" icon="ui_game_symbol_crops" display_name="lblCategoryTier1Seeds" />
        <crafting_category name="Tier2SeedResearch" icon="ui_game_symbol_book" display_name="lblCategoryTier2SeedResearch" />
        <crafting_category name="Tier2Seeds" icon="ui_game_symbol_add" display_name="lblCategoryTier2Seeds" />
        <crafting_category name="Tier3SeedResearch" icon="ui_game_symbol_book" display_name="lblCategoryTier3SeedResearch" />
        <crafting_category name="Tier3Seeds" icon="ui_game_symbol_healing_factor" display_name="lblCategoryTier3Seeds" />
    </crafting_category_list>
</append></config>`
}
