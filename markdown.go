package main

import (
    "net/http"

    "github.com/russross/blackfriday"
)

func main() {
    http.HandleFunc("/markdown", GenerateMarkdown)
    //this is a catch all - so it is defined at the end
    http.Handle("/", http.FileServer(http.Dir("public")))
    /*The last bit of this program starts the server, we pass nil as our handler, which assumes that the HTTP requests
    will be handled by the net/http packages default http.ServeMux, which is configured using http.Handle and http.HandleFunc, respectively.*/
    http.ListenAndServe(":8080", nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
    markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
    rw.Write(markdown)
    //In the absence of response code, it is assumed to be 200
}
