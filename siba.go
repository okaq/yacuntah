// okaq web server
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "siba.html"
)

func SibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("starting siba web server on localhost:8080")
    http.HandleFunc("/", SibaHandler)
    http.ListenAndServe(":8080", nil)
}
