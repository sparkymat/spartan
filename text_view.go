package spartan

import (
	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/gravity"
	"github.com/sparkymat/spartan/size"
)

type TextView struct {
	parent          ViewGroup
	text            string
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

func (box TextView) draw(left uint32, top uint32, right uint32, bottom uint32) {
	width := right - left + 1

	for i := left; i <= right; i++ {
		for j := top; j <= bottom; j++ {
			position := (j-top)*width + (i - left)
			char := ' '
			if position < uint32(len(box.text)) {
				char = rune(box.text[position])
			}
			termbox.SetCell(int(i), int(j), char, box.foregroundColor, box.backgroundColor)
		}
	}
}

func (box *TextView) SetWidth(width size.Size) {
	box.width = width
}

func (box *TextView) SetHeight(height size.Size) {
	box.height = height
}

func (box TextView) GetWidth() size.Size {
	return box.width
}

func (box TextView) GetHeight() size.Size {
	return box.height
}

func (box *TextView) SetLeftMargin(leftMargin uint32) {
	box.leftMargin = leftMargin
}

func (box *TextView) SetTopMargin(topMargin uint32) {
	box.topMargin = topMargin
}

func (box *TextView) SetRightMargin(rightMargin uint32) {
	box.rightMargin = rightMargin
}

func (box *TextView) SetBottomMargin(bottomMargin uint32) {
	box.bottomMargin = bottomMargin
}

func (box TextView) GetLeftMargin() uint32 {
	return box.leftMargin
}

func (box TextView) GetTopMargin() uint32 {
	return box.topMargin
}

func (box TextView) GetRightMargin() uint32 {
	return box.rightMargin
}

func (box TextView) GetBottomMargin() uint32 {
	return box.bottomMargin
}

func (box *TextView) SetColor(color termbox.Attribute) {
	box.foregroundColor = color
}

func (box TextView) GetColor() termbox.Attribute {
	return box.foregroundColor
}

func (box *TextView) SetBackgroundColor(color termbox.Attribute) {
	box.backgroundColor = color
}

func (box TextView) GetBackgroundColor() termbox.Attribute {
	return box.backgroundColor
}

func (box *TextView) SetParent(parent ViewGroup) {
	box.parent = parent
}

func (box TextView) GetParent() ViewGroup {
	return box.parent
}

func (box TextView) GetAbsoluteX() uint32 {
	if box.parent == nil {
		return box.leftMargin
	} else {
		return box.leftMargin + box.parent.GetAbsoluteX()
	}
}

func (box TextView) GetAbsoluteY() uint32 {
	if box.parent == nil {
		return box.topMargin
	} else {
		return box.topMargin + box.parent.GetAbsoluteY()
	}
}

func (box TextView) GetAbsoluteWidth() uint32 {
	return GetViewAbsoluteWidth(&box)
}

func (box TextView) GetAbsoluteHeight() uint32 {
	return GetViewAbsoluteHeight(&box)
}

func (box *TextView) SetLayoutGravity(gravity gravity.Type) {
	box.layoutGravity = gravity
}

func (box TextView) GetLayoutGravity() gravity.Type {
	return box.layoutGravity
}
