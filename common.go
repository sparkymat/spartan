package spartan

import (
	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/size"
)

type ChildView interface {
	GetParent() ViewGroup
	SetParent(parent ViewGroup)

	GetWidth() size.Size
	GetHeight() size.Size
}

func GetViewAbsoluteWidth(view ChildView) uint32 {
	if view.GetWidth() == size.MatchParent {
		if view.GetParent() == nil {
			width, _ := termbox.Size()
			return uint32(width)
		} else {
			return view.GetParent().GetAbsoluteWidth()
		}
	} else {
		return uint32(view.GetWidth())
	}
}

func GetViewAbsoluteHeight(view ChildView) uint32 {
	if view.GetHeight() == size.MatchParent {
		if view.GetParent() == nil {
			_, height := termbox.Size()
			return uint32(height)
		} else {
			return view.GetParent().GetAbsoluteHeight()
		}
	} else {
		return uint32(view.GetHeight())
	}
}
