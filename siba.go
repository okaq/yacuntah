// okaq web server
package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "strconv"
    "time"
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
    Time float64
    TimeKey string
    Id float64
    IdKey string
    // server comps
    Now int64
    NowKey string
    Lid float64
    LidKey string
    HashKey string
}

func QuidHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    // w.Write([]byte("ok quid!"));
    dec := json.NewDecoder(r.Body)
    defer r.Body.Close()
    // decode into new Quid
    var q Quid
    err := dec.Decode(&q)
    if err != nil {
        fmt.Println(err)
    }
    // fmt.Println(q)
    fmt.Printf("Time: %f, Id: %f.\n", q.Time, q.Id)
    q.TimeKey = strconv.FormatFloat(q.Time, 'f', -1, 64)
    q.IdKey = strconv.FormatFloat(q.Id, 'f', -1, 64)
    fmt.Printf("Time: %s, Id: %s\n", q.TimeKey, q.IdKey)
    // generate server side parts
    q.Now = time.Now().UnixNano()
    q.NowKey = strconv.FormatInt(q.Now, 10)
    // rand default source
    q.Lid = rand.Float64()
    q.LidKey = strconv.FormatFloat(q.Lid, 'f', -1, 64)
    fmt.Println(q)
}

func main() {
    fmt.Println("starting siba web server on localhost:8080")
    http.HandleFunc("/", SibaHandler)
    http.HandleFunc("/a", QuidHandler)
    http.ListenAndServe(":8080", nil)
}
