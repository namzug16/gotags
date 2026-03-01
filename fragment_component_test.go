package gotags

import "testing"

func TestFragmentEmptyReturnsEmptyString(t *testing.T) {
	got := Fragment().String()

	if got != "" {
		t.Fatalf("expected empty output for empty Fragment, got %q", got)
	}
}

func TestFragmentConcatenatesComponentsInOrder(t *testing.T) {
	got := Fragment(
		Text("A"),
		Raw("<em>B</em>"),
		Span("C"),
	).String()

	want := `A<em>B</em><span>C</span>`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentSupportsNestedFragments(t *testing.T) {
	got := Fragment(
		Text("start-"),
		Fragment(Text("inner-1"), Text("inner-2")),
		Text("-end"),
	).String()

	want := `start-inner-1inner-2-end`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentPreservesChildEscapingBehavior(t *testing.T) {
	got := Fragment(
		Text("<unsafe>"),
		Raw("<strong>ok</strong>"),
	).String()

	want := `&lt;unsafe&gt;<strong>ok</strong>`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentCanBeComposedInsideTags(t *testing.T) {
	got := Div(
		Fragment(Span("one"), Span("two")),
		Text("three"),
	).String()

	want := `<div><span>one</span><span>two</span>three</div>`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentNilComponentPanicsOnString(t *testing.T) {
	f := Fragment(Text("ok"), nil)

	defer func() {
		if recover() == nil {
			t.Fatal("expected panic when Fragment contains nil component")
		}
	}()

	_ = f.String()
}

func TestFragmentTypedNilComponentPanicsOnString(t *testing.T) {
	var nilTag *TagComponent
	f := Fragment(nilTag)

	defer func() {
		if recover() == nil {
			t.Fatal("expected panic when Fragment contains typed nil component")
		}
	}()

	_ = f.String()
}

func TestFragmentWithIfComponentTrueRendersConditionalChildren(t *testing.T) {
	got := Fragment(
		Text("before"),
		If(true, Text("visible")),
		Text("after"),
	).String()

	want := `beforevisibleafter`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentWithIfComponentFalseSkipsConditionalChildren(t *testing.T) {
	got := Fragment(
		Text("before"),
		If(false, Text("hidden")),
		Text("after"),
	).String()

	want := `beforeafter`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}
