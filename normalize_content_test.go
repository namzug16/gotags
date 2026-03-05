package gotags

import "testing"

func TestNormalizeContentEmptyInputReturnsEmptySlice(t *testing.T) {
	got, err := normalizeContent([]any{})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != 0 {
		t.Fatalf("expected empty result, got len=%d", len(got))
	}
}

func TestNormalizeContentConvertsStringsAndFlattensHTMLSlices(t *testing.T) {
	got, err := normalizeContent([]any{
		"plain",
		[]HTML{Span("nested"), Raw("<b>raw</b>")},
		Text("tail"),
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != 4 {
		t.Fatalf("unexpected normalized length: want 4, got %d", len(got))
	}

	if got[0].String() != "plain" {
		t.Fatalf("unexpected first item: %q", got[0].String())
	}

	if got[1].String() != "<span>nested</span>" {
		t.Fatalf("unexpected second item: %q", got[1].String())
	}

	if got[2].String() != "<b>raw</b>" {
		t.Fatalf("unexpected third item: %q", got[2].String())
	}

	if got[3].String() != "tail" {
		t.Fatalf("unexpected fourth item: %q", got[3].String())
	}
}

func TestNormalizeContentInvalidTypeReturnsError(t *testing.T) {
	_, err := normalizeContent([]any{42})

	if err == nil {
		t.Fatal("expected error for invalid child type")
	}

	want := "gotags: invalid child type int; expected string, HTML or []HTML"
	if err.Error() != want {
		t.Fatalf("unexpected error:\nwant: %q\ngot:  %q", want, err.Error())
	}
}
