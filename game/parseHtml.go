package game

import (
	"html/template"
)

type FrontPage struct {
	T *template.Template
}

func ReadHtmlFromFile(filepath string) (*template.Template, error) {
	temp, err := template.ParseFiles(filepath)
	if err != nil {
		return nil, err
	}
	return temp, nil
}
