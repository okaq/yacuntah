// okaq web server
package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "strconv"
    "time"
)

const (
    INDEX = "riba.html"
)

var (
    // rand seed and rng
    // peer id cache
    R *rand.Rand
    P map[string]string
)

// peer id and conn cache

func IdHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("ok"))
    r0 := R.Int()
    s0 := strconv.Itoa(r0)
    b0 := []byte(s0)
    w.Write(b0)
}

func PoolHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func RibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func main() {
    fmt.Println("starting riba web server on localhost:8080")
    R = rand.New(rand.NewSource(time.Now().UnixNano()))
    http.HandleFunc("/", RibaHandler)
    http.HandleFunc("/a", IdHandler)
    http.HandleFunc("/b", PoolHandler)
    http.ListenAndServe(":8080", nil)
}
