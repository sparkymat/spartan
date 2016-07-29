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

	layout := LinearLayout{
		Direction:  direction.Horizontal,
		IsBordered: true,
		Title:      "spartan",
	}
	layout.Height = size.MatchParent
	layout.Width = size.MatchParent

	menuLayout := LinearLayout{
		Direction:  direction.Vertical,
		IsBordered: true,
		Title:      "menu",
	}
	menuLayout.Width = 24
	menuLayout.Height = size.MatchParent

	contentLayout := LinearLayout{
		Direction:  direction.Vertical,
		IsBordered: true,
		Title:      "content",
	}
	contentLayout.Width = size.MatchParent
	contentLayout.Height = size.MatchParent

	helloBox := TextView{
		Text: "Hello, World!",
	}
	helloBox.Width = 20
	helloBox.Height = 1
	helloBox.ForegroundColor = termbox.ColorWhite
	helloBox.BackgroundColor = termbox.ColorRed

	triumphBox := TextView{
		Text: "This was a triumph",
	}
	triumphBox.Width = size.MatchParent
	triumphBox.Height = 3
	triumphBox.ForegroundColor = termbox.ColorBlack
	triumphBox.BackgroundColor = termbox.ColorMagenta

	noteBox := TextView{
		Text: "I am making a note here",
	}
	noteBox.Width = 6
	noteBox.Height = size.MatchParent
	noteBox.LayoutGravity = gravity.Center
	noteBox.ForegroundColor = termbox.ColorRed
	noteBox.BackgroundColor = termbox.ColorBlue

	finalBox := ImageView{
		ImagePath: "hello.jpg",
	}
	finalBox.Width = size.MatchParent
	finalBox.Height = size.MatchParent
	finalBox.ForegroundColor = termbox.ColorGreen
	finalBox.BackgroundColor = termbox.ColorYellow

	contentLayout.AddChild(&helloBox)
	contentLayout.AddChild(&triumphBox)
	contentLayout.AddChild(&noteBox)
	contentLayout.AddChild(&finalBox)

	layout.AddChild(&menuLayout)
	layout.AddChild(&contentLayout)

	app.SetContent(&layout)

	app.Run()
}
