package gotags

func If(condition bool, components ...any) []HTML {
	if condition {
		content, err := normalizeContent(components)

		if err != nil {
			panic("")
		}

		return content
	}

	return []HTML{}
}

func IfLazy(condition bool, component func() HTML) HTML {
	if condition {
		return component()
	}

	return Fragment()
}
