package gotags

type IfComponent struct {
	condition  bool
	components []HTML
}

func (d *IfComponent) AddToTag(target HTML) HTML {
	if !d.condition {
		return target
	}

	return AddToTag(target, d.components...)
}

func (d *IfComponent) String() string {
	// IfComponent cannot be composed directly because it may conditionally
	// contribute tags or attributes. It must be added to a TagComponent.
	panic("IfComponent cannot be composed directly; add it to a tag via AddToTag")
}

func If(condition bool, components ...HTML) *IfComponent {
	return &IfComponent{
		condition:  condition,
		components: components,
	}
}
