// okaq web server
package main

import (
    "crypto/md5"
    "encoding/ascii85"
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

type Peer struct {
    Q Quid
}

type Cache struct {
    sync.RWMutex
    Peers map[string]*Peer
}

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
    Hash string
    HashSum [md5.Size]byte
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
    q.Hash = fmt.Sprintf("%s:%s:%s:%s",q.TimeKey,q.IdKey,q.NowKey,q.LidKey)
    q.HashSum = md5.Sum([]byte(q.Hash))
    q.HashKey = string(q.HashSum[:])
    fmt.Println(q)
    // ascii representation
    b0 := make([]byte, ascii85.MaxEncodedLen(len(q.HashSum[:])))
    n0 := ascii85.Encode(b0, q.HashSum[:])
    fmt.Println(b0, n0)
    s0 := string(b0)
    fmt.Println(s0)
    w.Header().Set("Content-Type", "text/plain")
    // w.Write([]byte(s0))
    w.Write(b0)
    // pipeline calls for each stage
    // fan out work to available workers
    // cache id keys in map[string]Peer
    // golang maps not concurrency safe!
}

func main() {
    fmt.Println("starting siba web server on localhost:8080")
    http.HandleFunc("/", SibaHandler)
    http.HandleFunc("/a", QuidHandler)
    http.ListenAndServe(":8080", nil)
}
