package spartan

import (
	"github.com/nsf/termbox-go"
	"github.com/sparkymat/spartan/direction"
	"github.com/sparkymat/spartan/gravity"
	"github.com/sparkymat/spartan/size"
)

type LinearLayout struct {
	ViewGroup
	View

	Direction  direction.Type
	IsBordered bool
	Title      string
}

func (layout *LinearLayout) EnableBorder() {
	layout.IsBordered = true
}

func (layout *LinearLayout) DisableBorder() {
	layout.IsBordered = false
}

func (layout LinearLayout) GetAbsoluteX() uint32 {
	if layout.Parent == nil {
		return layout.LeftMargin
	} else {
		return layout.Parent.GetAbsoluteX() + layout.LeftMargin
	}
}

func (layout LinearLayout) GetAbsoluteY() uint32 {
	if layout.Parent == nil {
		return layout.TopMargin
	} else {
		return layout.Parent.GetAbsoluteY() + layout.TopMargin
	}
}

func (layout LinearLayout) GetAbsoluteWidth() uint32 {
	if layout.Width == size.MatchParent {
		if layout.Parent == nil {
			width, _ := termbox.Size()
			return uint32(width)
		} else {
			return layout.Parent.GetAbsoluteWidth()
		}
	} else {
		return uint32(layout.Width)
	}
}

func (layout LinearLayout) GetAbsoluteHeight() uint32 {
	if layout.Height == size.MatchParent {
		if layout.Parent == nil {
			_, height := termbox.Size()
			return uint32(height)
		} else {
			return layout.Parent.GetAbsoluteHeight()
		}
	} else {
		return uint32(layout.Height)
	}
}

func (layout LinearLayout) Draw(left uint32, top uint32, right uint32, bottom uint32) {
	containerLeft := left
	containerTop := top
	containerRight := right
	containerBottom := bottom
	containerWidth := containerRight - containerLeft + 1
	containerHeight := containerBottom - containerTop + 1

	if layout.IsBordered {
		containerLeft += 1
		containerRight -= 1
		containerWidth -= 2

		containerTop += 1
		containerBottom -= 1
		containerHeight -= 2

		layout.drawBorder(left, top, right, bottom)
		layout.drawTitle(left, top, right, bottom)
	}

	if layout.Direction == direction.Vertical {
		heights := make([]uint32, len(layout.children), len(layout.children))

		stretchiesCount := uint32(0)
		totalFixedHeight := uint32(0)

		for i, child := range layout.children {
			if child.GetHeight() == size.MatchParent {
				stretchiesCount += 1
			} else {
				heights[i] = uint32(child.GetHeight())
				totalFixedHeight += uint32(child.GetHeight())
			}
		}

		if stretchiesCount > 0 {

			for i, child := range layout.children {
				if child.GetHeight() == size.MatchParent {
					heights[i] = (containerHeight - totalFixedHeight) / stretchiesCount
				}
			}
		}

		currentTop := containerTop

		for i, child := range layout.children {
			height := heights[i]

			var width uint32

			if child.GetWidth() == size.MatchParent || uint32(child.GetWidth()) > containerWidth {
				width = containerWidth
			} else {
				width = uint32(child.GetWidth())
			}

			currentLeft := uint32(0)

			if child.GetLayoutGravity() == gravity.Left {
				currentLeft = containerLeft + child.GetLeftMargin()
			} else if child.GetLayoutGravity() == gravity.Right {
				currentLeft = containerLeft + (containerWidth - width - child.GetRightMargin())
			} else if child.GetLayoutGravity() == gravity.Middle {
				currentLeft = containerLeft + (containerWidth-width)/2
			}

			currentRight := currentLeft + width - 1
			currentBottom := currentTop + height - 1

			// Clip to container dimensions
			if currentLeft < containerLeft {
				currentLeft = containerLeft
				if layout.IsBordered {
					currentLeft += 1
				}
			}

			if currentRight > containerRight {
				currentRight = containerRight
				if layout.IsBordered {
					currentRight -= 1
				}
			}

			if currentTop < containerTop {
				currentTop = containerTop
				if layout.IsBordered {
					currentTop += 1
				}
			}

			if currentBottom > containerBottom {
				currentBottom = containerBottom
				if layout.IsBordered {
					currentBottom -= 1
				}
			}

			child.Draw(currentLeft, currentTop, currentRight, currentBottom)

			currentTop += height
		}
	} else if layout.Direction == direction.Horizontal {
		widths := make([]uint32, len(layout.children), len(layout.children))

		stretchiesCount := uint32(0)
		totalFixedWidth := uint32(0)

		for i, child := range layout.children {
			if child.GetWidth() == size.MatchParent {
				stretchiesCount += 1
			} else {
				widths[i] = uint32(child.GetWidth())
				totalFixedWidth += uint32(child.GetWidth())
			}
		}

		if stretchiesCount > 0 {

			for i, child := range layout.children {
				if child.GetWidth() == size.MatchParent {
					widths[i] = (containerWidth - totalFixedWidth) / stretchiesCount
				}
			}
		}

		currentLeft := containerLeft

		for i, child := range layout.children {
			width := widths[i]
			var height uint32

			if child.GetHeight() == size.MatchParent || uint32(child.GetHeight()) > containerHeight {
				height = containerHeight
			} else {
				height = uint32(child.GetHeight())
			}

			currentTop := uint32(0)

			if child.GetLayoutGravity() == gravity.Top {
				currentTop = containerTop + child.GetTopMargin()
			} else if child.GetLayoutGravity() == gravity.Bottom {
				currentTop = containerTop + (containerHeight - height - child.GetBottomMargin())
			} else if child.GetLayoutGravity() == gravity.Center {
				currentTop = containerTop + (containerHeight-height)/2
			}

			currentRight := currentLeft + width - 1
			currentBottom := currentTop + height - 1

			// Clip to container dimensions
			if currentLeft < containerLeft {
				currentLeft = containerLeft
				if layout.IsBordered {
					currentLeft += 1
				}
			}

			if currentRight > containerRight {
				currentRight = containerRight
				if layout.IsBordered {
					currentRight -= 1
				}
			}

			if currentTop < containerTop {
				currentTop = containerTop
				if layout.IsBordered {
					currentTop += 1
				}
			}

			if currentBottom > containerBottom {
				currentBottom = containerBottom
				if layout.IsBordered {
					currentBottom -= 1
				}
			}

			child.Draw(currentLeft, currentTop, currentRight, currentBottom)
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
	length := len(layout.Title)
	layoutWidth := right - left + 1
	start := int(left) + (int(layoutWidth)-length)/2
	end := start + length - 1

	for i := start; i <= end; i++ {
		termbox.SetCell(i, int(top), rune(layout.Title[i-start]), termbox.ColorWhite|termbox.AttrBold, termbox.ColorDefault)
	}

}
