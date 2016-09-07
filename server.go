
package main 

import "fmt"
import "net/http"
import "io"

func ProcessMessage(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world !")
    fmt.Println("Printer")
}

func main() {
    http.HandleFunc("/", ProcessMessage)
    http.ListenAndServe(":8000", nil)
}
