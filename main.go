package pixel

import (
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
	Inputs   Inputs
	window   *sdl.Window
	renderer *sdl.Renderer
	buffer   *sdl.Surface
}

type Inputs struct {
	A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z bool
	Escape, Tab, Shift, Alt, Ctrl, Comma, Dot, Slash, Semicolon, Backslash       bool
	LeftBracket, RightBracket                                                    bool
	Space, Backspace, Enter                                                      bool
	Option, Command, RightCtrl, Menu                                             bool
	Num1, Num2, Num3, Num4, Num5, Num6, Num7, Num8, Num9, Num0                   bool
	Minus, Equals                                                                bool
	Up, Left, Down, Right                                                        bool
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
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyboardEvent:
				if t.Repeat > 0 {
					break
				}
				switch t.Keysym.Sym {
				case 'q':
					e.Inputs.Q = t.State == sdl.PRESSED
				case 'w':
					e.Inputs.W = t.State == sdl.PRESSED
				case 'e':
					e.Inputs.E = t.State == sdl.PRESSED
				case 'r':
					e.Inputs.R = t.State == sdl.PRESSED
				case 't':
					e.Inputs.T = t.State == sdl.PRESSED
				case 'y':
					e.Inputs.Y = t.State == sdl.PRESSED
				case 'u':
					e.Inputs.U = t.State == sdl.PRESSED
				case 'i':
					e.Inputs.I = t.State == sdl.PRESSED
				case 'o':
					e.Inputs.O = t.State == sdl.PRESSED
				case 'p':
					e.Inputs.P = t.State == sdl.PRESSED
				case 'a':
					e.Inputs.A = t.State == sdl.PRESSED
				case 's':
					e.Inputs.S = t.State == sdl.PRESSED
				case 'd':
					e.Inputs.D = t.State == sdl.PRESSED
				case 'f':
					e.Inputs.F = t.State == sdl.PRESSED
				case 'g':
					e.Inputs.G = t.State == sdl.PRESSED
				case 'h':
					e.Inputs.H = t.State == sdl.PRESSED
				case 'j':
					e.Inputs.J = t.State == sdl.PRESSED
				case 'k':
					e.Inputs.K = t.State == sdl.PRESSED
				case 'l':
					e.Inputs.L = t.State == sdl.PRESSED
				case 'z':
					e.Inputs.Z = t.State == sdl.PRESSED
				case 'x':
					e.Inputs.X = t.State == sdl.PRESSED
				case 'c':
					e.Inputs.C = t.State == sdl.PRESSED
				case 'v':
					e.Inputs.V = t.State == sdl.PRESSED
				case 'b':
					e.Inputs.B = t.State == sdl.PRESSED
				case 'n':
					e.Inputs.N = t.State == sdl.PRESSED
				case 'm':
					e.Inputs.M = t.State == sdl.PRESSED
				case '\x1b': // Escape
					e.Inputs.Escape = t.State == sdl.PRESSED
				case '1':
					e.Inputs.Num1 = t.State == sdl.PRESSED
				case '2':
					e.Inputs.Num2 = t.State == sdl.PRESSED
				case '3':
					e.Inputs.Num3 = t.State == sdl.PRESSED
				case '4':
					e.Inputs.Num4 = t.State == sdl.PRESSED
				case '5':
					e.Inputs.Num5 = t.State == sdl.PRESSED
				case '6':
					e.Inputs.Num6 = t.State == sdl.PRESSED
				case '7':
					e.Inputs.Num7 = t.State == sdl.PRESSED
				case '8':
					e.Inputs.Num8 = t.State == sdl.PRESSED
				case '9':
					e.Inputs.Num9 = t.State == sdl.PRESSED
				case '0':
					e.Inputs.Num0 = t.State == sdl.PRESSED
				case '-':
					e.Inputs.Minus = t.State == sdl.PRESSED
				case '=':
					e.Inputs.Equals = t.State == sdl.PRESSED
				case '\t':
					e.Inputs.Tab = t.State == sdl.PRESSED
				case 1073742049: // shift
					e.Inputs.Shift = t.State == sdl.PRESSED
				case 1073742054: // alt
					e.Inputs.Alt = t.State == sdl.PRESSED
				case 1073742048: // ctrl
					e.Inputs.Ctrl = t.State == sdl.PRESSED
				case ',':
					e.Inputs.Comma = t.State == sdl.PRESSED
				case '.':
					e.Inputs.Dot = t.State == sdl.PRESSED
				case '/':
					e.Inputs.Slash = t.State == sdl.PRESSED
				case ';':
					e.Inputs.Semicolon = t.State == sdl.PRESSED
				case '\\':
					e.Inputs.Backslash = t.State == sdl.PRESSED
				case ' ':
					e.Inputs.Space = t.State == sdl.PRESSED
				case '\b':
					e.Inputs.Backspace = t.State == sdl.PRESSED
				case '\r':
					e.Inputs.Enter = t.State == sdl.PRESSED
				case 1073742051: //option
					e.Inputs.Option = t.State == sdl.PRESSED
				case 1073742050: //command
					e.Inputs.Command = t.State == sdl.PRESSED
				case 1073742052: // right ctrl
					e.Inputs.Command = t.State == sdl.PRESSED
				case 1073741925: // menu
					e.Inputs.Command = t.State == sdl.PRESSED
				case 1073741906: // up
					e.Inputs.Up = t.State == sdl.PRESSED
				case 1073741904: // left
					e.Inputs.Left = t.State == sdl.PRESSED
				case 1073741905: // down
					e.Inputs.Down = t.State == sdl.PRESSED
				case 1073741903: // right
					e.Inputs.Right = t.State == sdl.PRESSED
				}
			}
		}

		texture, _ := e.renderer.CreateTextureFromSurface(e.buffer)

		e.renderer.SetDrawColor(0.0, 0.0, 0.0, 1.0)
		e.renderer.Clear()

		update(e)
		if skips > 0 {
			skips--
			sdl.Delay(uint32(interval))
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
