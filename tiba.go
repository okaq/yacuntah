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
    INDEX = "tiba.html"
)

var (
    Rng *RNG
)

type RNG struct {
    seed int64
    source rand.Source
    R *rand.Rand
}

func NewRNG() *RNG {
    r0 := RNG{}
    r0.seed = time.Now().UnixNano()
    r0.source = rand.NewSource(r0.seed)
    r0.R = rand.New(r0.source)
    return &r0
}

func (r *RNG) Next() float32 {
    return r.R.Float32()
}

func TibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    w.Write([]byte("ready one"))
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // write a single random float32
    f0 := Rng.Next()
    f1 := float64(f0)
    s0 := strconv.FormatFloat(f1, 'g', -1, 32)
    b0 := []byte(s0)
    w.Write(b0)
    // w.Write([]byte(string(Rng.Next())))
}

func main() {
    fmt.Println("starting tiba web server on http://localhost:8080")
    Rng = NewRNG()
    fmt.Println(Rng.Next())
    http.HandleFunc("/", TibaHandler)
    http.HandleFunc("/u", UserHandler)
    http.HandleFunc("/d", DataHandler)
    http.ListenAndServe(":8080", nil)
}
