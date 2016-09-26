package spartan

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/direction"
)

type Divider struct {
	View
	Direction direction.Type
}

func (d Divider) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	if d.Direction == direction.Vertical {
		vertical := '\u2502'
		for i := top; i <= bottom; i++ {
			termbox.SetCell(int(left), int(i), vertical, d.ForegroundColor, d.BackgroundColor)
		}
	} else {
		horizontal := '\u2500'
		for i := left; i <= right; i++ {
			termbox.SetCell(int(i), int(top), horizontal, d.ForegroundColor, d.BackgroundColor)
		}
	}
}
