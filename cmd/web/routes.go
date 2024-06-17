package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) router() http.Handler {
	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("./view/static"))

	router.HandlerFunc(http.MethodGet, "/", app.home)

	router.HandlerFunc(http.MethodGet, "/blog", app.blog)
	router.HandlerFunc(http.MethodGet, "/blog/create", app.createPostPage)
	router.HandlerFunc(http.MethodPost, "/blog/create", app.createPost)
	router.HandlerFunc(http.MethodGet, "/blog/post/:id", app.showPost)

	return router
}
