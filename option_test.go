package option

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOption(t *testing.T) {
	option := Some(1)
	option2 := None[string]()
	option3 := TupleToOption(1, true)
	option4 := TupleToOption(1, false)

	require.True(t, option.IsPresent())
	require.False(t, option.IsAbsent())

	require.False(t, option2.IsPresent())
	require.True(t, option2.IsAbsent())

	require.True(t, option3.IsPresent())

	require.False(t, option4.IsPresent())
}

func TestString(t *testing.T) {
	option := Some("foo")
	option2 := None[string]()

	require.Equal(t, "Some(foo)", option.String())
	require.Equal(t, "None", option2.String())
}

func TestUpdate(t *testing.T) {
	t.Run("Update with value", func(t *testing.T) {
		option := Some(1)
		option.Update(2)
		require.Equal(t, 2, option.MustValue())
	})

	t.Run("Update without value", func(t *testing.T) {
		option2 := None[string]()
		option2.Update("foo")
		value, ok := option2.GetValue()
		require.True(t, ok)
		require.Equal(t, "foo", value)
	})
}

func TestApply(t *testing.T) {
	t.Run("Apply with value", func(t *testing.T) {
		option := Some("FoO")
		option.Apply(strings.ToLower)
		require.Equal(t, "foo", option.MustValue())
	})

	t.Run("Apply without value", func(t *testing.T) {
		option := None[string]()
		option.Apply(strings.ToLower)
		require.Equal(t, None[string](), option)
	})
}

func TestGet(t *testing.T) {
	option := Some(1)
	option2 := None[string]()

	got, ok := option.GetValue()
	require.Equal(t, 1, got)
	require.True(t, ok)

	got2, ok := option2.GetValue()
	require.Equal(t, "", got2)
	require.False(t, ok)
}

func TestMustGet(t *testing.T) {
	option := Some(1)
	option2 := None[string]()

	require.Equal(t, 1, option.MustValue())
	require.Panics(t, func() { option2.MustValue() })
}

func TestGetDefault(t *testing.T) {
	option := Some(1)
	option2 := None[string]()

	require.Equal(t, 1, option.ValueOrDefault(2))
	require.Equal(t, "foo", option2.ValueOrDefault("foo"))
}

type TestStruct struct {
	OptionalValueInt    Option[int]    `json:"optionalValueInt"`
	OptionalValueString Option[string] `json:"optionalValueString,omitempty"`
}
