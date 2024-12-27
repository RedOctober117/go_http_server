package internal

type Header int

const (
	UserAgent Header = iota
	ContentType
	ContentLength
	Host
	Accept
	AcceptLanguage
	AcceptEncoding
	Referer
	ContentDisposition
)

var headerMap = map[Header]string{
	UserAgent:          "User-Agent",
	ContentType:        "Content-Type",
	ContentLength:      "Content-Length",
	Host:               "Host",
	Accept:             "Accept",
	AcceptLanguage:     "Accept-Language",
	AcceptEncoding:     "Accept-Encoding",
	Referer:            "Referer",
	ContentDisposition: "Content-Disposition",
}

func (header Header) String() string {
	return headerMap[header]
}
