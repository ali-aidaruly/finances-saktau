package slice

func Reduce[T any, K comparable](data []T, combine func(K, T) (K, error)) (K, error) {
	if len(data) == 0 {
		return zeroValue[K](), nil
	}

	var (
		err error
		acc = zeroValue[K]()
	)

	for _, d := range data {
		acc, err = combine(acc, d)
		if err != nil {
			return zeroValue[K](), err
		}

	}
	return acc, nil
}

// Helper function to get the zero value of a type
func zeroValue[T any]() T {
	var v T
	return v
}
