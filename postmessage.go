package main

import "fmt"
import "gopkg.in/ini.v1"

func main() {
    cfg, _ := ini.Load([]byte("raw data"),"godbm.ini")
    key := cfg.Section("server").HasKey("PORT")
    fmt.Print(key)
}
