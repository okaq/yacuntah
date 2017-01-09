// okaq web server
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "strconv"
    "strings"
    "time"
)

const (
    INDEX = "riba.html"
    IMP = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890~!@#$%^&*()-_=+[{]}|;:'",<.>/?`
    SAVE = "save/noto.json"
)

var (
    // rand seed and rng
    // peer id cache
    R *rand.Rand
    P map[string]string
    I []string
    C int
    B bytes.Buffer
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
    fmt.Println(b0)
    // raw binary bytes recieved
    // marshal json into base64 string and save to file
    // bit array string, conv to golang byte slice
    defer r.Body.Close()
    j0, err := json.Marshal(b0)
    if err != nil {
        fmt.Println(err)
    }
    s1 := string(j0)
    fmt.Println(len(s1), s1)
    s2 := I[C]
    fmt.Println(C, s2)
    s3 :=fmt.Sprintf("\"%s\":%s\n", s2, s1)
    fmt.Println(s3)
    B.WriteString(s3)
    // timed ticker so will not race
    C = C + 1
    if C >= (len(I) - 1) {
        fmt.Println("dune")
        fmt.Println(B.String())
        ioutil.WriteFile(SAVE, B.Bytes(), 0777)
        fmt.Printf("File written: %s\n", SAVE)
    }
    s0 := fmt.Sprintf("%d bytes read ok!", len(b0))
    w.Write([]byte(s0))
}

func RibaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    http.ServeFile(w, r, INDEX)
}

func Import() {
    // set up rune array and counter
    I = strings.Split(IMP, "")
    fmt.Println(I, len(I))
    C = 0 // atomic increment
    // var B bytes.Buffer
}

func main() {
    fmt.Println("starting riba web server on localhost:8080")
    R = rand.New(rand.NewSource(time.Now().UnixNano()))
    Import()
    http.HandleFunc("/", RibaHandler)
    http.HandleFunc("/a", IdHandler)
    http.HandleFunc("/b", PoolHandler)
    http.ListenAndServe(":8080", nil)
}
