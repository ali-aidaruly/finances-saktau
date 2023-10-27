package slice

func Map[T1, T2 any](sl []T1, cast func(T1) T2) []T2 {
	if sl == nil {
		return nil
	}

	res := make([]T2, len(sl))
	for i := range sl {
		res[i] = cast(sl[i])
	}

	return res
}

func MapOptional[T1, T2 any](sl []T1, cast func(T1) (T2, error)) ([]T2, error) {
	if sl == nil {
		return nil, nil
	}

	var (
		res = make([]T2, len(sl))
		err error
	)
	for i := range sl {
		res[i], err = cast(sl[i])
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
