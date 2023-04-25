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
