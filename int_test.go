package serialisation

import (
	"errors"
	"testing"
)

func TestEncodeInt(t *testing.T) {
	typeb := string(byte_type_int)
	cases := map[int]string{
		24:           typeb + "24",
		-24:          typeb + "-24",
		111111222222: typeb + "111111222222",
	}

	for v := range cases {
		res := encodeInt(v)
		if res != cases[v] {
			t.Errorf("Expected %s for value %d got %s", cases[v], v, res)
		}
	}
}

func TestDecodeIntOk(t *testing.T) {
	typeb := string(byte_type_int)
	cases := map[string]int{
		typeb + "33":           33,
		typeb + "-11":          -11,
		typeb + "222333222333": 222333222333,
	}

	for v := range cases {
		res, _, err := decodeInt(v, 0)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if res != cases[v] {
			t.Errorf("Expected %d for value %s got %d", cases[v], v, res)
		}
	}
}

func TestDecodeIntErr(t *testing.T) {
	{
		// no type byte in the begginning
		_, _, err := decodeInt("12312313", 0)
		if !errors.Is(err, ErrInvalidFormat) {
			t.Errorf("Expected to receive [%s] error, got [%s]", ErrInvalidFormat, err)
		}
	}
	{
		typeb := string(byte_type_int)
		// conversion error: not a valid int
		_, _, err := decodeInt(typeb+"1231.2313", 0)
		if !errors.Is(err, ErrConversion) {
			t.Errorf("Expected to receive [%s] error, got [%s]", ErrInvalidFormat, err)
		}
	}
}
