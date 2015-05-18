package spartan

import (
	"testing"

	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/direction"
	"github.com/sparkymat/spartan/gravity"
	"github.com/sparkymat/spartan/size"
)

func TestSanity(t *testing.T) {
	app := New()

	layout := LinearLayout{}
	layout.SetDirection(direction.Vertical)
	layout.SetWidth(size.MatchParent)
	layout.SetHeight(size.MatchParent)

	helloBox := TextView{text: "Hello, World!"}
	helloBox.SetWidth(20)
	helloBox.SetHeight(1)
	helloBox.SetColor(termbox.ColorWhite)
	helloBox.SetBackgroundColor(termbox.ColorRed)

	triumphBox := TextView{text: "This was a triumph"}
	triumphBox.SetWidth(size.MatchParent)
	triumphBox.SetHeight(3)
	triumphBox.SetColor(termbox.ColorBlack)
	triumphBox.SetBackgroundColor(termbox.ColorMagenta)

	noteBox := TextView{text: "I am making a note here"}
	noteBox.SetWidth(6)
	noteBox.SetHeight(size.MatchParent)
	noteBox.SetLayoutGravity(gravity.Center)
	noteBox.SetColor(termbox.ColorRed)
	noteBox.SetBackgroundColor(termbox.ColorBlue)

	layout.AddView(&helloBox)
	layout.AddView(&triumphBox)
	layout.AddView(&noteBox)

	app.SetLayout(&layout)

	app.Run()
}
