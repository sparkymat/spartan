package spartan

import (
	"errors"

	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/direction"
	"github.com/sparkymat/spartan/gravity"
	"github.com/sparkymat/spartan/size"
)

type LinearLayout struct {
	parent     ViewGroup
	views      []View
	x          uint32
	y          uint32
	width      size.Size
	height     size.Size
	direction  direction.Type
	isBordered bool
	title      string
}

func (layout *LinearLayout) SetTitle(title string) {
	layout.title = title
}

func (layout LinearLayout) GetTitle() string {
	return layout.title
}

func (layout *LinearLayout) EnableBorder() {
	layout.isBordered = true
}

func (layout *LinearLayout) DisableBorder() {
	layout.isBordered = false
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
	containerLeft := layout.GetAbsoluteX()
	containerTop := layout.GetAbsoluteWidth()
	containerWidth := layout.GetAbsoluteWidth()
	containerHeight := layout.GetAbsoluteHeight()
	containerRight := containerLeft + containerWidth
	containerBottom := containerTop + containerHeight

	if layout.isBordered {
		containerLeft += 1
		containerRight -= 1
		containerWidth -= 2

		containerTop += 1
		containerBottom -= 1
		containerHeight -= 2
	}

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
					heights[i] = (containerHeight - totalFixedHeight) / stretchiesCount
				}
			}
		}

		currentTop := containerTop

		for i, view := range layout.views {
			height := heights[i]
			width := view.GetAbsoluteWidth()

			currentLeft := uint32(0)

			if view.GetLayoutGravity() == gravity.Left {
				currentLeft = containerLeft + view.GetLeftMargin()
			} else if view.GetLayoutGravity() == gravity.Right {
				currentLeft = containerLeft + (containerWidth - width - view.GetRightMargin())
			} else if view.GetLayoutGravity() == gravity.Middle {
				currentLeft = containerLeft + (containerWidth-width)/2
			}

			currentRight := currentLeft + width - 1
			currentBottom := currentTop + height - 1

			// Clip to container dimensions
			if currentLeft < containerLeft {
				currentLeft = containerLeft
			}

			if currentRight > containerRight {
				currentRight = containerRight
			}

			if currentTop < containerTop {
				currentTop = containerTop
			}

			if currentBottom > containerBottom {
				currentBottom = containerBottom
			}

			view.draw(currentLeft, currentTop, currentRight, currentBottom)

			currentTop += height
		}
	} else if layout.direction == direction.Horizontal {
		widths := make([]uint32, len(layout.views), len(layout.views))

		stretchiesCount := uint32(0)
		totalFixedWidth := uint32(0)

		for i, view := range layout.views {
			if view.GetWidth() == size.MatchParent {
				stretchiesCount += 1
			} else {
				widths[i] = uint32(view.GetWidth())
				totalFixedWidth += uint32(view.GetWidth())
			}
		}

		if stretchiesCount > 0 {

			for i, view := range layout.views {
				if view.GetWidth() == size.MatchParent {
					widths[i] = (containerWidth - totalFixedWidth) / stretchiesCount
				}
			}
		}

		currentLeft := containerLeft

		for i, view := range layout.views {
			width := widths[i]
			height := view.GetAbsoluteHeight()

			currentTop := uint32(0)

			if view.GetLayoutGravity() == gravity.Top {
				currentTop = containerTop + view.GetTopMargin()
			} else if view.GetLayoutGravity() == gravity.Bottom {
				currentTop = containerTop + (containerHeight - height - view.GetBottomMargin())
			} else if view.GetLayoutGravity() == gravity.Center {
				currentTop = containerTop + (containerHeight-height)/2
			}

			currentRight := currentLeft + width - 1
			currentBottom := currentTop + height - 1

			// Clip to container dimensions
			if currentLeft < containerLeft {
				currentLeft = containerLeft
			}

			if currentRight > containerRight {
				currentRight = containerRight
			}

			if currentTop < containerTop {
				currentTop = containerTop
			}

			if currentBottom > containerBottom {
				currentBottom = containerBottom
			}

			view.draw(currentLeft, currentTop, currentLeft+width, currentTop+height)
			currentLeft += width
		}
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
