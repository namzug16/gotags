package gotags

import "fmt"

func Range[T any](items []T, toComponent func(int, T) HTML) []HTML {
	if toComponent == nil {
		panic("Range: toComponent mapper cannot be nil")
	}

	components := make([]HTML, len(items))

	for i, item := range items {
		component := toComponent(i, item)
		if component == nil {
			panic(fmt.Sprintf("Range: mapper returned nil HTML at index %d", i))
		}

		components[i] = component
	}

	return components
}
