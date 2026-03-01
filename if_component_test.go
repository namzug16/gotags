package gotags

import "testing"

func TestIfComponentStringPanics(t *testing.T) {
	t.Helper()

	defer func() {
		recovered := recover()
		if recovered == nil {
			t.Fatal("expected panic from IfComponent.String")
		}

		if recovered != "IfComponent cannot be composed directly; add it to a tag via AddToTag" {
			t.Fatalf("unexpected panic message: %v", recovered)
		}
	}()

	_ = If(true, Text("hello")).String()
}

func TestIfTrueAddsChildrenAndAttributesToTag(t *testing.T) {
	tag := Div(
		If(true,
			X.Class("enabled"),
			Span("visible"),
		),
	)

	got := tag.String()
	want := `<div class="enabled"><span>visible</span></div>`

	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestIfFalseSkipsChildrenAndAttributes(t *testing.T) {
	tag := Div(
		X.Id("root"),
		If(false,
			X.Class("should-not-exist"),
			Span("hidden"),
		),
		Text("always"),
	)

	got := tag.String()
	want := `<div id="root">always</div>`

	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestNestedIfKeepsOrderForTruthyBranches(t *testing.T) {
	tag := Div(
		Text("A"),
		If(true,
			Text("B"),
			If(true, Text("C")),
		),
		Text("D"),
	)

	got := tag.String()
	want := `<div>ABCD</div>`

	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestIfWithoutComponentsIsNoOp(t *testing.T) {
	tag := Div(If(true), Text("ok"))

	got := tag.String()
	want := `<div>ok</div>`

	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}
