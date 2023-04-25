package serialisation

import "strings"

func Encode(arr []any) string {
	res := make([]string, 0, len(arr))
	for _, val := range arr {
		if v, ok := val.(int); ok {
			res = append(res, encodeInt(v))
		}
		if v, ok := val.(string); ok {
			res = append(res, encodeString(v))
		}
		if v, ok := val.(float64); ok {
			res = append(res, encodeFloat(v))
		}
		if v, ok := val.([]any); ok {
			res = append(res, Encode(v))
		}
	}

	var s strings.Builder
	s.WriteByte(byte_arr_open)
	s.WriteString(strings.Join(res, byte_sep_str))
	s.WriteByte(byte_arr_close)
	return s.String()
}

func Decode(s string) ([]any, error) {
	val, _, err := decodeArr(s, 0)
	return val, err
}

func isControlByte(b byte) bool {
	return b == byte_arr_open || b == byte_arr_close || b == byte_sep || b == byte_esc
}

func isTypeByte(b byte) bool {
	return b == byte_type_int || b == byte_type_float
}
