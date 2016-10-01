// code art okaq.com home
// random font and smaple
// aq@okaq.com
// 2016-09-18
package main

import (
    "fmt"
    "net/http"
    "strconv"
    "sync/atomic"
)

const (
    INDEX = "viba.html"
)

var (
    ID uint64 // uinque id generator
)

func VibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
}

func AnkhHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // atomic int64 counter for user ids
    // server is id:0, admin is id:1, player aq is id:2
    new_id := atomic.AddUint64(&ID, 1)
    w.Header().Set("Content-Type", "application/json")
    s0 := strconv.FormatUint(new_id, 10)
    b0 := []byte(s0)
    w.Write(b0)
}

func main() {
    fmt.Println("starting viba web server on http://localhost:8080");
    ID = 2
    http.HandleFunc("/", VibaHandler)
    http.HandleFunc("/json", JsonHandler)
    http.HandleFunc("/a", AnkhHandler)
    http.ListenAndServe(":8080", nil)
}
