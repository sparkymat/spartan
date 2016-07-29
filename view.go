package spartan

import (
	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/gravity"
	"github.com/sparkymat/spartan/size"
)

type View struct {
	Parent          *ViewGroup
	ForegroundColor termbox.Attribute
	BackgroundColor termbox.Attribute

	LeftMargin   uint32
	TopMargin    uint32
	RightMargin  uint32
	BottomMargin uint32

	Width  size.Size
	Height size.Size

	LayoutGravity gravity.Type
}

func (v View) GetHeight() size.Size {
	return v.Height
}

func (v View) GetWidth() size.Size {
	return v.Width
}

func (v View) GetLeftMargin() uint32 {
	return v.LeftMargin
}

func (v View) GetRightMargin() uint32 {
	return v.RightMargin
}

func (v View) GetTopMargin() uint32 {
	return v.TopMargin
}

func (v View) GetBottomMargin() uint32 {
	return v.BottomMargin
}

func (v View) GetLayoutGravity() gravity.Type {
	return v.LayoutGravity
}
