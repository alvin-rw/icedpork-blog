package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alvin-rw/icedpork-blog/cmd/internal/data"
	"github.com/go-playground/form/v4"
	"github.com/julienschmidt/httprouter"
)

const defaultPageSize int = 5

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.loadTemplate(w, app.templateCache["home.html"], nil)
}

func (app *application) blog(w http.ResponseWriter, r *http.Request) {
	posts, err := app.models.PostModel.GetPosts(defaultPageSize)
	if err != nil {
		app.internalServerErrorResponse(w, err)
	}

	data := templateData{
		Posts: posts,
	}

	app.loadTemplate(w, app.templateCache["blog.html"], data)
}

func (app *application) createPostPage(w http.ResponseWriter, r *http.Request) {
	app.loadTemplate(w, app.templateCache["create_post.html"], nil)
}

type createPostForm struct {
	Title     string `form:"title"`
	Content   string `form:"content"`
	Published bool   `form:"published"`
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.internalServerErrorResponse(w, err)
		return
	}

	var input createPostForm
	formdecoder := form.NewDecoder()

	err = formdecoder.Decode(&input, r.Form)
	if err != nil {
		app.internalServerErrorResponse(w, err)
	}

	post := &data.Post{
		Title:     input.Title,
		Content:   input.Content,
		Published: input.Published,
	}

	err = app.models.PostModel.Create(post)
	if err != nil {
		app.internalServerErrorResponse(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/blog/post/%d", post.ID), http.StatusSeeOther)
}

func (app *application) showPost(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.notFoundResponse(w)
		return
	}

	post, err := app.models.PostModel.GetPost(id)
	if err != nil {
		app.internalServerErrorResponse(w, err)
		return
	}

	data := templateData{
		Posts: []data.Post{
			*post,
		},
	}

	app.loadTemplate(w, app.templateCache["show_post.html"], data)
}
