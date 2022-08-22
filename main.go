package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Color struct {
	R, G, B uint8
}

type Engine struct {
	Title    string
	Width    int32
	Height   int32
	Scale    int
	FPS      int
	window   *sdl.Window
	renderer *sdl.Renderer
	buffer   *sdl.Surface
}

type UpdateFn func(*Engine)
type RenderFn func(*Engine)

func NewEngine(title string, width int32, height int32, scale int, fps int) (*Engine, error) {
	engine := Engine{
		Title:  title,
		Width:  width,
		Height: height,
		Scale:  scale,
		FPS:    fps,
	}

	window, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		width*int32(scale),
		height*int32(scale),
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return &engine, err
	}
	engine.window = window

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return &engine, err
	}
	engine.renderer = renderer
	renderer.SetScale(float32(scale), float32(scale))

	surface, err := sdl.CreateRGBSurface(0, int32(width), int32(height), 32, 0xff000000, 0x00ff0000, 0x0000ff00, 0x000000ff)
	if err != nil {
		return &engine, err
	}
	engine.buffer = surface

	return &engine, nil
}

func (e *Engine) Destroy() {
	if e.renderer != nil {
		e.renderer.Destroy()
		e.renderer = nil
	}
	if e.window != nil {
		e.window.Destroy()
		e.window = nil
	}
	e.buffer = nil
}

func (e *Engine) Draw(x int, y int, r uint8, g uint8, b uint8) {
	e.buffer.Set(x, y, sdl.RGBA8888{R: r, G: g, B: b, A: 255})
}

func (e *Engine) Loop(update UpdateFn, render RenderFn) {
	running := true
	skips := 0
	interval := 1000 / e.FPS

	for running {
		before := time.Now().UnixMilli()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		texture, _ := e.renderer.CreateTextureFromSurface(e.buffer)

		e.renderer.SetDrawColor(0.0, 0.0, 0.0, 1.0)
		e.renderer.Clear()

		update(e)
		if skips > 0 {
			skips--
			sdl.Delay(uint32(interval))
			fmt.Println("frame skip")
			continue
		} else {
			render(e)
			e.renderer.Copy(
				texture,
				&sdl.Rect{
					X: 0,
					Y: 0,
					W: e.Width,
					H: e.Height,
				},
				&sdl.Rect{
					X: 0,
					Y: 0,
					W: e.Width,
					H: e.Height,
				})

			e.renderer.Present()
		}

		after := time.Now().UnixMilli()
		delay := int64(interval) - (after - before)
		for delay < 0 {
			skips++
			delay += int64(interval)
		}

		sdl.Delay(uint32(delay))
	}
}

func main() {

	engine, err := NewEngine("Pixels", 300, 200, 4, 60)
	if err != nil {
		panic(err)
	}
	defer engine.Destroy()

	engine.Loop(Update, Render)
}

func Update(e *Engine) {}

func Render(e *Engine) {
	for i := 0; i < int(e.Width)*int(e.Height); i++ {
		nx := rand.Intn(int(e.Width))
		ny := rand.Intn(int(e.Height))
		red := uint8(rand.Intn(255))
		green := uint8(rand.Intn(255))
		blue := uint8(rand.Intn(255))

		e.Draw(nx, ny, red, green, blue)
	}
}
