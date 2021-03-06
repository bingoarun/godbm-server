package main
import (
    "encoding/json"
    "os"
    "fmt"
)
type server struct {
    Port string
    Log_dir string
    Storage string
}
type search struct {
    Algorithm string
    Concurrency string
}
type Configuration struct {
    Test string
    Server server
    Search search
}

func main() {
    file, err := os.Open("config.json")
    decoder := json.NewDecoder(file)

    var Confobj Configuration
    if err = decoder.Decode(&Confobj); err != nil{
        fmt.Print(err)
    }
    fmt.Printf("%+v\n", Confobj.Server.Port)
    fmt.Printf("%+v\n",Confobj)
}
