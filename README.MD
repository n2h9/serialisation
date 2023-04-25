## Description

Control bytes:
- `[` open array 
- `]` close array
- `,` separator of the elements in the array
- `\` esc byte

Type mark bytes:
- `i` integer type
- `f` float type 

Encoding rules:
- Integers are marked with `i` in the beginning
- Floats are marked with `f` in the beginning
- Arrays are surrounded with `[` `]`
- Elements in array are separated with `,`

Escaping rules:
- if string starts with type mark bytes  this bytes are prepend with `\`: `\i`, `\f`
- if string contains control bytes this bytes are prepend with `\`: `\[`, `\]`, `\,`, `\\`


## Usage

```go
package main

import (
	"fmt"
	"serialisation"
)

func main() {
	arr := []any{"a", "b", "c"}

	encoded := serialisation.Encode(arr)
	fmt.Println(encoded)

	decoded, err := serialisation.Decode(encoded)
	fmt.Println(decoded, err)
}
```


Go version >= 1.18
