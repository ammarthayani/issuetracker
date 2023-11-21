package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// TODO: Create API Object to store mux router
	r := mux.NewRouter()
	r.HandleFunc("/", HomeController)

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

func HomeController(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("views/home/index.html")

	tmpl.Execute(w, nil)
}
