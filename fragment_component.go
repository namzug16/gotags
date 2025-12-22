package gotags

import "strings"

type FragmentComponent struct {
	components []HTML
}

func (f *FragmentComponent) String() string {
	var sb strings.Builder

	for _, c := range f.components {
		sb.WriteString(c.String())
	}

	return sb.String()
}

func Fragment(components ...HTML) *FragmentComponent {
	return &FragmentComponent{
		components: components,
	}
}
