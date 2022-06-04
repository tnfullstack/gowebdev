package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", diner)
	http.HandleFunc("/diner.jpeg", dinerPic)
	http.ListenAndServe(":8080", nil)
}

func diner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="diner.jpeg">`)
}

func dinerPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("diner.jpeg")
	if err != nil {
		http.Error(w, "file not found", 404)
	}
	defer f.Close()

	io.Copy(w, f)
}
