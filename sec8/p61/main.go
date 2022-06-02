package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

type Person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", handleHome)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {

	// body
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)

	err := tmpl.ExecuteTemplate(w, "index.html", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
	}
}
