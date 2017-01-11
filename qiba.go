// okaq.com web
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "qiba.html"
)

func QibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("starting server on http:localhost:8080")
    http.HandleFunc("/", QibaHandler)
    http.ListenAndServe(":8080", nil)
}
