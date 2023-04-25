package serialisation

import (
	"fmt"
	"strconv"
	"strings"
)

func encodeFloat(v float64) string {
	f := strconv.FormatFloat(v, 'f', -1, 64)
	var s strings.Builder
	s.WriteByte(byte_type_float)
	s.WriteString(f)

	return s.String()
}

func decodeFloat(s string, i int) (val float64, nexti int, err error) {
	if s[i] != byte_type_float {
		return -1, i, fmt.Errorf(
			"%w: expected first byte to be %c found %c at position %d",
			ErrInvalidFormat,
			byte_type_float,
			s[i],
			i,
		)
	}

	// get position right after the end of the number
	var j int
	for j = i + 1; j < len(s) && s[j] != byte_sep && s[j] != byte_arr_close; j++ {
	}

	val, err = strconv.ParseFloat(s[i+1:j], 64)
	if err != nil {
		return -1, i, fmt.Errorf(
			"%w: could not convert a float at postion %d: %s",
			ErrConversion,
			i,
			err,
		)
	}

	return val, j, nil
}
