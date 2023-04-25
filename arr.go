package serialisation

import "fmt"

func decodeArr(s string, i int) (val []any, nexti int, err error) {
	if s[i] != byte_arr_open {
		return nil, i, fmt.Errorf(
			"%w: expected first byte to be %c found %c at position %d",
			ErrInvalidFormat,
			byte_arr_open,
			s[i],
			i,
		)
	}

	i++
loop:
	for i < len(s) {
		switch s[i] {
		case byte_sep:
			i++
		case byte_arr_close:
			break loop
		case byte_arr_open:
			v, nexti, err := decodeArr(s, i)
			if err != nil {
				return nil, nexti, err
			}
			val = append(val, v)
			i = nexti
		case byte_type_int:
			v, nexti, err := decodeInt(s, i)
			if err != nil {
				return nil, nexti, err
			}
			val = append(val, v)
			i = nexti
		case byte_type_float:
			v, nexti, err := decodeFloat(s, i)
			if err != nil {
				return nil, nexti, err
			}
			val = append(val, v)
			i = nexti
		default:
			v, nexti, err := decodeString(s, i)
			if err != nil {
				return nil, nexti, err
			}
			val = append(val, v)
			i = nexti
		}
	}

	if i >= len(s) {
		return nil, i, fmt.Errorf(
			"%w: expected %c found end of line",
			ErrInvalidFormat,
			byte_arr_close,
		)
	}

	i++
	return val, i, nil
}
