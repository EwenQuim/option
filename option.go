package option

import (
	"fmt"
)

// Some builds an Option when value is present.
func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

// None builds an Option when value is absent.
func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

// TupleToOption converts a tuple to an Option.
func TupleToOption[T any](value T, ok bool) Option[T] {
	if ok {
		return Some(value)
	}

	return None[T]()
}

// Option is a container for an optional value of type T.
// It uses nil to represent the absence of a value.
type Option[T any] struct {
	value *T
}

// String returns a string representation of the Option.
func (option Option[T]) String() string {
	if option.IsAbsent() {
		return "None"
	}

	return fmt.Sprintf("Some(%v)", *option.value)
}

// IsPresent returns true when value is absent.
func (option Option[T]) IsPresent() bool {
	return option.value != nil
}

// IsAbsent returns true when value is present.
func (option Option[T]) IsAbsent() bool {
	return option.value == nil
}

// GetValue returns value and presence. If value is absent, returns zero value and false.
func (option Option[T]) GetValue() (T, bool) {
	if option.IsPresent() {
		return *option.value, true
	}

	return *new(T), false
}

// MustValue returns value if present or panics instead.
func (option Option[T]) MustValue() T {
	return *option.value
}

// ValueOrDefault returns value if present or the default value given instead.
func (option Option[T]) ValueOrDefault(fallback T) T {
	res, ok := option.GetValue()
	if ok {
		return res
	}

	return fallback
}

// Update updates the value of the Option.
func (option *Option[T]) Update(value T) {
	if option.IsPresent() {
		*option.value = value
	}

	*option = Some(value)
}

// Apply applies the given function to the value of the Option.
func (option *Option[T]) Apply(f func(T) T) {
	if option.IsPresent() {
		*option.value = f(*option.value)
	}
}
