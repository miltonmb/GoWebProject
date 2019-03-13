package main

import (
	"html/template"
	"net/http"
)

type newPage struct {
	Title string
}

type Libro struct {
	ISBN    string
	name    string
	inStock int
	editora string
}

func makeLibro(ISBN string, name string, inStock int, editora string) Libro {
	return Libro{ISBN, name, inStock, editora}
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	p := newPage{Title: "Proyecto 2"}
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, p)
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":8000", nil)
}
