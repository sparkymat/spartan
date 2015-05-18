package spartan

import "github.com/sparkymat/spartan/size"

type ViewGroup interface {
	AddView(view View)
	GetChildAt(index uint32) (View, error)
	GetChildCount() uint32

	GetAbsoluteX() uint32
	GetAbsoluteY() uint32
	GetAbsoluteWidth() uint32
	GetAbsoluteHeight() uint32

	SetWidth(width size.Size)
	SetHeight(height size.Size)
	SetLeftMargin(x uint32)
	SetTopMargin(y uint32)

	SetParent(parent ViewGroup)
	GetParent() ViewGroup

	draw()
}
