package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Hello from GoDo"))
}

func todoView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Displaying specific todo with ID %d\n", id)
}

func todoCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating todo item...\n"))
}
