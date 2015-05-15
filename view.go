package spartan

import "github.com/sparkymat/spartan/size"

type View interface {
	draw() error
	Width() uint32
	Height() uint32
	SizeType() size.Type
	ResizeTo(width uint32, height uint32) error
	Left() uint32
	Top() uint32
	Right() uint32
	Bottom() uint32
	PositionTo(left uint32, top uint32) error
}
