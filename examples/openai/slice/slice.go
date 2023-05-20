package slice

func Map[T any, U any](xs []T, f func(T) U) []U {
	ys := make([]U, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func MapE[T any, U any](xs []T, f func(T) (U, error)) ([]U, error) {
	ys := make([]U, len(xs))
	for i, x := range xs {
		y, err := f(x)
		if err != nil {
			return ys, err
		}
		ys[i] = y
	}
	return ys, nil
}

func Filter[T any](xs []T, f func(T) bool) []T {
	ys := make([]T, 0)
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

func Find[T any](xs []T, f func(T) bool) (T, bool) {
	var x T
	for _, x = range xs {
		if f(x) {
			return x, true
		}
	}
	return x, false
}
