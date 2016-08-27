// LD36 - Ancient Tech
// okaq.com entry
// Puczicalchunkuna (lit. “Temple of the Tree of Minds”)
// okaq.ld36@gmail.com
// 2016-08-28

// Dzonocaob - Bitmap Sampler
package main

import (
    "fmt"
    "net/http"
)

const (
    INDEX = "dzono.html"
)

func DzonoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func LoadHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    fmt.Println("Sample to load: ")
    s0 := ""
    fmt.Scanln(&s0)
    fmt.Println(s0)
}

func main() {
    fmt.Println("starting dzono web server on localhost:8080")
    http.HandleFunc("/", DzonoHandler)
    http.HandleFunc("/load", LoadHandler)
    http.ListenAndServe(":8080", nil)
}
