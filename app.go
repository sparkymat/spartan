package spartan

import "github.com/nsf/termbox-go"

type app struct {
	content Drawable
}

func New() app {
	newApp := app{}
	return newApp
}

func (a app) Run() error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	termbox.SetOutputMode(termbox.Output256)

	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	a.Redraw()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventResize:
			a.Redraw()
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}

	return err
}

func (a app) Redraw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	a.draw()
	termbox.Flush()
}

func (a app) draw() {
	width, height := termbox.Size()
	a.content.Draw(0, 0, uint32(width)-1, uint32(height)-1)
}

func (a *app) SetContent(content Drawable) {
	a.content = content
}
