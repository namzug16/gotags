package htmx

import (
	"testing"

	gt "github.com/namzug16/gotags"
)

func TestDisableRendersBooleanAttribute(t *testing.T) {
	got := gt.Button(
		Disable(),
		"Do not process",
	).String()

	want := `<button hx-disable>Do not process</button>`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}
