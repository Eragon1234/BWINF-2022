package utils

func Remove[T any](a []T, i int) []T {
	if len(a) <= i {
		return a
	}
	a = append(a[:i], a[i+1:]...)
	return a
}

func RemoveAndReturn[T any](a []T, i int) ([]T, T) {
	if len(a) <= i {
		return a, *new(T)
	}
	removedItem := a[i]
	a = Remove(a, i)
	return a, removedItem
}
