package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/markdown", GenerateMarkdown)
	r.PathPrefix("/").Handler(http.FileServer((http.Dir("public"))))

	http.ListenAndServe(":"+port, r)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
