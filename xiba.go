// modern aliens
// ai nano game
// aq@okaq.com
// 2016-08-30
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "xiba.html"
)

func XibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("starting xiba web server on http://localhost:8080");
    http.HandleFunc("/", XibaHandler)
    http.ListenAndServe(":8080", nil)
}
