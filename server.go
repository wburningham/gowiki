package main

import (
    "fmt"
    "net/http"
)

func handler(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Hi there, I love %s!", req.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}