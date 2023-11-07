package slice

func Filter[T1 any](sl []T1, f func(T1) bool) []T1 {
	if sl == nil {
		return nil
	}

	res := make([]T1, 0, len(sl))
	for i := range sl {
		if f(sl[i]) {
			res = append(res, sl[i])
		}
	}

	return res
}
