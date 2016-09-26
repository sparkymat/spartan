package spartan

import "github.com/nsf/termbox-go"

type App struct {
	content Drawable
}

func New() App {
	newApp := App{}
	return newApp
}

func (a App) Run(eventChannel chan termbox.Event) error {
	err := termbox.Init()
	if err != nil {
		return err
	}
	termbox.SetOutputMode(termbox.Output256)

	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	a.Redraw()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventResize:
			a.Redraw()
		case termbox.EventKey:
			eventChannel <- ev
		case termbox.EventError:
			panic(ev.Err)
		}
	}

	return err
}

func (a App) Redraw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	a.draw()
	termbox.Flush()
}

func (a App) draw() {
	width, height := termbox.Size()
	a.content.Draw(0, 0, uint32(width)-1, uint32(height)-1)
}

func (a *App) SetContent(content Drawable) {
	a.content = content
}
