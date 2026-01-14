package collection

// Pair represents a simple key-value pair, where the key is of type K and the value is of type V.
// This type is useful for storing and working with individual key-value pairs in various contexts, such as in a Dictionary.
//
// Type parameters:
//   - K: The type of the key in the Pair.
//   - V: The type of the value in the Pair.
//
// Fields:
//   - key: The key of the Pair, of type K.
//   - value: The value of the Pair, of type V.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1})
//     pair := dict.Pairs()[0]
//     fmt.Println(pair.key)   // Outputs: "a"
//     fmt.Println(pair.value) // Outputs: 1
type Pair[K, V any] struct {
	key K
	value V
}

func NewPair[K, V any](key K, value V) Pair[K, V] {
	return Pair[K, V]{
		key: key,
		value: value,
	}
}

// Key returns the key of the Pair.
//
// Returns:
//   - The key of type K from the Pair.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1})
//     pair := dict.Pairs()[0]
//     key := pair.Key() // key will be "a"
func (p Pair[K, V]) Key() K {
	return p.key
}

// Value returns the value of the Pair.
//
// Returns:
//   - The value of type V from the Pair.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1})
//     pair := dict.Pairs()[0]
//     value := pair.Value() // value will be 1
func (p Pair[K, V]) Value() V {
	return p.value
}