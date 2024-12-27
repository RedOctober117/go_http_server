package internal

type Method int

const (
	Get Method = iota
	Post
	Head
	Put
	Delete
	Connect
	Options
	Trace
	Patch
)

var methodMap = map[Method]string{
	Get:     "GET",
	Post:    "POST",
	Head:    "HEAD",
	Put:     "PUT",
	Delete:  "DELETE",
	Connect: "CONNECT",
	Options: "OPTIONS",
	Trace:   "TRACE",
	Patch:   "PATCH",
}

func (method Method) String() string {
	return methodMap[method]
}
