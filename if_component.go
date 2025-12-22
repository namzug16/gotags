package gotags

type IfComponent struct {
	condition  bool
	components []HTML
}

func (d *IfComponent) String() string {
	return Fragment(d.components...).String()
}

func If(condition bool, components ...HTML) *IfComponent {
	return &IfComponent{
		condition:  condition,
		components: components,
	}
}
