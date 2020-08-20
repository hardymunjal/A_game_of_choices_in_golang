package game

import (
	"html/template"
	"net/http"
)

type handler struct {
	s    Story
	temp *template.Template
}

func (h handler) getTemplate(t *template.Template) {
	h.temp = t
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.temp.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}

func RouterHandler(s Story, t *template.Template) (http.Handler, error) {
	return handler{
		s,
		t,
	}, nil
}
