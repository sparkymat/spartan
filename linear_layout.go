package spartan

import (
	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/direction"
	"github.com/sparkymat/spartan/gravity"
	"github.com/sparkymat/spartan/size"
)

type LinearLayout struct {
	parent        ViewGroup
	views         []View
	leftMargin    uint32
	topMargin     uint32
	rightMargin   uint32
	bottomMargin  uint32
	width         size.Size
	height        size.Size
	direction     direction.Type
	isBordered    bool
	title         string
	layoutGravity gravity.Type
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

func (layout *LinearLayout) SetLayoutGravity(gravity gravity.Type) {
	layout.layoutGravity = gravity
}

func (layout LinearLayout) GetLayoutGravity() gravity.Type {
	return layout.layoutGravity
}

func (layout *LinearLayout) SetWidth(width size.Size) {
	layout.width = width
}

func (layout LinearLayout) GetWidth() size.Size {
	return layout.width
}

func (layout *LinearLayout) SetHeight(height size.Size) {
	layout.height = height
}

func (layout LinearLayout) GetHeight() size.Size {
	return layout.height
}

func (layout *LinearLayout) SetLeftMargin(leftMargin uint32) {
	layout.leftMargin = leftMargin
}

func (layout *LinearLayout) SetTopMargin(topMargin uint32) {
	layout.topMargin = topMargin
}

func (layout *LinearLayout) SetRightMargin(rightMargin uint32) {
	layout.rightMargin = rightMargin
}

func (layout *LinearLayout) SetBottomMargin(bottomMargin uint32) {
	layout.bottomMargin = bottomMargin
}

func (layout LinearLayout) GetLeftMargin() uint32 {
	return layout.leftMargin
}

func (layout LinearLayout) GetTopMargin() uint32 {
	return layout.topMargin
}

func (layout LinearLayout) GetRightMargin() uint32 {
	return layout.rightMargin
}

func (layout LinearLayout) GetBottomMargin() uint32 {
	return layout.bottomMargin
}

func (layout *LinearLayout) AddView(view View) {
	layout.views = append(layout.views, view)
}

func (layout LinearLayout) GetAbsoluteX() uint32 {
	if layout.GetParent() == nil {
		return layout.leftMargin
	} else {
		return layout.GetParent().GetAbsoluteX() + layout.leftMargin
	}
}

func (layout LinearLayout) GetAbsoluteY() uint32 {
	if layout.GetParent() == nil {
		return layout.topMargin
	} else {
		return layout.GetParent().GetAbsoluteY() + layout.topMargin
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

func (layout LinearLayout) draw(left uint32, top uint32, right uint32, bottom uint32) {
	containerLeft := left
	containerTop := top
	containerRight := right
	containerBottom := bottom
	containerWidth := containerRight - containerLeft + 1
	containerHeight := containerBottom - containerTop + 1

	if layout.isBordered {
		containerLeft += 1
		containerRight -= 1
		containerWidth -= 2

		containerTop += 1
		containerBottom -= 1
		containerHeight -= 2

		layout.drawBorder(left, top, right, bottom)
		layout.drawTitle(left, top, right, bottom)
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

			var width uint32

			if view.GetWidth() == size.MatchParent || uint32(view.GetWidth()) > containerWidth {
				width = containerWidth
			} else {
				width = uint32(view.GetWidth())
			}

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
				if layout.isBordered {
					currentLeft += 1
				}
			}

			if currentRight > containerRight {
				currentRight = containerRight
				if layout.isBordered {
					currentRight -= 1
				}
			}

			if currentTop < containerTop {
				currentTop = containerTop
				if layout.isBordered {
					currentTop += 1
				}
			}

			if currentBottom > containerBottom {
				currentBottom = containerBottom
				if layout.isBordered {
					currentBottom -= 1
				}
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
			var height uint32

			if view.GetHeight() == size.MatchParent || uint32(view.GetHeight()) > containerHeight {
				height = containerHeight
			} else {
				height = uint32(view.GetHeight())
			}

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
				if layout.isBordered {
					currentLeft += 1
				}
			}

			if currentRight > containerRight {
				currentRight = containerRight
				if layout.isBordered {
					currentRight -= 1
				}
			}

			if currentTop < containerTop {
				currentTop = containerTop
				if layout.isBordered {
					currentTop += 1
				}
			}

			if currentBottom > containerBottom {
				currentBottom = containerBottom
				if layout.isBordered {
					currentBottom -= 1
				}
			}

			view.draw(currentLeft, currentTop, currentRight, currentBottom)
			currentLeft += width
		}
	}
}

func (layout LinearLayout) drawBorder(left uint32, top uint32, right uint32, bottom uint32) {
	leftTop := '\u250c'
	rightTop := '\u2510'
	leftBottom := '\u2514'
	rightBottom := '\u2518'

	termbox.SetCell(int(left), int(top), leftTop, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(int(right), int(top), rightTop, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(int(left), int(bottom), leftBottom, termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(int(right), int(bottom), rightBottom, termbox.ColorDefault, termbox.ColorDefault)

	// horizontal
	horizontal := '\u2500'
	for i := left + 1; i <= right-1; i++ {
		termbox.SetCell(int(i), int(top), horizontal, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(int(i), int(bottom), horizontal, termbox.ColorDefault, termbox.ColorDefault)
	}

	// vertical
	vertical := '\u2502'
	for j := top + 1; j <= bottom-1; j++ {
		termbox.SetCell(int(left), int(j), vertical, termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(int(right), int(j), vertical, termbox.ColorDefault, termbox.ColorDefault)
	}
}

func (layout LinearLayout) drawTitle(left uint32, top uint32, right uint32, bottom uint32) {
	length := len(layout.title)
	layoutWidth := right - left + 1
	start := int(left) + (int(layoutWidth)-length)/2
	end := start + length - 1

	for i := start; i <= end; i++ {
		termbox.SetCell(i, int(top), rune(layout.title[i-start]), termbox.ColorWhite|termbox.AttrBold, termbox.ColorDefault)
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
