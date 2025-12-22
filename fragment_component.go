package gotags

import "strings"

type FragmentComponent struct {
	Components []HTML
}

func (f *FragmentComponent) String() string {
	var sb strings.Builder

	for _, c := range f.Components {
		sb.WriteString(c.String())
	}

	return sb.String()
}

func Fragment(components ...HTML) *FragmentComponent {
	return &FragmentComponent{
		Components: components,
	}
}
