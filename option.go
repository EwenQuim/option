package optional

import (
	"encoding/json"
	"fmt"
)

// Some builds an Option when value is present.
func Some[T comparable](value T) Option[T] {
	return Option[T]{value: &value}
}

// None builds an Option when value is absent.
func None[T comparable]() Option[T] {
	return Option[T]{value: nil}
}

// TupleToOption converts a tuple to an Option.
func TupleToOption[T comparable](value T, ok bool) Option[T] {
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

// Value returns value and presence. If value is absent, returns zero value and false.
func (option Option[T]) Value() (T, bool) {
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
	res, ok := option.Value()
	if ok {
		return res
	}
	return fallback
}

// UnmarshalJSON implements json.Unmarshaler.
func (option *Option[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		option.value = nil
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	option.value = &value
	return nil
}

// MarshalJSON implements json.Marshaler.
func (option Option[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(option.value)
}
