package spartan

import (
	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/size"
)

type View interface {
	GetWidth() size.Size
	GetHeight() size.Size
	SetWidth(width size.Size)
	SetHeight(height size.Size)

	SetLeftMargin(leftMargin uint32)
	SetTopMargin(topMargin uint32)
	GetLeftMargin() uint32
	GetTopMargin() uint32

	GetAbsoluteX() uint32
	GetAbsoluteY() uint32
	GetAbsoluteWidth() uint32
	GetAbsoluteHeight() uint32

	SetParent(parent ViewGroup)
	GetParent() ViewGroup

	SetColor(color termbox.Attribute)
	GetColor() termbox.Attribute
	SetBackgroundColor(color termbox.Attribute)
	GetBackgroundColor() termbox.Attribute

	draw()
}
