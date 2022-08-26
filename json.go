package option

import "encoding/json"

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
