package game

import (
	"fmt"
	"html/template"
	"log"
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
	path := r.URL.Path[1:]
	fmt.Println("Path clicked: ", path)
	if path == "" {
		path = "intro"
	}
	if chap, foundMap := h.s[path]; foundMap {
		err := h.temp.Execute(w, chap)
		if err != nil {
			log.Println("Error: ", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Didn't find the chapter.", http.StatusInternalServerError)
	}
}

func RouterHandler(s Story, t *template.Template) (http.Handler, error) {
	return handler{
		s,
		t,
	}, nil
}
