package serialisation

import (
	"fmt"
	"strconv"
	"strings"
)

func encodeInt(v int) string {
	var s strings.Builder
	s.WriteByte(byte_type_int)
	s.WriteString(strconv.Itoa(v))

	return s.String()
}

func decodeInt(s string, i int) (val int, nexti int, err error) {
	if s[i] != byte_type_int {
		return -1, i, fmt.Errorf(
			"%w: expected first byte to be %c found %c at position %d",
			ErrInvalidFormat,
			byte_type_int,
			s[i],
			i,
		)
	}

	// get position right after the end of the number
	var j int
	for j = i + 1; j < len(s) && s[j] != byte_sep && s[j] != byte_arr_close; j++ {
	}

	val, err = strconv.Atoi(s[i+1 : j])
	if err != nil {
		return -1, i, fmt.Errorf(
			"%w: could not convert an int at postion %d: %s",
			ErrConversion,
			i,
			err,
		)
	}

	return val, j, nil
}
