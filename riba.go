// okaq web server
package main

import (
    "fmt"
    "io/ioutil"
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
    // post format for joruri waka roboto font data
    // arraybuffer, uint8array from threshold bit array
    // slice and join 8 "0" and "1" chars to form byte string
    // parseInt(byte, 2) and push to buffer array
    // json.stringify uint8array
    // send to server as type arraybuffer
    // read from req body as new bytes buffer
    // json marshal bytes
    // re-encode to json file, or base64 string
    // simplify, keep as bit string
    b0, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(b0))
    // bit array string, conv to golang byte slice
    defer r.Body.Close()
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
