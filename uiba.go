// okaq web server
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "uiba.html"
)

func UibaHandler() {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("starting viba web server on http://localhost:8080");
    http.HandleFunc("/", VibaHandler)
    http.ListenAndServe(":8080", nil)
}
