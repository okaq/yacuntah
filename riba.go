// okaq web server
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "riba.html"
)

// peer id and conn cache

func IdHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("ok"))
}

func RibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("starting riba web server on localhost:8080")
    http.HandleFunc("/", RibaHandler)
    http.HandleFunc("/a", IdHandler)
    http.ListenAndServe(":8080", nil)
}
