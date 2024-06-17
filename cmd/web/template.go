package main

import (
	"html/template"
	"path/filepath"

	"github.com/alvin-rw/icedpork-blog/cmd/internal/data"
)

type templateData struct {
	Posts []data.Post
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./view/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles("./view/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./view/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
