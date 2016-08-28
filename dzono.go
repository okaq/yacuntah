// LD36 - Ancient Tech
// okaq.com entry
// Puczicalchunkuna (lit. “Temple of the Tree of Minds”)
// okaq.ld36@gmail.com
// 2016-08-28

// Dzonocaob - Bitmap Sampler
package main

import (
    "encoding/json"
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
    fmt.Printf("Sending %s.\n", s0)
    s1 := make([]string, 1)
    s1[0] = s0
    j0, err := json.Marshal(s1)
    if err != nil {
        fmt.Println(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(j0)
}

func SampHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    fmt.Println(r.URL)
    http.ServeFile(w, r, r.URL.Path[1:])
    // http.FileServer behavior change in go1.7?
}

func main() {
    fmt.Println("starting dzono web server on localhost:8080")
    // http.Handle("/samp/", http.StripPrefix("/samp", http.FileServer(http.Dir("samp/"))))
    http.HandleFunc("/", DzonoHandler)
    http.HandleFunc("/load", LoadHandler)
    // fs0 := http.FileServer(http.Dir("samp"))
    // http.Handle("/samp/", http.StripPrefix("/samp", http.FileServer(http.Dir("samp/"))))
    http.HandleFunc("/samp/", SampHandler)
    http.ListenAndServe(":8080", nil)
}
