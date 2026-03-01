package gotags

import "strings"

type FragmentComponent struct {
	components []HTML
}

func (f *FragmentComponent) String() string {
	var sb strings.Builder

	for _, c := range f.components {
		ifComp, ok := c.(*IfComponent)
		if ok {
			if ifComp.condition {
				sb.WriteString(Fragment(ifComp.components...).String())
			}

			continue
		}

		sb.WriteString(c.String())
	}

	return sb.String()
}

func Fragment(components ...HTML) *FragmentComponent {
	return &FragmentComponent{
		components: components,
	}
}
