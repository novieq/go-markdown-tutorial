package main

import (
    "net/http"
    "os"
    "fmt"
    "github.com/russross/blackfriday"
)

func main() {
    //Heroku gives us a PORT environment variable and expects our web application to bind to it.
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/markdown", GenerateMarkdown)
    http.HandleFunc("/hello", HelloWorld)
    //this is a catch all - so it is defined at the end
    http.Handle("/", http.FileServer(http.Dir("public")))
    /*The last bit of this program starts the server, we pass nil as our handler, which assumes that the HTTP requests
    will be handled by the net/http packages default http.ServeMux, which is configured using http.Handle and http.HandleFunc, respectively.*/
    http.ListenAndServe(":"+port, nil)
}

func HelloWorld(res http.ResponseWriter, req *http.Request) {
    fmt.Fprint(res, "Hello World")
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
    markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
    rw.Write(markdown)
    //In the absence of response code, it is assumed to be 200
}
