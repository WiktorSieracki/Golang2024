package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
    fmt.Println("Hello, World!")

    err := sdl.Init(sdl.INIT_EVERYTHING)
    if err != nil {
        log.Fatalf("Failed to initialize SDL: %s", err)
    }
    defer sdl.Quit()
}