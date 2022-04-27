package gen

type Plant interface {
	IsCompatibleWith(string) bool
	WriteBlockStages(chan string, int, string)
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
