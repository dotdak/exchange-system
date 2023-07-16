package utils

func Any[K comparable](vals ...K) K {
	var zero K
	for _, v := range vals {
		if v != zero {
			return v
		}
	}

	return zero
}

func AnyError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
