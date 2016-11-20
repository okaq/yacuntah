// okaq web server
package main

import (
    "encoding/json"
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

type Quid struct {
    Time float32
    Id float32
}

func QuidHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("ok quid!"));
    dec := json.NewDecoder(r.Body)
    var q Quid
    err := dec.Decode(&q)
    if err != nil {
        fmt.Println(err)
    }
    // fmt.Println(q)
    fmt.Printf("Time: %f, Id: %f.\n", q.Time, q.Id)
    defer r.Body.Close()
}

func main() {
    fmt.Println("starting siba web server on localhost:8080")
    http.HandleFunc("/", SibaHandler)
    http.HandleFunc("/a", QuidHandler)
    http.ListenAndServe(":8080", nil)
}
