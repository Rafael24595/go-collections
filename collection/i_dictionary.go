package collection

type DictionaryConstructor[T comparable, K any, D IDictionary[T, K]] func(map[T]K) D

type IDictionary[T comparable, K any] interface {
	Size() int
	Exists(key T) bool
	Find(predicate func(T, K) bool) []K
	FindOne(predicate func(T, K) bool) (*K, bool)
	Get(key T) (*K, bool)
	Put(key T, item K) (*K, bool)
	PutIfAbsent(key T, item K) (*K, bool)
	PutAll(items map[T]K) IDictionary[T, K]
	Merge(other IDictionary[T, K]) IDictionary[T, K]
	Filter(predicate func(T, K) bool) IDictionary[T, K]
	FilterSelf(predicate func(T, K) bool) IDictionary[T, K]
	Remove(key T) (*K, bool)
	ForEach(predicate func(T, K)) IDictionary[T, K]
	Map(predicate func(T, K) K) IDictionary[T, K]
	Clean() IDictionary[T, K]
	Clone() IDictionary[T, K]
	Keys() []T
	KeysVector() *Vector[T]
	Values() []K
	ValuesVector() *Vector[K]
	Pairs() []Pair[T, K]
	Collect() map[T]K
}

// IDictionaryMap creates a new IDictionary by applying the provided predicate function to each key-value pair in the original IDictionary.
// The predicate function is applied to each key and value, and its result is used as the new value in the returned IDictionary.
//
// Parameters:
//   - c: A pointer to the Dictionary[T, K] from which the key-value pairs will be transformed.
//   - predicate: A function that takes a key of type T and a value of type K, and returns a new value of type E. This function is applied to each key-value pair.
//   - constructor: A function that instance a new IDictionary implementation, and return it with the mapped values.
//
// Returns:
//   - A new IDictionary[T, E] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//	newDict := IDictionaryMap(dict, func(k string, v int) string { return fmt.Sprintf("%d", v) }, MakeDictionary)
//	// newDict will contain {"a": "1", "b": "2"}, where the values are transformed to strings
func IDictionaryMap[T comparable, K, E any, ID IDictionary[T, K], OD IDictionary[T, E]](c ID, predicate func(T, K) E, constructor DictionaryConstructor[T, E, OD]) OD {
	return MapToIDictionary(c.Collect(), predicate, constructor)
}

// MapToIDictionary creates a new Dictionary by applying the provided predicate function to each key-value pair in the provided map.
// The predicate function is applied to each key and value, and its result is used as the new value in the returned IDictionary.
//
// Parameters:
//   - c: A map[T]K from which the key-value pairs will be transformed.
//   - predicate: A function that takes a key of type T and a value of type K, and returns a new value of type E. This function is applied to each key-value pair.
//   - constructor: A function that instance a new IDictionary implementation, and return it with the mapped values.
//
// Returns:
//   - A new IDictionary[T, E] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	dict := map[string]int{"a": 1, "b": 2}
//	newDict := MapToIDictionary(dict, func(k string, v int) string { return fmt.Sprintf("%d", v) }, MakeDictionary)
//	// newDict will contain {"a": "1", "b": "2"}, where the values are transformed to strings
func MapToIDictionary[T comparable, K, E any, OD IDictionary[T, E]](c map[T]K, predicate func(T, K) E, constructor DictionaryConstructor[T, E, OD]) OD {
	mapped := map[T]E{}
	for key, item := range c {
		mapped[key] = predicate(key, item)
	}
	return constructor(mapped)
}

// VectorMapToIDictionary applies the given predicate function to each element in the IVector,
// transforming each element of type T into an tuple of types E, that implements comparable, and K, then returns
// a new IDictionary with the transformed elements.
//
// Parameters:
//   - c: The source IVector containing elements of type T.
//   - predicate: A function that takes an element of type T and transforms it into an element of type K.
//   - constructor: A function that instance a new IDictionary implementation, and return it with the mapped values.
//
// Returns:
//   - A new IDictionary[E, K] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	vec := VectorFromList([]int{1, 2, 3, 4})
//	transformed := VectorMapToIDictionary(vec, func(v int) (string, int) { return fmt.Sprintf("Item %d", v), v }, MakeDictionary)
//	// transformed will be a new Vector with elements: {"Item 1": 1, "Item 2": 2, "Item 3": 3, "Item 4": 4}
func VectorMapToIDictionary[T, K any, E comparable, OD IDictionary[E, K]](c IVector[T], predicate func(T) (E, K), constructor DictionaryConstructor[E, K, OD]) IDictionary[E, K] {
	return ListMapToIDictionary(c.Collect(), predicate, constructor)
}

// ListMapToIDictionary applies the given predicate function to each element in the slice,
// transformng each element of type T into an tuple of types E, that implements comparable, and K, then returns
// a new IDictionary with the transformed elements.
//
// Parameters:
//   - c: The slice IVector containing elements of type T.
//   - predicate: A function that takes an element of type T and transforms it into an element of type K.
//   - constructor: A function that instance a new IDictionary implementation, and return it with the mapped values.
//
// Returns:
//   - A new IDictionary[E, K] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	slc := []int{1, 2, 3, 4}
//	transformed := ListMapToIDictionary(slc, func(v int) (string, int) { return fmt.Sprintf("Item %d", v), v }, MakeDictionary)
//	// transformed will be a new Vector with elements: {"Item 1": 1, "Item 2": 2, "Item 3": 3, "Item 4": 4}
func ListMapToIDictionary[T, K any, E comparable, OD IDictionary[E, K]](c []T, predicate func(T) (E, K), constructor DictionaryConstructor[E, K, OD]) IDictionary[E, K] {
	m := make(map[E]K, 0)

	for _, v := range c {
		e, k := predicate(v)
		m[e] = k
	}

	return constructor(m)
}
