package spartan

import "github.com/sparkymat/spartan/size"

type TextBox struct {
	text string

	left   uint32
	top    uint32
	width  uint32
	height uint32
}

func (box TextBox) draw() error {
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
	return box.left + box.width
}

func (box TextBox) Bottom() uint32 {
	return box.top + box.height
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
