package spartan

import (
	"errors"

	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/size"
)

type LinearLayout struct {
	parent ViewGroup
	views  []View
	x      uint32
	y      uint32
	width  size.Size
	height size.Size
}

func (layout *LinearLayout) SetWidth(width size.Size) {
	layout.width = width
}

func (layout *LinearLayout) SetHeight(height size.Size) {
	layout.height = height
}

func (layout *LinearLayout) SetLeftMargin(x uint32) {
	layout.x = x
}

func (layout *LinearLayout) SetTopMargin(y uint32) {
	layout.y = y
}

func (layout *LinearLayout) AddView(view View) {
	layout.views = append(layout.views, view)
}

func (layout LinearLayout) GetChildCount() uint32 {
	return uint32(len(layout.views))
}

func (layout LinearLayout) GetChildAt(index uint32) (View, error) {
	if index >= layout.GetChildCount() {
		return nil, errors.New("index out of bounds")
	}

	return layout.views[index], nil
}

func (layout LinearLayout) GetAbsoluteX() uint32 {
	if layout.GetParent() == nil {
		return layout.x
	} else {
		return layout.GetParent().GetAbsoluteX() + layout.x
	}
}

func (layout LinearLayout) GetAbsoluteY() uint32 {
	if layout.GetParent() == nil {
		return layout.y
	} else {
		return layout.GetParent().GetAbsoluteY() + layout.y
	}
}

func (layout LinearLayout) GetAbsoluteWidth() uint32 {
	if layout.width == size.MatchParent {
		if layout.parent == nil {
			width, _ := termbox.Size()
			return uint32(width)
		} else {
			return layout.parent.GetAbsoluteWidth()
		}
	} else {
		return uint32(layout.width)
	}
}

func (layout LinearLayout) GetAbsoluteHeight() uint32 {
	if layout.height == size.MatchParent {
		if layout.parent == nil {
			_, height := termbox.Size()
			return uint32(height)
		} else {
			return layout.parent.GetAbsoluteHeight()
		}
	} else {
		return uint32(layout.height)
	}
}

func (layout LinearLayout) draw() {
	for _, view := range layout.views {
		view.draw()
	}
}

func (layout LinearLayout) GetParent() ViewGroup {
	return layout.parent
}

func (layout *LinearLayout) SetParent(parent ViewGroup) {
	layout.parent = parent
}
