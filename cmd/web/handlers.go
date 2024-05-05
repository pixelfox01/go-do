package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.Error(w, "Not Found", http.StatusNotFound)
		// http.NotFound(w, r)
		app.notFound(w)
		return
	}
	w.Write([]byte("Hello from GoDo"))
}

func (app *application) todoView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		// http.Error(w, "Not Found", http.StatusNotFound)
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Displaying specific todo with ID %d\n", id)
}

func (app *application) todoCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new todo"))
}
