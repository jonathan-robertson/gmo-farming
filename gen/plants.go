package gen

type Plant interface {
	IsCompatibleWith(string) bool
	WriteStages(chan string, int, string)
}

var Plants []Plant = []Plant{
	CreateMushroom(),
	CreateCorn(),
}
