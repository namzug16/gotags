package gotags

import "testing"

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

func TestIfTrueNormalizesStringAndHTMLSliceContent(t *testing.T) {
	tag := Div(
		If(true,
			"before",
			[]HTML{Span("middle"), Text("<after>")},
		),
	)

	got := tag.String()
	want := `<div>before<span>middle</span>&lt;after&gt;</div>`

	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestIfFalseSkipsInvalidContent(t *testing.T) {
	tag := Div(
		Text("before"),
		If(false, 123),
		Text("after"),
	)

	got := tag.String()
	want := `<div>beforeafter</div>`

	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestIfTruePanicsOnInvalidContent(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic when If(true) contains invalid content")
		}
	}()

	_ = If(true, 123)
}

func TestIfLazyFalseSkipsCallback(t *testing.T) {
	called := false

	tag := Div(
		Text("before"),
		IfLazy(false, func() HTML {
			called = true
			return Span("hidden")
		}),
		Text("after"),
	)

	got := tag.String()
	want := `<div>beforeafter</div>`

	if called {
		t.Fatal("expected callback to not run when condition is false")
	}

	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestIfLazyTrueEvaluatesCallbackOnceAndRendersResult(t *testing.T) {
	called := 0

	tag := Div(
		IfLazy(true, func() HTML {
			called++
			return Span("visible")
		}),
	)

	got := tag.String()
	want := `<div><span>visible</span></div>`

	if called != 1 {
		t.Fatalf("expected callback to run exactly once, got %d", called)
	}

	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestIfLazyCanContributeAttribute(t *testing.T) {
	tag := Div(
		IfLazy(true, func() HTML {
			return X.Class("enabled")
		}),
		Text("ok"),
	)

	got := tag.String()
	want := `<div class="enabled">ok</div>`

	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestIfLazyFalseAllowsNilCallback(t *testing.T) {
	tag := Div(IfLazy(false, nil), Text("ok"))

	got := tag.String()
	want := `<div>ok</div>`

	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}
