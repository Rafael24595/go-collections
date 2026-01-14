package collection

type DictionaryConstructor[K comparable, V any, D IDictionary[K, V]] func(map[K]V) D

type IDictionary[K comparable, V any] interface {
	Size() int
	Exists(key K) bool
	Find(predicate func(K, V) bool) []V
	FindOne(predicate func(K, V) bool) (V, bool)
	Get(key K) (V, bool)
	Put(key K, item V) (V, bool)
	PutIfAbsent(key K, item V) (V, bool)
	PutAll(items map[K]V) IDictionary[K, V]
	Merge(other IDictionary[K, V]) IDictionary[K, V]
	Filter(predicate func(K, V) bool) IDictionary[K, V]
	FilterSelf(predicate func(K, V) bool) IDictionary[K, V]
	Remove(key K) (V, bool)
	ForEach(predicate func(K, V)) IDictionary[K, V]
	Map(predicate func(K, V) V) IDictionary[K, V]
	Clean() IDictionary[K, V]
	Clone() IDictionary[K, V]
	Max(predicate func(K, V) int) (Pair[K, V], int, bool)
	Min(predicate func(K, V) int) (Pair[K, V], int, bool)
	Keys() []K
	KeysVector() *Vector[K]
	Values() []V
	ValuesVector() *Vector[V]
	Pairs() []Pair[K, V]
	Collect() map[K]V
}

// IDictionaryMap creates a new IDictionary by applying the provided predicate function to each key-value pair in the original IDictionary.
// The predicate function is applied to each key and value, and its result is used as the new value in the returned IDictionary.
//
// Parameters:
//   - c: A pointer to the Dictionary[K, I] from which the key-value pairs will be transformed.
//   - predicate: A function that takes a key of type K and a value of type I, and returns a new value of type V. This function is applied to each key-value pair.
//   - constructor: A function that instance a new IDictionary implementation, and return it with the mapped values.
//
// Returns:
//   - A new IDictionary[K, V] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//	newDict := IDictionaryMap(dict, func(k string, v int) string { return fmt.Sprintf("%d", v) }, MakeDictionary)
//	// newDict will contain {"a": "1", "b": "2"}, where the values are transformed to strings
func IDictionaryMap[K comparable, I, V any, ID IDictionary[K, I], OD IDictionary[K, V]](c ID, predicate func(K, I) V, constructor DictionaryConstructor[K, V, OD]) OD {
	return MapToIDictionary(c.Collect(), predicate, constructor)
}

// MapToIDictionary creates a new Dictionary by applying the provided predicate function to each key-value pair in the provided map.
// The predicate function is applied to each key and value, and its result is used as the new value in the returned IDictionary.
//
// Parameters:
//   - c: A map[K]I from which the key-value pairs will be transformed.
//   - predicate: A function that takes a key of type K and a value of type I, and returns a new value of type V. This function is applied to each key-value pair.
//   - constructor: A function that instance a new IDictionary implementation, and return it with the mapped values.
//
// Returns:
//   - A new IDictionary[K, V] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	dict := map[string]int{"a": 1, "b": 2}
//	newDict := MapToIDictionary(dict, func(k string, v int) string { return fmt.Sprintf("%d", v) }, MakeDictionary)
//	// newDict will contain {"a": "1", "b": "2"}, where the values are transformed to strings
func MapToIDictionary[K comparable, I, V any, OD IDictionary[K, V]](c map[K]I, predicate func(K, I) V, constructor DictionaryConstructor[K, V, OD]) OD {
	mapped := map[K]V{}
	for key, item := range c {
		mapped[key] = predicate(key, item)
	}
	return constructor(mapped)
}

// VectorMapToIDictionary applies the given predicate function to each element in the IVector,
// transforming each element of type I into an tuple of types K, that implements comparable, and V, then returns
// a new IDictionary with the transformed elements.
//
// Parameters:
//   - c: The source IVector containing elements of type I.
//   - predicate: A function that takes an element of type I and transforms it into an element of type V.
//   - constructor: A function that instance a new IDictionary implementation, and return it with the mapped values.
//
// Returns:
//   - A new IDictionary[K, V] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	vec := VectorFromList([]int{1, 2, 3, 4})
//	transformed := VectorMapToIDictionary(vec, func(v int) (string, int) { return fmt.Sprintf("Item %d", v), v }, MakeDictionary)
//	// transformed will be a new Vector with elements: {"Item 1": 1, "Item 2": 2, "Item 3": 3, "Item 4": 4}
func VectorMapToIDictionary[I, V any, K comparable, OD IDictionary[K, V]](c IVector[I], predicate func(I) (K, V), constructor DictionaryConstructor[K, V, OD]) IDictionary[K, V] {
	return ListMapToIDictionary(c.Collect(), predicate, constructor)
}

// ListMapToIDictionary applies the given predicate function to each element in the slice,
// transformng each element of type I into an tuple of types K, that implements comparable, and V, then returns
// a new IDictionary with the transformed elements.
//
// Parameters:
//   - c: The slice IVector containing elements of type I.
//   - predicate: A function that takes an element of type I and transforms it into an element of type V.
//   - constructor: A function that instance a new IDictionary implementation, and return it with the mapped values.
//
// Returns:
//   - A new IDictionary[K, V] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	slc := []int{1, 2, 3, 4}
//	transformed := ListMapToIDictionary(slc, func(v int) (string, int) { return fmt.Sprintf("Item %d", v), v }, MakeDictionary)
//	// transformed will be a new Vector with elements: {"Item 1": 1, "Item 2": 2, "Item 3": 3, "Item 4": 4}
func ListMapToIDictionary[I, V any, K comparable, OD IDictionary[K, V]](c []I, predicate func(I) (K, V), constructor DictionaryConstructor[K, V, OD]) IDictionary[K, V] {
	m := make(map[K]V, 0)

	for _, v := range c {
		e, k := predicate(v)
		m[e] = k
	}

	return constructor(m)
}
