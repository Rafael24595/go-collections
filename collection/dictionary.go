package collection

// Dictionary is a generic key-value store where each key is of type K and each value is of type V.
// The Dictionary provides methods to manipulate and interact with key-value pairs efficiently, including
// operations like adding, removing, and transforming pairs.
//
// Type parameters:
//   - K: The type of the keys in the Dictionary. The keys must be comparable.
//   - V: The type of the values in the Dictionary.
//
// Fields:
//   - items: A map storing the actual key-value pairs. The keys are of type K, and the values are of type V.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     dict.Put("c", 3)
//     value, exists := dict.Get("a") // value will be 1, exists will be true
type Dictionary[K comparable, V any] struct {
	items map[K]V
}

// MakeDictionary creates a new Dictionary from a given map.
// It takes a map with keys of type K and values of type V and
// returns a pointer to a IDictionary containing the same items.
//
// K must be a comparable type to be used as a map key.
// V can be any type.
//
// Example usage:
//     myMap := map[string]int{"a": 1, "b": 2}
//     dict := MakeDictionary(myMap)
func MakeDictionary[K comparable, V any](items map[K]V) IDictionary[K, V] {
	return &Dictionary[K, V]{
		items,
	}
}

// DictionaryFromMap creates a new Dictionary from a given map.
// It takes a map with keys of type K and values of type V and
// returns a pointer to a Dictionary containing the same items.
//
// K must be a comparable type to be used as a map key.
// V can be any type.
//
// Example usage:
//     myMap := map[string]int{"a": 1, "b": 2}
//     dict := DictionaryFromMap(myMap)
func DictionaryFromMap[K comparable, V any](items map[K]V) *Dictionary[K, V] {
	return &Dictionary[K, V]{
		items,
	}
}

// DictionaryEmpty creates and returns a new, empty Dictionary.
//
// K must be a comparable type to be used as a map key.
// V can be any type.
//
// Example usage:
//     emptyDict := DictionaryEmpty[string, int]()
func DictionaryEmpty[K comparable, V any]() *Dictionary[K, V] {
	return DictionaryFromMap(make(map[K]V))
}

// DictionaryFromVector creates a Dictionary from a Vector by applying a mapping function.
//
// K must be a comparable type to be used as a dictionary key.
// V can be any type.
//
// Parameters:
//   - vector: A Vector containing values of type V.
//   - mapper: A function that converts a value of type V into a key of type K.
//
// Returns:
//   - A pointer to a Dictionary[K, V].
//
// Example usage:
//     vec := VectorFromList([]int{10, 20, 30})
//     dict := DictionaryFromVector(vec, func(i int) string { return fmt.Sprintf("key-%d", i) })
func DictionaryFromVector[K comparable, V any](vector Vector[V], mapper func(V) K) *Dictionary[K, V] {
	return DictionaryFromList(vector.items, mapper)
}

// DictionaryFromList creates a Dictionary from a slice by applying a mapping function.
//
// K must be a comparable type to be used as a dictionary key.
// V can be any type.
//
// Parameters:
//   - vector: A slice of values of type V.
//   - mapper: A function that converts a value of type V into a key of type K.
//
// Returns:
//   - A pointer to a Dictionary[K, V] containing the mapped key-value pairs.
//
// Example usage:
//     list := []int{10, 20, 30}
//     dict := DictionaryFromList(list, func(i int) string { return fmt.Sprintf("key-%d", i)
func DictionaryFromList[K comparable, V any](vector []V, mapper func(V) K) *Dictionary[K, V] {
	mapp := DictionaryEmpty[K, V]()
	for _, v := range vector {
		mapp.Put(mapper(v), v)
	}
	return mapp
}

// Size returns the number of key-value pairs in the Dictionary.
//
// Returns:
//   - An integer representing the number of elements in the Dictionary.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     size := dict.Size() // size will be 2
func (c *Dictionary[K, V]) Size() int {
	return len(c.items)
}

