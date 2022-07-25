package main

import (
	"io"
	"net/http"
)

const (
	envLogPath     = "APP_LOGFILE_PATH"
	defaultLogPath = "log.txt"
)

var errInternalServer = []byte("internal error")

type handler struct {
	w io.Writer
}

func newHandler(w io.Writer) *handler {
	return &handler{w: w}
}

func main() {
	// TODO: код писать здесь
}

func (h *handler) logHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: код писать здесь
}
