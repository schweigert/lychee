package spec

import "fmt"

func Expect[T any](value T, f func(T) error) error {
	return f(value)
}

func Eq[T comparable](value T) func(T) error {
	return func(expected T) error {
		if value == expected {
			return nil
		}

		return fmt.Errorf("value(%v) != expected(%v)", value, expected)
	}
}
