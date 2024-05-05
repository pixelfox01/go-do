package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/todo/view", app.todoView)
	mux.HandleFunc("/todo/create", app.todoCreate)
	mux.HandleFunc("/todo/delete", app.todoDelete)

	return mux
}
