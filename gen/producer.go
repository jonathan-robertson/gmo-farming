package gen

type Producer interface {
	GetPath() string
	GetFilename() string
	Produce(chan string)
}
