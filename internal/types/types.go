package types

import (
	"net/http"
)

// Command represents what to execute in the program
type Command struct {
	Comment string
	Method  string
	URL     string
	Headers http.Header
	Body    []byte
}
