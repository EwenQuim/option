package option

import (
	"encoding/json"
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

func TestMarshalOption(t *testing.T) {
	t.Run("Marshal with value", func(t *testing.T) {
		option := Some("foo")

		data, err := json.Marshal(option)
		require.NoError(t, err)
		require.Equal(t, `"foo"`, string(data))
	})

	t.Run("Marshal struct with value", func(t *testing.T) {
		testStruct := TestStruct{
			OptionalValueInt:    Some(1),
			OptionalValueString: Some("foo"),
		}

		data, err := json.Marshal(testStruct)
		require.NoError(t, err)
		require.Equal(t, `{"optionalValueInt":1,"optionalValueString":"foo"}`, string(data))
	})

	t.Run("Marshal without value", func(t *testing.T) {
		option := None[string]()
		data, err := json.Marshal(option)
		require.NoError(t, err)
		require.Equal(t, `null`, string(data))
	})

	t.Run("Marshal struct without value", func(t *testing.T) {
		testStruct := TestStruct{
			OptionalValueInt:    None[int](),
			OptionalValueString: Some("foo"),
		}

		data, err := json.Marshal(testStruct)
		require.NoError(t, err)
		require.Equal(t, `{"optionalValueInt":null,"optionalValueString":"foo"}`, string(data))
	})
}

func TestUnmarshalOption(t *testing.T) {
	t.Run("Unmarshal with value", func(t *testing.T) {
		var option Option[string]
		err := json.Unmarshal([]byte(`"foo"`), &option)
		require.NoError(t, err)
		require.Equal(t, Some("foo"), option)
	})

	t.Run("Unmarshal without value", func(t *testing.T) {
		var option Option[string]
		err := json.Unmarshal([]byte(`null`), &option)
		require.NoError(t, err)
		require.Equal(t, None[string](), option)
	})

	t.Run("Unmarshal struct with value", func(t *testing.T) {
		var testStruct TestStruct
		err := json.Unmarshal([]byte(`{"optionalValueInt":1,"optionalValueString":"foo"}`), &testStruct)
		require.NoError(t, err)
		require.Equal(t, TestStruct{
			OptionalValueInt:    Some(1),
			OptionalValueString: Some("foo"),
		}, testStruct)
	})

	t.Run("Unmarshal struct without int value", func(t *testing.T) {
		var testStruct TestStruct
		err := json.Unmarshal([]byte(`{"optionalValueInt":null,"optionalValueString":"foo"}`), &testStruct)
		require.NoError(t, err)
		require.Equal(t, TestStruct{
			OptionalValueInt:    None[int](),
			OptionalValueString: Some("foo"),
		}, testStruct)
	})

	t.Run("Unmarshal struct without struct value", func(t *testing.T) {
		var testStruct TestStruct
		err := json.Unmarshal([]byte(`{"optionalValueInt":1,"optionalValueString":null}`), &testStruct)
		require.NoError(t, err)
		require.Equal(t, TestStruct{
			OptionalValueInt:    Some(1),
			OptionalValueString: None[string](),
		}, testStruct)
	})

	t.Run("Unmarshal struct with value and empty optional", func(t *testing.T) {
		var testStruct TestStruct
		err := json.Unmarshal([]byte(`{"optionalValueInt":1}`), &testStruct)
		require.NoError(t, err)
		require.Equal(t, TestStruct{
			OptionalValueInt:    Some(1),
			OptionalValueString: None[string](),
		}, testStruct)
	})
}
