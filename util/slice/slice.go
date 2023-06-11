package slice

import "errors"

func First[T any](values []T, fn func(v T) bool) (ret T, err error) {
	for _, v := range values {
		if fn(v) {
			return v, nil
		}
	}
	err = errors.New("no elements")
	return ret, err
}

func Where[T any](values []T, fn func(v T) bool) (result []T) {
	for _, v := range values {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func SelectString[T any](values []T, fn func(v T) string) (result []string) {
	for _, v := range values {
		result = append(result, fn(v))
	}
	return result
}
