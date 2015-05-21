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
	layout.SetDirection(direction.Horizontal)
	layout.SetWidth(size.MatchParent)
	layout.SetHeight(size.MatchParent)
	layout.EnableBorder()
	layout.SetTitle("spartan")

	menuLayout := LinearLayout{}
	menuLayout.SetDirection(direction.Vertical)
	menuLayout.SetWidth(size.MatchParent)
	menuLayout.SetHeight(size.MatchParent)
	menuLayout.EnableBorder()
	menuLayout.SetTitle("menu")

	contentLayout := LinearLayout{}
	contentLayout.SetDirection(direction.Vertical)
	contentLayout.SetWidth(size.MatchParent)
	contentLayout.SetHeight(size.MatchParent)
	contentLayout.EnableBorder()
	contentLayout.SetTitle("content")

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

	successBox := TextView{text: "Huge success"}
	successBox.SetWidth(15)
	successBox.SetHeight(size.MatchParent)
	successBox.SetRightMargin(10)
	successBox.SetLayoutGravity(gravity.Right)
	successBox.SetColor(termbox.ColorGreen)
	successBox.SetBackgroundColor(termbox.ColorYellow)

	contentLayout.AddView(&helloBox)
	contentLayout.AddView(&triumphBox)
	contentLayout.AddView(&noteBox)
	contentLayout.AddView(&successBox)

	layout.AddView(&menuLayout)
	layout.AddView(&contentLayout)

	app.SetLayout(&layout)

	app.Run()
}
