package types

import "testing"

func TestDeserializeAndSerializeGoodDate(t *testing.T) {
	input := "\"2015-10-26\""
	d := MGDate{}

	d.UnmarshalJSON([]byte(input))
	if bytesOut, err := d.MarshalJSON(); err != nil {
		t.Error("Error marshalling value: %v", err)
	} else {
		stringOut := string(bytesOut)
		if input != stringOut {
			t.Error("Marshalled value did not match input (%v != %v)", stringOut, input)
		}
	}
}
