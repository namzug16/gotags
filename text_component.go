package gotags

import "html"

type Text string

func (t Text) String() string { return html.EscapeString(string(t)) }
