package main

import (
	"net/http"
)

func (app *application) internalServerErrorResponse(w http.ResponseWriter, err error) {
	app.logger.Println(err)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) notFoundResponse(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
