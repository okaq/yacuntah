// okaq web server
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "siba.html"
)

// cache list of peer quids
// maintain broker for any peer net

func SibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func QuidHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("ok quid!"));
}

func main() {
    fmt.Println("starting siba web server on localhost:8080")
    http.HandleFunc("/", SibaHandler)
    http.HandleFunc("/a", QuidHandler)
    http.ListenAndServe(":8080", nil)
}
