package spartan

import (
	"github.com/sparkymat/spartan/gravity"
	"github.com/sparkymat/spartan/size"
)

type View interface {
	GetWidth() size.Size
	GetHeight() size.Size
	SetWidth(width size.Size)
	SetHeight(height size.Size)

	SetLeftMargin(leftMargin uint32)
	SetTopMargin(topMargin uint32)
	SetRightMargin(rightMargin uint32)
	SetBottomMargin(bottomMargin uint32)
	GetLeftMargin() uint32
	GetTopMargin() uint32
	GetRightMargin() uint32
	GetBottomMargin() uint32

	SetParent(parent ViewGroup)
	GetParent() ViewGroup

	SetLayoutGravity(gravity gravity.Type)
	GetLayoutGravity() gravity.Type

	draw(left uint32, top uint32, right uint32, bottom uint32)
}
