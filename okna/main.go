package main

import "github.com/veandco/go-sdl2/sdl"

type Layout string

const (
	Static  Layout = "static"
	Dynamic Layout = "dynamic"
)

type Rect struct {
	Width, Height int32
	TypeRect      Layout
	ListRect      []Rect
}

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	r := Rect{
		Width:    500,
		Height:   500,
		TypeRect: Static,
	}
}