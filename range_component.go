package gotags

import "fmt"

type RangeComponent[T any] struct {
	items       []T
	toComponent func(index int, item T) HTML
}

func (r *RangeComponent[T]) String() string {
	components := make([]HTML, len(r.items))

	for i, item := range r.items {
		component := r.toComponent(i, item)
		if component == nil {
			panic(fmt.Sprintf("Range: mapper returned nil HTML at index %d", i))
		}

		components[i] = component
	}

	return Fragment(components).String()
}

func Range[T any](items []T, toComponent func(int, T) HTML) *RangeComponent[T] {
	if toComponent == nil {
		panic("Range: toComponent mapper cannot be nil")
	}

	return &RangeComponent[T]{
		items:       items,
		toComponent: toComponent,
	}
}
