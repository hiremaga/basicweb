package main

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/hiremaga/basicweb/books"
	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/markdown", GenerateMarkdown)
	r.HandleFunc("/books", books.Show)
	r.HandleFunc("/books.html", books.ShowHtml)

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":" + port)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
