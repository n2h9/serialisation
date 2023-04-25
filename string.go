package serialisation

import (
	"fmt"
	"strings"
)

func encodeString(v string) string {
	if len(v) <= 0 {
		return v
	}

	var s strings.Builder
	s.Grow(len(v))

	if isTypeByte(v[0]) || isControlByte(v[0]) {
		s.WriteByte(byte_esc)
	}
	s.WriteByte(v[0])

	for i := 1; i < len(v); i++ {
		if isControlByte(v[i]) {
			s.WriteByte(byte_esc)
		}
		s.WriteByte(v[i])
	}

	return s.String()
}

func decodeString(s string, i int) (val string, nexti int, err error) {
	var b strings.Builder

	for ; i < len(s); i++ {
		if s[i] == byte_sep || s[i] == byte_arr_close {
			break
		}
		if s[i] == byte_esc {
			if (i + 1) >= len(s) {
				return "", i + 1, fmt.Errorf(
					"%w: expected char after %c found end of line",
					ErrInvalidFormat,
					byte_esc,
				)
			}
			// skip byte_esc char, write next char to res
			i++
		}
		b.WriteByte(s[i])
	}

	return b.String(), i, nil
}