// Exists checks if the given key exists in the Dictionary.
//
// Parameters:
//   - key: The key of type K to check for in the Dictionary.
//
// Returns:
//   - A boolean indicating whether the key exists in the Dictionary.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     exists := dict.Exists("a") // exists will be true
//     exists = dict.Exists("c")  // exists will be false
func (c *Dictionary[K, V]) Exists(key K) bool {
	_, exists := c.items[key]
	return exists
}

// Find returns a slice of values from the Dictionary that satisfy the given predicate function.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and returns a boolean.
//                The function should return true for the values that should be included in the result.
//
// Returns:
//   - A slice of values of type V that satisfy the predicate function.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     result := dict.Find(func(k string, v int) bool { return v > 1 })
//     // result will be [2, 3]
func (c *Dictionary[K, V]) Find(predicate func(K, V) bool) []V {
	filter := []V{}
	for k, v := range c.items {
		if predicate(k, v) {
			filter = append(filter, v)
		}
	}
	return filter
}

// FindOne searches for the first key-value pair in the Dictionary that satisfies the given predicate function.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and returns a boolean.
//                The function should return true for the first pair that matches the search criteria.
//
// Returns:
//   - A pointer to the value of type V if a matching key-value pair is found, or nil if not found.
//   - A boolean indicating whether a match was found (true if found, false otherwise).
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     value, found := dict.FindOne(func(k string, v int) bool { return v == 2 })
//     // value will be a pointer to 2, found will be true
//     value, found = dict.FindOne(func(k string, v int) bool { return v == 4 })
//     // value will be nil, found will be false
func (c *Dictionary[K, V]) FindOne(predicate func(K, V) bool) (V, bool) {
	for k, v := range c.items {
		if predicate(k, v) {
			return v, true
		}
	}
	var zero V
	return zero, false
}

// Get retrieves the value associated with the given key in the Dictionary.
// It returns a pointer to the value if the key exists, and a boolean indicating whether the key was found.
//
// Parameters:
//   - key: The key of type K whose associated value is to be retrieved.
//
// Returns:
//   - A pointer to the value of type V associated with the key, or nil if the key does not exist.
//   - A boolean indicating whether the key was found in the Dictionary (true if found, false otherwise).
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     value, found := dict.Get("a") // value will be a pointer to 1, found will be true
//     value, found = dict.Get("c")  // value will be nil, found will be false
func (c *Dictionary[K, V]) Get(key K) (V, bool) {
	value, exists := c.items[key]
	return value, exists
}

// Put adds a key-value pair to the Dictionary, updating the value if the key already exists.
// It returns the old value associated with the key, if any, and a boolean indicating whether
// the key already existed in the Dictionary (true if it existed, false otherwise).
//
// Parameters:
//   - key: The key of type K to associate with the given value.
//   - item: The value of type V to be associated with the key.
//
// Returns:
//   - A pointer to the old value associated with the key, or nil if the key did not exist.
//   - A boolean indicating whether the key was already present in the Dictionary (true if it existed).
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     oldValue, exists := dict.Put("a", 3) // oldValue will be a pointer to 1, exists will be true
//     oldValue, exists = dict.Put("c", 4)  // oldValue will be nil, exists will be false
func (c *Dictionary[K, V]) Put(key K, item V) (V, bool) {
	old, exists := c.Get(key)
	c.items[key] = item
	return old, exists
}

