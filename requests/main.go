package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
    resp, err := http.Get("http://example.com")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(string(body))
}