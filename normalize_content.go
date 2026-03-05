package gotags

import (
	"fmt"
)

func normalizeContent(content []any) ([]HTML, error) {
	res := []HTML{}
	for _, c := range content {
		switch v := c.(type) {
		case HTML:
			res = append(res, v)
		case []HTML:
			res = append(res, v...)
		case string:
			res = append(res, Text(v))
		default:
			return []HTML{}, fmt.Errorf(
				"gotags: invalid child type %T; expected string, HTML or []HTML",
				c,
			)
		}
	}
	return res, nil
}