// PutIfAbsent adds a key-value pair to the Dictionary only if the key does not already exist.
// If the key is already present, it does nothing and returns the existing value associated with the key,
// along with a boolean indicating that the key was already present.
//
// Parameters:
//   - key: The key of type K to associate with the given value.
//   - item: The value of type V to be associated with the key if the key is absent.
//
// Returns:
//   - A pointer to the old value associated with the key, or nil if the key was not found.
//   - A boolean indicating whether the key was already present in the Dictionary (true if it existed, false if it was absent).
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     oldValue, exists := dict.PutIfAbsent("a", 3) // oldValue will be a pointer to 1, exists will be true
//     oldValue, exists = dict.PutIfAbsent("c", 4)  // oldValue will be nil, exists will be false
func (c *Dictionary[K, V]) PutIfAbsent(key K, item V) (V, bool) {
	old, exists := c.Get(key)
	if !exists {
		c.Put(key, item)
	}
	return old, exists
}

// PutAll adds all key-value pairs from another map to the Dictionary
// overwriting any existing values for the keys that already exist in the Dictionary.
//
// Parameters:
//   - items: A map of type map[K]V containing the key-value pairs to add to the Dictionary.
//
// Returns:
//   - The Dictionary itself, with all the new key-value pairs added.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     otherMap := map[string]int{"b": 3, "c": 4}
//     dict.PutAll(otherMap) // dict will contain {"a": 1, "b": 3, "c": 4}
func (c *Dictionary[K, V]) PutAll(items map[K]V) IDictionary[K, V] {
	for key := range items {
		c.items[key] = items[key]
	}
	return c
}

// Merge combines all key-value pairs from another Dictionary into the current Dictionary
// overwriting any existing values for the keys that already exist.
//
// Parameters:
//   - other: A Dictionary of type Dictionary[K, V] to merge into the current Dictionary.
//
// Returns:
//   - The Dictionary itself, with the key-value pairs from the other Dictionary added.
//
// Example usage:
//     dict1 := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     dict2 := DictionaryFromMap(map[string]int{"b": 3, "c": 4})
//     dict1.Merge(dict2) // dict1 will contain {"a": 1, "b": 3, "c": 4}
func (c *Dictionary[K, V]) Merge(other IDictionary[K, V]) IDictionary[K, V] {
	return c.PutAll(other.Collect())
}

// Filter creates a new Dictionary by filtering the key-value pairs in the current Dictionary
// based on the provided predicate function. It iterates over all key-value pairs and retains
// those that satisfy the condition defined in the predicate.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and returns a boolean.
//                The function should return true for the key-value pairs that should be kept in the result.
//
// Returns:
//   - A new Dictionary itself, after filtering out the key-value pairs that do not satisfy the predicate.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     filtered := dict.Filter(func(k string, v int) bool { return v > 1 })
//     // filtered will contain {"b": 2, "c": 3}
func (c *Dictionary[K, V]) Filter(predicate func(K, V) bool) IDictionary[K, V] {
	filter := map[K]V{}
	for key, v := range c.items {
		if predicate(key, v) {
			filter[key] = v
		}
	}
	return DictionaryFromMap(filter)
}

// FilterSelf filters the key-value pairs in the current Dictionary based on the provided predicate function.
// It updates the Dictionary itself, removing key-value pairs that do not satisfy the condition defined in the predicate.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and returns a boolean.
//                The function should return true for the key-value pairs that should be retained in the Dictionary.
//
// Returns:
//   - The Dictionary itself, with only the key-value pairs that satisfy the predicate.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     dict.FilterSelf(func(k string, v int) bool { return v > 1 })
//     // dict will contain {"b": 2, "c": 3}
func (c *Dictionary[K, V]) FilterSelf(predicate func(K, V) bool) IDictionary[K, V] {
	filter := map[K]V{}
	for key, v := range c.items {
		if predicate(key, v) {
			filter[key] = v
		}
	}
	c.items = filter
	return c
}

