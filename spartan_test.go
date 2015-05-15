package spartan

import (
	"testing"

	"github.com/nsf/termbox-go"
)

func TestSanity(t *testing.T) {
	app := New()

	helloBox := TextBox{text: "Hello, World!"}
	helloBox.ResizeTo(40, 2)
	helloBox.PositionTo(4, 4)
	helloBox.SetForegroundColor(termbox.ColorWhite)
	helloBox.SetBackgroundColor(termbox.ColorRed)

	app.AddView(&helloBox)

	app.Run()
}
