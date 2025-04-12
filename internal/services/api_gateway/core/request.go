package core

import (
	"io"
	"net/http"
	"time"
)

type RequestOptions struct {
	Method  string
	Path    string
	Headers http.Header
	Body    io.Reader
	Timeout time.Duration // optional timeout override
}
