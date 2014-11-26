package main

import (
	"net/http"
	"os"

	"github.com/GeertJohan/go.rice"
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

	public := rice.MustFindBox("public").HTTPBox()

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(public),
	)
	n.UseHandler(r)
	n.Run(":" + port)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
