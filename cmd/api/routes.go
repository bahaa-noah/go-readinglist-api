package main

import "net/http"

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthCheck)
	mux.HandleFunc("/v1/books", app.getCreateBookHandler)
	mux.HandleFunc("/v1/books/", app.getUpdateDeleteBooksHandler)
	return mux
}
