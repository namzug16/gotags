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

func TestFragmentWithStringsAndHTMLSlicesNormalizesContent(t *testing.T) {
	got := Fragment(
		"before-",
		[]HTML{Span("one"), Text("<two>")},
		Raw("<i>three</i>"),
	).String()

	want := `before-<span>one</span>&lt;two&gt;<i>three</i>`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentInvalidChildTypePanicsWithContext(t *testing.T) {
	defer func() {
		recovered := recover()
		if recovered == nil {
			t.Fatal("expected panic when Fragment contains invalid child type")
		}

		want := "gotags: invalid child type int; expected string, HTML or []HTML; in Fragment"
		if recovered != want {
			t.Fatalf("unexpected panic message:\nwant: %q\ngot:  %v", want, recovered)
		}
	}()

	_ = Fragment(123)
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

func TestFragmentWithDeeplyNestedFragments(t *testing.T) {
	got := Fragment(
		Text("level0-"),
		Fragment(
			Text("level1-"),
			Fragment(Text("level2")),
		),
		Text("-end"),
	).String()

	want := `level0-level1-level2-end`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentWithMultipleNestedFragments(t *testing.T) {
	got := Fragment(
		Fragment(Text("one"), Text("two")),
		Fragment(Text("three"), Text("four")),
	).String()

	want := `onetwothreefour`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentContainingFragmentWithMixedContent(t *testing.T) {
	got := Fragment(
		Text("start-"),
		Fragment(
			Text("text"),
			Raw("<em>raw</em>"),
			Span("span"),
		),
		Text("-end"),
	).String()

	want := `start-text<em>raw</em><span>span</span>-end`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentWithNestedFragmentAsOnlyChild(t *testing.T) {
	got := Fragment(
		Fragment(Text("only-child")),
	).String()

	want := `only-child`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentWithNestedFragmentInsideTag(t *testing.T) {
	got := Div(
		X.Class("container"),
		Fragment(
			Span("inner-1"),
			Fragment(Span("deep-inner")),
			Span("inner-2"),
		),
	).String()

	want := `<div class="container"><span>inner-1</span><span>deep-inner</span><span>inner-2</span></div>`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentWithNestedFragmentAndIfComponents(t *testing.T) {
	got := Fragment(
		Text("a"),
		Fragment(
			If(true, Text("b")),
			If(false, Text("c")),
		),
		Text("d"),
	).String()

	want := `abd`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentWithNestedFragmentPreservesEscaping(t *testing.T) {
	got := Fragment(
		Fragment(
			Text("<script>"),
			Raw("<b>bold</b>"),
		),
	).String()

	want := `&lt;script&gt;<b>bold</b>`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestFragmentWithNilFragmentChildPanics(t *testing.T) {
	defer func() {
		recovered := recover()
		if recovered == nil {
			t.Fatal("expected panic when Fragment contains nil fragment")
		}
	}()

	var nestedFrag *FragmentComponent
	_ = Fragment(nestedFrag).String()
}
