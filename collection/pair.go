package collection

// Pair represents a simple key-value pair, where the key is of type T and the value is of type K.
// This type is useful for storing and working with individual key-value pairs in various contexts, such as in a Dictionary.
//
// Type parameters:
//   - T: The type of the key in the Pair.
//   - K: The type of the value in the Pair.
//
// Fields:
//   - key: The key of the Pair, of type T.
//   - value: The value of the Pair, of type K.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1})
//     pair := dict.Pairs()[0]
//     fmt.Println(pair.key)   // Outputs: "a"
//     fmt.Println(pair.value) // Outputs: 1
type Pair[T, K any] struct {
	key T
	value K
}

func NewPair[T, K any](key T, value K) Pair[T, K] {
	return Pair[T, K]{
		key: key,
		value: value,
	}
}

// Key returns the key of the Pair.
//
// Returns:
//   - The key of type T from the Pair.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1})
//     pair := dict.Pairs()[0]
//     key := pair.Key() // key will be "a"
func (p Pair[T, K]) Key() T {
	return p.key
}

// Value returns the value of the Pair.
//
// Returns:
//   - The value of type K from the Pair.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1})
//     pair := dict.Pairs()[0]
//     value := pair.Value() // value will be 1
func (p Pair[T, K]) Value() K {
	return p.value
}