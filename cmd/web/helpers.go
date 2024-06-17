package main

import (
	"html/template"
	"net/http"
)

func (app *application) loadTemplate(w http.ResponseWriter, ts *template.Template, data interface{}) error {
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		return err
	}

	return nil
}