// Remove deletes a key-value pair from the Dictionary by the provided key.
// It returns the old value associated with the key, if it exists, along with a boolean
// indicating whether the key was found and removed from the Dictionary.
//
// Parameters:
//   - key: The key of type K to remove from the Dictionary.
//
// Returns:
//   - A pointer to the old value associated with the key, or nil if the key was not found.
//   - A boolean indicating whether the key was present and removed (true if removed, false if not).
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     oldValue, exists := dict.Remove("a", 1) // oldValue will be a pointer to 1, exists will be true
//     oldValue, exists = dict.Remove("c", 3)  // oldValue will be nil, exists will be false
func (c *Dictionary[K, V]) Remove(key K) (V, bool) {
	old, exists := c.Get(key)
	delete(c.items, key)
	return old, exists
}

// ForEach iterates over all key-value pairs in the Dictionary, applying the provided predicate function to each pair.
// The predicate is called with each key and value, allowing side effects or custom actions for every entry in the Dictionary.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and performs an action or operation.
//
// Returns:
//   - The Dictionary itself, allowing for method chaining.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     dict.ForEach(func(k string, v int) { fmt.Println(k, v) })
//     // Output:
//     // a 1
//     // b 2
func (c *Dictionary[K, V]) ForEach(predicate func(K, V)) IDictionary[K, V] {
	for k, v := range c.items {
		predicate(k, v)
	}
	return c
}

// Map transforms the values in the Dictionary by applying the provided predicate function to each key-value pair.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and returns a new value of type V.
//
// Returns:
//   - The Dictionary itself, with the transformed values.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     dict.Map(func(k string, v int) int { return v * 2 })
//     // dict will contain {"a": 2, "b": 4}
func (c *Dictionary[K, V]) Map(predicate func(K, V) V) IDictionary[K, V] {
	for k, v := range c.items {
		c.items[k] = predicate(k, v)
	}
	return c
}

// Clean removes all key-value pairs from the Dictionary, effectively clearing its contents.
// After calling this method, the Dictionary will be empty, and its size will be zero.
//
// Returns:
//   - The Dictionary itself, now empty, allowing for method chaining.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     dict.Clean() // dict will be empty: {}
func (c *Dictionary[K, V]) Clean() IDictionary[K, V] {
	c.items = make(map[K]V)
	return c
}

// Clone creates a shallow copy of the Dictionary, including all key-value pairs.
// The new Dictionary will have the same keys and values as the original, but modifications to one
// will not affect the other.
//
// Returns:
//   - A new Dictionary that is a clone of the current Dictionary.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     clonedDict := dict.Clone() // clonedDict is a new Dictionary with the same contents as dict
func (c *Dictionary[K, V]) Clone() IDictionary[K, V] {
	cloned := make(map[K]V)
	for k, v := range c.items {
		cloned[k] = v
	}
	return DictionaryFromMap(cloned)
}

// Max returns the key-value pair from the Dictionary that yields the maximum
// score when evaluated with the provided predicate function.
//
// The predicate function is applied to each key-value pair in the Dictionary.
// The pair that produces the highest integer score is returned along with
// that score.
//
// Due to the unordered nature of maps, if multiple pairs produce the same
// maximum score, the returned pair is not deterministic.
//
// Parameters:
//   - predicate: A function that takes a key and a value, and returns an
//     integer score used for comparison.
//
// Returns:
//   - A Pair containing the key and value with the maximum score.
//   - The maximum score returned by the predicate.
//   - A boolean indicating whether the Dictionary was non-empty.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"go": 14, "rust": 11, "zig": 3})
//     pair, score, ok := dict.Max(func(k string, v int) int { return v })
//     // pair.key == "go", pair.value == 14, score == 92, ok == true
func (c *Dictionary[K, V]) Max(predicate func(k K, v V) int) (Pair[K, V], int, bool) {
	if len(c.items) == 0 {
		var zero_key K
		var zero_val V
		return NewPair(zero_key, zero_val), 0, false
	}

	var (
		maxKey   K
		maxValue V
		maxScore int
		init     bool
	)

	for k, v := range c.items {
		score := predicate(k, v)

		if !init || score >= maxScore {
			maxKey = k
			maxValue = v
			maxScore = score
			init = true
		}
	}

	return NewPair(maxKey, maxValue), maxScore, true
}

