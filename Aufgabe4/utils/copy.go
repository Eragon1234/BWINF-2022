package utils

func CopyPointerSlice[T any](a []*T) []*T {
	var a2 []*T
	for _, e := range a {
		e2 := *e
		a2 = append(a2, &e2)
	}
	return a2
}
