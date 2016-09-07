
package main 

import (
    "io"
    "net/http"
)

func RootCall(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "API server is running\n")
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
    server := http.Server {
        Addr: ":8000",
        Handler: &myHandler{},
    }
    mux = make(map[string]func(http.ResponseWriter, *http.Request))
    mux["/"] = RootCall
    server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    if h, ok := mux[r.URL.String()]; ok {
        h(w, r)
        return
    }
    io.WriteString(w, "Unknown URL call: "+r.URL.String())
}
