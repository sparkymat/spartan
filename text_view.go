package spartan

import "github.com/nsf/termbox-go"

type TextView struct {
	View
	Text string
}

func (box TextView) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	width := right - left + 1

	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			position := (j-top)*width + (i - left)
			char := ' '
			if position < uint32(len(box.Text)) {
				char = rune(box.Text[position])
			}
			termbox.SetCell(int(i), int(j), char, box.ForegroundColor, box.BackgroundColor)
		}
	}
}
