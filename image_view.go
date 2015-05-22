package spartan

import (
	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/gravity"
	"github.com/sparkymat/spartan/size"
)

type ImageView struct {
	parent          ViewGroup
	imagePath       string
	foregroundColor termbox.Attribute
	backgroundColor termbox.Attribute

	leftMargin   uint32
	topMargin    uint32
	rightMargin  uint32
	bottomMargin uint32

	width  size.Size
	height size.Size

	layoutGravity gravity.Type
}

func (box ImageView) draw(left uint32, top uint32, right uint32, bottom uint32) {

	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			char := ' '
			termbox.SetCell(int(i), int(j), char, box.foregroundColor, box.backgroundColor)
		}
	}
}

func (box *ImageView) SetWidth(width size.Size) {
	box.width = width
}

func (box *ImageView) SetHeight(height size.Size) {
	box.height = height
}

func (box ImageView) GetWidth() size.Size {
	return box.width
}

func (box ImageView) GetHeight() size.Size {
	return box.height
}

func (box *ImageView) SetLeftMargin(leftMargin uint32) {
	box.leftMargin = leftMargin
}

func (box *ImageView) SetTopMargin(topMargin uint32) {
	box.topMargin = topMargin
}

func (box *ImageView) SetRightMargin(rightMargin uint32) {
	box.rightMargin = rightMargin
}

func (box *ImageView) SetBottomMargin(bottomMargin uint32) {
	box.bottomMargin = bottomMargin
}

func (box ImageView) GetLeftMargin() uint32 {
	return box.leftMargin
}

func (box ImageView) GetTopMargin() uint32 {
	return box.topMargin
}

func (box ImageView) GetRightMargin() uint32 {
	return box.rightMargin
}

func (box ImageView) GetBottomMargin() uint32 {
	return box.bottomMargin
}

func (box *ImageView) SetColor(color termbox.Attribute) {
	box.foregroundColor = color
}

func (box ImageView) GetColor() termbox.Attribute {
	return box.foregroundColor
}

func (box *ImageView) SetBackgroundColor(color termbox.Attribute) {
	box.backgroundColor = color
}

func (box ImageView) GetBackgroundColor() termbox.Attribute {
	return box.backgroundColor
}

func (box *ImageView) SetParent(parent ViewGroup) {
	box.parent = parent
}

func (box ImageView) GetParent() ViewGroup {
	return box.parent
}

func (box *ImageView) SetLayoutGravity(gravity gravity.Type) {
	box.layoutGravity = gravity
}

func (box ImageView) GetLayoutGravity() gravity.Type {
	return box.layoutGravity
}
