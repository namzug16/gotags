package gotags

import "testing"

func TestTagFunctionsNormalizeStringAndHTMLSliceContent(t *testing.T) {
	got := Div(
		"before",
		[]HTML{Span("<mid>"), Raw("<i>end</i>")},
	).String()

	want := `<div>before<span>&lt;mid&gt;</span><i>end</i></div>`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestTagFunctionsInvalidChildTypePanicsWithTagContext(t *testing.T) {
	defer func() {
		recovered := recover()
		if recovered == nil {
			t.Fatal("expected panic for invalid tag child type")
		}

		want := "gotags: invalid child type int; expected string, HTML or []HTML; in Tag <div>"
		if recovered != want {
			t.Fatalf("unexpected panic message:\nwant: %q\ngot:  %v", want, recovered)
		}
	}()

	_ = Div(123)
}

func TestVoidTagWithStringContentStillRendersVoidElement(t *testing.T) {
	got := Br("ignored").String()

	want := `<br />`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}
