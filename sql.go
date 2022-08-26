package option

import (
	"database/sql/driver"
	"fmt"
)

// Value implements the driver.Valuer interface.
// To safely get the value of an Option, use GetValue().
func (t Option[T]) Value() (driver.Value, error) {
	value, ok := t.GetValue()
	if !ok {
		return nil, nil
	}

	return value, nil
}

// Scan implements the sql.Scanner interface.
// It means that the Option can be scanned from a database value.
func (t *Option[T]) Scan(v interface{}) error {
	if v == nil {
		*t = None[T]()
		return nil
	}

	switch vt := v.(type) {
	case T:
		*t = Some(vt)

	default:
		return fmt.Errorf("unsupported type: %T", v)
	}

	return nil
}
