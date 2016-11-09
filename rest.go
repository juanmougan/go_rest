package main
 
import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Message struct {
    Text string	`json:"text"`
}
 
func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

func about (w http.ResponseWriter, r *http.Request) {
 
    m := Message{"Sample API written by Juan Manuel Mougan"}
    b, err := json.Marshal(m)
 
    if err != nil {
        panic(err)
    }
 
     w.Write(b)
}
 
func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/about/", about)
    http.ListenAndServe(":8080", nil)
}
