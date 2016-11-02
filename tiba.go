// okaq web server
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "tiba.html"
)

func TibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("ready one"))
}

func main() {
    fmt.Println("starting tiba web server on http://localhost:8080")
    http.HandleFunc("/", TibaHandler)
    http.HandleFunc("/u", UserHandler)
    http.ListenAndServe(":8080", nil)
}
