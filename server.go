
package main 

import (
    "io"
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "fmt"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "API server is running\n")
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
    key := mux.Vars(r)["key"]
    io.WriteString(w, fmt.Sprintf("Need to process key: %s", key))
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
    key := mux.Vars(r)["key"]
    value := mux.Vars(r)["value"]
    io.WriteString(w, fmt.Sprintf("Need to process \n key: %s \n value: %s\n", key, value))
}


func main() {
    router := mux.NewRouter()
    router.HandleFunc("/",HomeHandler)
    router.HandleFunc("/get/{key}",GetHandler)
    router.HandleFunc("/post/{key}/{value}",PostHandler)
    log.Fatal(http.ListenAndServe(":8000", router))
}
