package spartan

import "github.com/nsf/termbox-go"

type app struct {
	layout ViewGroup
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
	a.layout.draw()
}

func (a *app) SetLayout(group ViewGroup) {
	a.layout = group
}
