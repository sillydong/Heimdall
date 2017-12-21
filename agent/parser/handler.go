package parser

import "net/http"

type handler struct {
}

func (h *handler) Handle(req *http.Request) (string, int32) {
	return "", 0
}
