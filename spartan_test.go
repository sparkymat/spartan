package spartan

import (
	"testing"

	"github.com/nsf/termbox-go"
)

func TestSanity(t *testing.T) {
	app := New()

	layout := LinearLayout{}

	helloBox := TextView{text: "Hello, World!"}
	helloBox.SetWidth(20)
	helloBox.SetHeight(1)
	helloBox.SetLeftMargin(4)
	helloBox.SetTopMargin(4)
	helloBox.SetColor(termbox.ColorWhite)
	helloBox.SetBackgroundColor(termbox.ColorRed)

	layout.AddView(&helloBox)
	app.SetLayout(&layout)

	app.Run()
}
