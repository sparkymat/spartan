package spartan

import "github.com/nsf/termbox-go"

type ImageView struct {
	View
	imagePath string
}

func (box ImageView) Draw(left uint32, top uint32, right uint32, bottom uint32) {

	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			char := ' '
			termbox.SetCell(int(i), int(j), char, box.ForegroundColor, box.BackgroundColor)
		}
	}
}
