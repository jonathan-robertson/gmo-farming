package gen

// Producer is an interface representing a struct that can produce string content to a channel
type Producer interface {
	GetPath() string
	GetFilename() string
	Produce(chan string)
}
