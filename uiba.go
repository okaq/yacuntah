// okaq web server
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "uiba.html"
)

func UibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("starting uiba web server on http://localhost:8080");
    http.HandleFunc("/", UibaHandler)
    http.ListenAndServe(":8080", nil)
}
