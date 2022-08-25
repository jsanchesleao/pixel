# Pixel Engine

This is a simple engine that displays a screen, allows for individual pixel changing and keeps a constant frame rate.

I created this module as a sandbox for some experimentation on retro gaming.
It is not the most optimized way of rendering graphics, but I found it was easy to port into different languages and environments.

## Installation

```
go get github.com/jsanchesleao/pixel@v1.0.1
```

## Usage

The engine is pretty simple and minimalistic.
This example shows all it can do:

```go
package main

import "github.com/jsanchesleao/pixel"

func main() {
	// Initialize a new Engine with
	// Width of 320 pixels
	// Height of 240 pixels
	// Pixel size of 4
	// 60 Frames/Second
	engine, err := pixel.NewEngine("test", 320, 240, 4, 60)
	if err != nil {
		panic(err)
	}

	// This will decomission all graphics resources used by the engine
	defer engine.Destroy()

	// Define some state to be updated and rendered
	px, py := 0, 0

	// Update function.
	// Only operations related to updating the state should be used here
	// This runs once every iteration of the draw loop
	update := func(e *pixel.Engine) {
		if e.Inputs.W && py > 0 {
			py--
		}
		if e.Inputs.S && py < int(e.Height)-1 {
			py++
		}
		if e.Inputs.A && px > 0 {
			px--
		}
		if e.Inputs.D && px < int(e.Width)-1 {
			px++
		}
	}

	// Render function.
	// Here you can draw individual pixels on screen with the engine.Draw method
	// The canvas is preserved every loop iteration, so you have to manually clear it
	// Here we clear it with double for loop
	render := func(e *pixel.Engine) {
		for x := 0; x < int(e.Width); x++ {
			for y := 0; y < int(e.Height); y++ {
				if x == px && y == py {
					e.Draw(px, py, 0, 0, 255)
				} else {
					e.Draw(x, y, 0, 0, 0)
				}
			}
		}
	}

	// Starts the loop and also shows the window.
	engine.Loop(update, render)
}
```
