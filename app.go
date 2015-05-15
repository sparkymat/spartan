package spartan

import "github.com/nsf/termbox-go"

type app struct {
	views []View
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
	for _, view := range a.views {
		view.draw()
	}
}

func (a *app) AddView(view View) {
	a.views = append(a.views, view)
}
