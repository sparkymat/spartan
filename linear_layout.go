package spartan

import (
	"errors"

	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/direction"
	"github.com/sparkymat/spartan/gravity"
	"github.com/sparkymat/spartan/size"
)

type LinearLayout struct {
	parent    ViewGroup
	views     []View
	x         uint32
	y         uint32
	width     size.Size
	height    size.Size
	direction direction.Type
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
	if layout.direction == direction.Vertical {
		heights := make([]uint32, len(layout.views), len(layout.views))

		stretchiesCount := uint32(0)
		totalFixedHeight := uint32(0)

		for i, view := range layout.views {
			if view.GetHeight() == size.MatchParent {
				stretchiesCount += 1
			} else {
				heights[i] = uint32(view.GetHeight())
				totalFixedHeight += uint32(view.GetHeight())
			}
		}

		if stretchiesCount > 0 {

			for i, view := range layout.views {
				if view.GetHeight() == size.MatchParent {
					heights[i] = (layout.GetAbsoluteHeight() - totalFixedHeight) / stretchiesCount
				}
			}
		}

		currentTop := uint32(0)

		var parentWidth uint32

		if layout.parent == nil {
			pw, _ := termbox.Size()
			parentWidth = uint32(pw)
		} else {
			currentTop = layout.parent.GetAbsoluteY()
			parentWidth = layout.parent.GetAbsoluteWidth()
		}

		for i, view := range layout.views {
			height := heights[i]
			width := view.GetAbsoluteWidth()

			currentLeft := uint32(0)

			if view.GetLayoutGravity() == gravity.Left {
				currentLeft = view.GetLeftMargin()
			} else if view.GetLayoutGravity() == gravity.Right {
				currentLeft = (parentWidth - width - view.GetRightMargin())
			} else if view.GetLayoutGravity() == gravity.Middle {
				currentLeft = (parentWidth - width) / 2
			}

			view.draw(currentLeft, currentTop, currentLeft+width, currentTop+height)
			currentTop += height
		}
	} else if layout.direction == direction.Horizontal {
	}
}

func (layout LinearLayout) GetParent() ViewGroup {
	return layout.parent
}

func (layout *LinearLayout) SetParent(parent ViewGroup) {
	layout.parent = parent
}

func (layout *LinearLayout) SetDirection(direction direction.Type) {
	layout.direction = direction
}

func (layout LinearLayout) GetDirection() direction.Type {
	return layout.direction
}
