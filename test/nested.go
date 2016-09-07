package main

import (
    "encoding/json"
    "fmt"
    "log"
)

func main() {
    jStr := `
    {
        "AAA": {
            "assdfdff": ["asdf"],
            "fdsfa": ["1231", "123"]
        }
    }
    `

    type Inner struct {
        Key2 []string `json:"assdfdff"`
        Key3 []string `json:"fdsfa"`
    }
    type Container struct {
        Key Inner `json:"AAA"`
    }
    var cont Container
    if err := json.Unmarshal([]byte(jStr), &cont); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", cont)
}
