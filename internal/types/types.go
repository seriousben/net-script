package types

import (
	"net/http"
)

type Command struct {
	Comment string
	Method  string
	URL     string
	Headers http.Header
	Body    []byte
}
