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

	title := "Take a shower"
	completed := true

	id, err := app.todos.Insert(title, completed)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/todo/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) todoDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		// http.Error(w, "Not Found", http.StatusNotFound)
		app.notFound(w)
		return
	}

	rows, err := app.todos.Delete(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "Success. Rows Affected: %d\n", rows)
}
