package spartan

import (
	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/size"
)

type TextBox struct {
	text            string
	foregroundColor termbox.Attribute
	backgroundColor termbox.Attribute

	left   uint32
	top    uint32
	width  uint32
	height uint32
}

func (box TextBox) draw() error {
	for i := box.Left(); i <= box.Right(); i++ {
		for j := box.Top(); j <= box.Bottom(); j++ {
			position := (j-box.Top())*box.Width() + (i - box.Left())
			if position < uint32(len(box.text)) {
				char := rune(box.text[position])
				termbox.SetCell(int(i), int(j), char, box.foregroundColor, box.backgroundColor)
			} else {
				termbox.SetCell(int(i), int(j), ' ', box.foregroundColor, box.backgroundColor)
			}
		}
	}

	return nil
}

func (box TextBox) Width() uint32 {
	return box.width
}

func (box TextBox) Height() uint32 {
	return box.height
}

func (box TextBox) Left() uint32 {
	return box.left
}

func (box TextBox) Top() uint32 {
	return box.top
}

func (box TextBox) Right() uint32 {
	return box.left + box.width - 1
}

func (box TextBox) Bottom() uint32 {
	return box.top + box.height - 1
}

func (box TextBox) SizeType() size.Type {
	return size.ResponsiveBoth
}

func (box *TextBox) ResizeTo(width uint32, height uint32) error {
	box.width = width
	box.height = height
	return nil
}

func (box *TextBox) PositionTo(left uint32, top uint32) error {
	box.left = left
	box.top = top
	return nil
}

func (box *TextBox) SetForegroundColor(color termbox.Attribute) {
	box.foregroundColor = color
}

func (box *TextBox) SetBackgroundColor(color termbox.Attribute) {
	box.backgroundColor = color
}
