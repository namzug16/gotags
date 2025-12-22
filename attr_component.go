package gotags

import (
	"fmt"
	"html"
	"strings"
)

type AttributeComponent struct {
	name  string
	values []string
}

func (a *AttributeComponent) String() string {
	if len(a.values) == 0 {
		return a.name
	}
	value := strings.Join(a.values, " ")
	return fmt.Sprintf(`%s="%s"`, a.name, html.EscapeString(value))
}
