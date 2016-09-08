// spelling salts
// ai nano game
// aq@okaq.com
// 2016-09-12
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "wiba.html"
)

func WibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func main() {
    fmt.Println("starting wiba web server on http://localhost:8080");
    http.HandleFunc("/", WibaHandler)
    http.HandleFunc("/json", JsonHandler)
    http.ListenAndServe(":8080", nil)
}
