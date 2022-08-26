package option

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

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

	t.Run("Unmarshal invalid JSON with value and empty optional", func(t *testing.T) {
		var testStruct TestStruct
		err := json.Unmarshal([]byte(`{ not a JSON`), &testStruct)
		require.Error(t, err)
		require.Equal(t, TestStruct{
			OptionalValueInt:    None[int](),
			OptionalValueString: None[string](),
		}, testStruct)
	})
}
