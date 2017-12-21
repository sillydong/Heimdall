package mirror

import (
	"net/http"

	"github.com/sillydong/heimdall/agent/parser"
	"github.com/sirupsen/logrus"
)

//mirror listen on the port to accept requests from nginx
type Mirror struct {
	option *MirrorOption
	p      *parser.Parser
}

//mirror option
type MirrorOption struct {
	Listen string `yaml:"listen"`
}

//create mirror instance
func NewMirror(option *MirrorOption, p *parser.Parser) *Mirror {
	return &Mirror{
		option: option,
		p:      p,
	}
}

//run mirror in gorutine
func (m *Mirror) Run() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			logrus.Infof("request: %v/%v", r.Host, r.RequestURI)
			m.p.Parse(r)
			w.Write([]byte{})
		})
		logrus.Fatal(http.ListenAndServe(m.option.Listen, nil))
	}()
}
