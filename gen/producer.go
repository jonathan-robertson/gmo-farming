package gen

type Producer interface {
	GetPath() string
	Produce(chan string)
}
