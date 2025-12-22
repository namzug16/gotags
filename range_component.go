package gotags

type RangeComponent[T any] struct {
	items       []T
	toComponent func(index int, item T) HTML
}

func (r *RangeComponent[T]) String() string {
	components := make([]HTML, len(r.items))

	for i, item := range r.items {
		components[i] = r.toComponent(i, item)
	}

	return Fragment(components...).String()
}

func Range[T any](items []T, toComponent func(int, T) HTML) *RangeComponent[T] {
	return &RangeComponent[T]{
		items:       items,
		toComponent: toComponent,
	}
}
