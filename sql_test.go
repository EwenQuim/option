package option

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValue(t *testing.T) {
	t.Run("Value with a string value", func(t *testing.T) {
		option := Some("foo")

		value, err := option.Value()
		require.NoError(t, err)
		require.Equal(t, "foo", value)
	})

	t.Run("Value with an int value", func(t *testing.T) {
		option := Some(42)

		value, err := option.Value()
		require.NoError(t, err)
		require.Equal(t, 42, value)
	})

	t.Run("Value without value", func(t *testing.T) {
		option := None[string]()

		value, err := option.Value()
		require.NoError(t, err)
		require.Nil(t, value)
	})
}

func TestScan(t *testing.T) {
	t.Run("Scan with string value", func(t *testing.T) {
		option := None[string]()
		err := option.Scan("foo")
		require.NoError(t, err)
		require.Equal(t, Some("foo"), option)
	})

	t.Run("Scan with int value", func(t *testing.T) {
		option := None[int]()
		err := option.Scan(7)
		require.NoError(t, err)
		require.Equal(t, Some(7), option)
	})

	t.Run("Scan without value", func(t *testing.T) {
		option := None[string]()
		err := option.Scan(nil)
		require.NoError(t, err)
		require.Equal(t, None[string](), option)
	})

	t.Run("Scan without value on an initialized Option", func(t *testing.T) {
		option := Some("hi")
		err := option.Scan(nil)
		require.NoError(t, err)
		require.Equal(t, None[string](), option)
	})

	t.Run("Scan with unsupported type", func(t *testing.T) {
		option := None[string]()
		err := option.Scan(1)
		require.Error(t, err)
		require.Equal(t, None[string](), option)
	})

	t.Run("Scan with unsupported type", func(t *testing.T) {
		option := None[int]()
		err := option.Scan("foo")
		require.Error(t, err)
		require.Equal(t, None[int](), option)
	})
}
