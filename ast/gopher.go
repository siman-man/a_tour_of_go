package main

import "fmt"

type Gopher struct {
    gopher string `json:"gopher"`
}

func main() {
    const gopher = "GOPHER"
    gogopher := GOPHER()
    gogopher.gopher = gopher
    fmt.Println(gogopher)
}

func GOPHER() (gopher *Gopher) {
    gopher = &Gopher{
        gopher: "gopher",
    }
    return
}