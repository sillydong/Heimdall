package web

import "github.com/sillydong/heimdall/server/model"

type Web struct {
	m *model.Model
}

func NewWeb(m *model.Model) *Web {
	return &Web{
		m: m,
	}
}

func (w *Web) Run() {
}
