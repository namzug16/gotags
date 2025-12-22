package gotags

import "strings"

type DocComponent struct {
	Components []HTML
}

func (d *DocComponent) String() string {
	var sb strings.Builder

	html := NewTagComponent("html", false, d.Components...)

	sb.WriteString("<!DOCTYPE html>")

	sb.WriteString(html.String())

	return sb.String()
}

func Doc(components ...HTML) *DocComponent {
	return &DocComponent{
		Components: components,
	}
}
