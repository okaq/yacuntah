// code art okaq.com home
// random font and smaple
// aq@okaq.com
// 2016-09-18
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "viba.html"
)

func VibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func main() {
    fmt.Println("starting viba web server on http://localhost:8080");
    http.HandleFunc("/", VibaHandler)
    http.HandleFunc("/json", JsonHandler)
    http.ListenAndServe(":8080", nil)
}
