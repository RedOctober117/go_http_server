package internal

type StatusCode int

const (
	Code200 StatusCode = iota
	Code201
	Code400
	Code404
	Code500
	Code501
)

var statusCodeMap = map[StatusCode]string{
	Code200: "200 OK",
	Code201: "201 OK",
	Code400: "400 OK",
	Code404: "404 OK",
	Code500: "500 OK",
	Code501: "501 OK",
}

func (status StatusCode) String() string {
	return statusCodeMap[status]
}