// Min returns the key-value pair from the Dictionary that yields the minimum
// score when evaluated with the provided predicate function.
//
// The predicate function is applied to each key-value pair in the Dictionary.
// The pair that produces the smallest integer score is returned along with
// that score.
//
// Due to the unordered nature of maps, if multiple pairs produce the same
// minimum score, the returned pair is not deterministic.
//
// Parameters:
//   - predicate: A function that takes a key and a value, and returns an
//     integer score used for comparison.
//
// Returns:
//   - A Pair containing the key and value that produced the minimum score.
//   - The minimum score returned by the predicate.
//   - A boolean indicating whether the Dictionary was non-empty.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"go": 90, "rust": 85, "zig": 92})
//     pair, score, ok := dict.Min(func(k string, v int) int { return v })
//     // pair.key == "rust", pair.value == 85, score == 85, ok == true
func (c *Dictionary[K, V]) Min(
	predicate func(k K, v V) int,
) (Pair[K, V], int, bool) {
	if len(c.items) == 0 {
		var zero_key K
		var zero_val V
		return Pair[K, V]{key: zero_key, value: zero_val}, 0, false
	}

	var (
		minPair  Pair[K, V]
		minScore int
		init     bool
	)

	for k, v := range c.items {
		score := predicate(k, v)

		if !init || score <= minScore {
			minPair = Pair[K, V]{key: k, value: v}
			minScore = score
			init = true
		}
	}

	return minPair, minScore, true
}

// Keys returns a slice of all the keys in the Dictionary. The keys are returned in no specific order.
//
// Returns:
//   - A slice of type []K containing all the keys in the Dictionary.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     keys := dict.Keys() // keys will contain []string{"a", "b", "c"}
func (c Dictionary[K, V]) Keys() []K {
	keys := make([]K, 0, len(c.items))
	for key := range c.items {
		keys = append(keys, key)
	}
	return keys
}

// KeysVector returns a Vector containing all the keys in the Dictionary.
//
// Returns:
//   - A Vector[K] containing all the keys from the Dictionary.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     keysVector := dict.KeysVector() // keysVector will be a Vector containing ["a", "b", "c"]
func (c Dictionary[K, V]) KeysVector() *Vector[K] {
	return VectorFromList(c.Keys())
}

// Values returns a slice containing all the values in the Dictionary. The values are returned in no specific order.
//
// Returns:
//   - A slice of type []V containing all the values in the Dictionary.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     values := dict.Values() // values will contain []int{1, 2, 3}
func (c *Dictionary[K, V]) Values() []V {
	values := make([]V, 0, len(c.items))
	for key := range c.items {
		values = append(values, c.items[key])
	}
	return values
}

// ValuesVector returns a Vector containing all the values in the Dictionary.
//
// Returns:
//   - A Vector[V] containing all the values from the Dictionary.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     valuesVector := dict.ValuesVector() // valuesVector will be a Vector containing [1, 2, 3]
func (c Dictionary[K, V]) ValuesVector() *Vector[V] {
	return VectorFromList(c.Values())
}

// Pairs returns a slice of key-value pairs in the Dictionary, where each pair is represented as a Pair[K, V].
// The pairs are returned in no specific order.
//
// Returns:
//   - A slice of type []Pair[K, V] containing all key-value pairs from the Dictionary.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     pairs := dict.Pairs()
//     // pairs will contain [{a 1}, {b 2}, {c 3}], where each Pair holds a key-value pair
func (c *Dictionary[K, V]) Pairs() []Pair[K, V] {
	pairs := make([]Pair[K, V], 0, len(c.items))
	for k, v := range c.items {
		pairs = append(pairs, NewPair(k, v))
	}
	return pairs
}

