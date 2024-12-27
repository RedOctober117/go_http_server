package internal

// import "fmt"

type Protocol int

const (
	Version1_0 Protocol = iota
	Version1_1
	Version2_0
)

var protocolMap = map[Protocol]string{
	Version1_0: "HTTP/1.0",
	Version1_1: "HTTP/1.1",
	Version2_0: "HTTP/2.0",
}

func (ver Protocol) String() string {
	return protocolMap[ver]
}
