package slices

// Unique returns a unique subset of the slice provided.
func Unique[T comparable](input []T) []T {
	u := []T{}
	m := make(map[T]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}
