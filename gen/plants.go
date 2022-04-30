package gen

type Plant interface {
	IsCompatibleWith(rune) bool
	WriteBlockStages(chan string, string)
	GetName() string
	GetNamePlural() string
	GetDisplayName() string
	GetDescription() string
	GetPreferredConsumer() string
	GetCraftTime() int
}

var Plants []Plant = []Plant{
	CreateMushroom(),
	CreateCorn(),
}