// Collect returns an intance map containing all the key-value pairs in the Dictionary.
//
// Returns:
//   - A map of type map[K]V containing all key-value pairs in the Dictionary.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     collectedMap := dict.Collect() // collectedMap will be map[string]int{"a": 1, "b": 2}
func (c Dictionary[K, V]) Collect() map[K]V {
	return c.items
}

// DictionaryMap creates a new Dictionary by applying the provided predicate function to each key-value pair in the original IDictionary.
// The predicate function is applied to each key and value, and its result is used as the new value in the returned Dictionary.
//
// Parameters:
//   - c: A pointer to the Dictionary[K, V] from which the key-value pairs will be transformed.
//   - predicate: A function that takes a key of type K and a value of type V, and returns a new value of type E. This function is applied to each key-value pair.
//
// Returns:
//   - A new IDictionary[K, E] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//	newDict := DictionaryMap(dict, func(k string, v int) string { return fmt.Sprintf("%d", v) })
//	// newDict will contain {"a": "1", "b": "2"}, where the values are transformed to strings
func DictionaryMap[K comparable, V, E any](c IDictionary[K, V], predicate func(K, V) E) IDictionary[K, E] {
	return MapToDictionary(c.Collect(), predicate)
}

// MapToDictionary creates a new Dictionary by applying the provided predicate function to each key-value pair in the provided map.
// The predicate function is applied to each key and value, and its result is used as the new value in the returned Dictionary.
//
// Parameters:
//   - c: A map[K]V from which the key-value pairs will be transformed.
//   - predicate: A function that takes a key of type K and a value of type V, and returns a new value of type E. This function is applied to each key-value pair.
//
// Returns:
//   - A new IDictionary[K, E] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	dict := map[string]int{"a": 1, "b": 2}
//	newDict := MapToDictionary(dict, func(k string, v int) string { return fmt.Sprintf("%d", v) })
//	// newDict will contain {"a": "1", "b": "2"}, where the values are transformed to strings
func MapToDictionary[K comparable, V, E any](c map[K]V, predicate func(K, V) E) IDictionary[K, E] {
	return MapToIDictionary(c, predicate, MakeDictionary)
}

// VectorMapToDictionary applies the given predicate function to each element in the IVector,
// transforming each element of type K into an tuple of types E, that implements comparable, and V, then returns
// a new Dictionary with the transformed elements.
//
// Parameters:
//   - c: The source IVector containing elements of type K.
//   - predicate: A function that takes an element of type K and transforms it into an element of type V.
//
// Returns:
//   - A new IDictionary[E, V] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	vec := VectorFromList([]int{1, 2, 3, 4})
//	transformed := VectorMapToDictionary(vec, func(v int) (string, int) { return fmt.Sprintf("Item %d", v), v })
//	// transformed will be a new Vector with elements: {"Item 1": 1, "Item 2": 2, "Item 3": 3, "Item 4": 4}
func VectorMapToDictionary[K, V any, E comparable](c IVector[K], predicate func(K) (E, V)) IDictionary[E, V] {
	return ListMapToDictionary(c.Collect(), predicate)
}

// ListMapToDictionary applies the given predicate function to each element in the slice,
// transformng each element of type K into an tuple of types E, that implements comparable, and V, then returns
// a new Dictionary with the transformed elements.
//
// Parameters:
//   - c: The slice IVector containing elements of type K.
//   - predicate: A function that takes an element of type K and transforms it into an element of type V.
//
// Returns:
//   - A new IDictionary[E, V] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	slc := []int{1, 2, 3, 4}
//	transformed := ListMapToDictionary(slc, func(v int) (string, int) { return fmt.Sprintf("Item %d", v), v })
//	// transformed will be a new Vector with elements: {"Item 1": 1, "Item 2": 2, "Item 3": 3, "Item 4": 4}
func ListMapToDictionary[K, V any, E comparable](c []K, predicate func(K) (E, V)) IDictionary[E, V] {
	return ListMapToIDictionary(c, predicate, MakeDictionary)
}
