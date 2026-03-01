package gotags

import (
	"fmt"
	"testing"
)

func TestRangeRendersItemsInOrderWithIndex(t *testing.T) {
	items := []string{"alpha", "beta", "gamma"}

	got := Range(items, func(index int, item string) HTML {
		return Li(fmt.Sprintf("%d:%s", index, item))
	}).String()

	want := `<li>0:alpha</li><li>1:beta</li><li>2:gamma</li>`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestRangeEmptySliceDoesNotInvokeMapper(t *testing.T) {
	called := false

	got := Range([]int{}, func(index int, item int) HTML {
		called = true
		return Text("unexpected")
	}).String()

	if called {
		t.Fatal("expected mapper not to be called for empty slice")
	}

	if got != "" {
		t.Fatalf("expected empty output for empty slice, got %q", got)
	}
}

func TestRangeNilSliceDoesNotInvokeMapper(t *testing.T) {
	called := false

	var items []string
	got := Range(items, func(index int, item string) HTML {
		called = true
		return Text("unexpected")
	}).String()

	if called {
		t.Fatal("expected mapper not to be called for nil slice")
	}

	if got != "" {
		t.Fatalf("expected empty output for nil slice, got %q", got)
	}
}

func TestRangeMapperCanReturnDifferentComponentKinds(t *testing.T) {
	got := Range([]int{1, 2, 3}, func(index int, item int) HTML {
		if item%2 == 0 {
			return Raw("<em>even</em>")
		}
		return Text("odd")
	}).String()

	want := `odd<em>even</em>odd`
	if got != want {
		t.Fatalf("unexpected rendered HTML:\nwant: %q\ngot:  %q", want, got)
	}
}

func TestRangeMapperNilFunctionPanicsOnString(t *testing.T) {
	defer func() {
		recovered := recover()
		if recovered == nil {
			t.Fatal("expected panic when mapper is nil")
		}

		if recovered != "Range: toComponent mapper cannot be nil" {
			t.Fatalf("unexpected panic message: %v", recovered)
		}
	}()

	_ = Range([]int{1}, nil)
}

func TestRangeMapperReturningNilHTMLPanicsOnString(t *testing.T) {
	r := Range([]int{1}, func(index int, item int) HTML {
		return nil
	})

	defer func() {
		recovered := recover()
		if recovered == nil {
			t.Fatal("expected panic when mapper returns nil HTML")
		}

		if recovered != "Range: mapper returned nil HTML at index 0" {
			t.Fatalf("unexpected panic message: %v", recovered)
		}
	}()

	_ = r.String()
}
