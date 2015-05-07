package main

import (
    "fmt"
    "net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Hi there, I love %s!", req.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}