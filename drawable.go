package spartan

import (
	"github.com/sparkymat/spartan/gravity"
	"github.com/sparkymat/spartan/size"
)

type Drawable interface {
	Draw(left uint32, top uint32, right uint32, bottom uint32)

	GetLeftMargin() uint32
	GetRightMargin() uint32
	GetTopMargin() uint32
	GetBottomMargin() uint32
	GetHeight() size.Size
	GetWidth() size.Size
	GetLayoutGravity() gravity.Type

	OnStart()
}
