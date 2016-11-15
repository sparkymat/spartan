package spartan

type ViewGroup struct {
	children []Drawable
}

func (group *ViewGroup) AddChild(child Drawable) {
	group.children = append(group.children, child)
}

func (group ViewGroup) GetAbsoluteX() uint32 {
	return 0
}

func (group ViewGroup) GetAbsoluteY() uint32 {
	return 0
}

func (group ViewGroup) GetAbsoluteWidth() uint32 {
	return 0
}

func (group ViewGroup) GetAbsoluteHeight() uint32 {
	return 0
}

func (layout LinearLayout) OnStart() {
	for _, c := range layout.children {
		c.OnStart()
	}
}
