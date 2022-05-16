package gen

type VanillaUiDisplay struct{}

func (*VanillaUiDisplay) GetPath() string {
	return "Config-Vanilla/ui_display.xml"
}

func (*VanillaUiDisplay) Produce(c chan string) {
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
