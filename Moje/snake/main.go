package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type Point struct{ X, Y int }
type Direction int

const (
    Up Direction = iota
    Down
    Left
    Right
)

var (
    snake         []Point
    direction     Direction
    apple         Point
    width, height int
)

func main() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()

    width, height = termbox.Size()
    snake = append(snake, Point{width / 2, height / 2})
    apple = Point{width / 2, height / 3}

    go func() {
        for {
            switch ev := termbox.PollEvent(); ev.Type {
            case termbox.EventKey:
                switch ev.Key {
                case termbox.KeyArrowUp:
                    direction = Up
                case termbox.KeyArrowDown:
                    direction = Down
                case termbox.KeyArrowLeft:
                    direction = Left
                case termbox.KeyArrowRight:
                    direction = Right
				case termbox.KeyEsc:
					termbox.Close()
                }
            }
        }
    }()

    for {
        draw()
        update()
        time.Sleep(50 * time.Millisecond)
    }
}

func draw() {
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    for _, p := range snake {
        termbox.SetCell(p.X, p.Y, 'S', termbox.ColorGreen, termbox.ColorDefault)
    }
    termbox.SetCell(apple.X, apple.Y, 'A', termbox.ColorRed, termbox.ColorDefault)
    termbox.Flush()
}

func update() {
    head := snake[0]
    switch direction {
    case Up:
        head.Y--
    case Down:
        head.Y++
    case Left:
        head.X--
    case Right:
        head.X++
    }
    if head.X < 0 || head.Y < 0 || head.X >= width || head.Y >= height {
        termbox.Close()
        panic("Game Over")
    }
    snake = append([]Point{head}, snake[:len(snake)-1]...)
    if head.X == apple.X && head.Y == apple.Y {
        snake = append(snake, Point{})
        apple.X = rand.Intn(width)
        apple.Y = rand.Intn(height)
    }
}