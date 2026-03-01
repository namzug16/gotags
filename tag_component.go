package gotags

import (
	"fmt"
	"strings"
)

type TagComponent struct {
	name     string
	isVoid   bool
	attrs    []HTML
	children []HTML
}

func NewTagComponent(name string, isVoid bool, content ...HTML) *TagComponent {
	component := &TagComponent{
		name:     name,
		isVoid:   isVoid,
		attrs:    make([]HTML, 0),
		children: make([]HTML, 0),
	}

	AddToTag(component, content...)

	return component
}

func (t *TagComponent) String() string {
	var sb strings.Builder

	for _, attr := range t.attrs {
		sb.WriteString(" " + attr.String())
	}
	attrStr := sb.String()

	if t.isVoid {
		return fmt.Sprintf("<%s%s />", t.name, attrStr)
	}

	sb.Reset()
	for _, c := range t.children {
		sb.WriteString(c.String())
	}

	return fmt.Sprintf("<%s%s>%s</%s>", t.name, attrStr, sb.String(), t.name)
}

func AddToTag(target HTML, components ...HTML) HTML {
	t, ok := target.(*TagComponent)
	if !ok {
		panic("AddToTag: target is not a *TagComponent")
	}

	for i, c := range components {
		if c == nil {
			panic(fmt.Sprintf("AddToTag: nil component at index %d", i))
		}

		switch v := c.(type) {
		case *AttributeComponent:
			t.attrs = append(t.attrs, v)
		case *IfComponent:
			AddToTag(t, v.components...)
		default:
			t.children = append(t.children, c)
		}
	}

	return t
}
