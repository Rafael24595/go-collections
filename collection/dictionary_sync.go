package collection

import (
	"maps"
	"sync"
)

// DictionarySync is a thread-safe generic key-value store where each key is of type K and each value is of type V.
// The DictionarySync provides methods to manipulate and interact with key-value pairs while ensuring safe concurrent access.
//
// Thread Safety:
//   - A read-write mutex (sync.RWMutex) is used to protect access to the underlying map (`items`).
//   - Read operations (e.g., Get, Size) use a read lock (`RLock()`), allowing concurrent reads.
//   - Write operations (e.g., Put, Delete) use a write lock (`Lock()`) to ensure exclusive access.
//
// Thread Safety:
//   - Allows multiple goroutines to safely read and modify the dictionary.
//   - Internally manages synchronization to prevent race conditions.
//
// Fields:
//   - items: A map storing the actual key-value pairs. The keys are of type K, and the values are of type V.
//
// Example usage:
//
//	dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//	dict.Put("c", 3)
//	value, exists := dict.Get("a") // value will be 1, exists will be true
type DictionarySync[K comparable, V any] struct {
	mu    sync.RWMutex
	items map[K]V
}

// MakeDictionarySync creates a new DictionarySync from a given map.
// It takes a map with keys of type K and values of type V and
// returns a pointer to a IDictionary containing the same items.
//
// K must be a comparable type to be used as a map key.
// V can be any type.
//
// Example usage:
//
//	myMap := map[string]int{"a": 1, "b": 2}
//	dict := MakeDictionarySync(myMap)
func MakeDictionarySync[K comparable, V any](items map[K]V) IDictionary[K, V] {
	return DictionarySyncFromMap(items)
}

// DictionarySyncFromMap creates a new DictionarySync from a given map.
// It takes a map with keys of type K and values of type V and
// returns a pointer to a DictionarySync containing the same items.
//
// K must be a comparable type to be used as a map key.
// V can be any type.
//
// Example usage:
//
//	myMap := map[string]int{"a": 1, "b": 2}
//	dict := DictionarySyncFromMap(myMap)
func DictionarySyncFromMap[K comparable, V any](items map[K]V) *DictionarySync[K, V] {
	return &DictionarySync[K, V]{
		items: items,
	}
}

// DictionarySyncEmpty creates and returns a new, empty DictionarySync.
//
// K must be a comparable type to be used as a map key.
// V can be any type.
//
// Example usage:
//
//	emptyDict := DictionarySyncEmpty[string, int]()
func DictionarySyncEmpty[K comparable, V any]() *DictionarySync[K, V] {
	return DictionarySyncFromMap(make(map[K]V))
}

// DictionarySyncFromVector creates a DictionarySync from a Vector by applying a mapping function.
//
// K must be a comparable type to be used as a dictionary key.
// V can be any type.
//
// Parameters:
//   - vector: A Vector containing values of type V.
//   - mapper: A function that converts a value of type V into a key of type K.
//
// Returns:
//   - A pointer to a DictionarySync[K, V].
//
// Example usage:
//
//	vec := VectorFromList([]int{10, 20, 30})
//	dict := DictionaryFromVector(vec, func(i int) string { return fmt.Sprintf("key-%d", i) })
func DictionarySyncFromVector[K comparable, V any](vector Vector[V], mapper func(V) K) *DictionarySync[K, V] {
	return DictionarySyncFromList(vector.items, mapper)
}

// DictionarySyncFromList creates a DictionarySync from a slice by applying a mapping function.
//
// K must be a comparable type to be used as a dictionary key.
// V can be any type.
//
// Parameters:
//   - vector: A slice of values of type V.
//   - mapper: A function that converts a value of type V into a key of type K.
//
// Returns:
//   - A pointer to a DictionarySync[K, V] containing the mapped key-value pairs.
//
// Example usage:
//
//	list := []int{10, 20, 30}
//	dict := DictionarySyncFromList(list, func(i int) string { return fmt.Sprintf("key-%d", i)
func DictionarySyncFromList[K comparable, V any](vector []V, mapper func(V) K) *DictionarySync[K, V] {
	mapp := DictionarySyncEmpty[K, V]()
	for _, v := range vector {
		mapp.Put(mapper(v), v)
	}
	return mapp
}

// Size returns the number of key-value pairs in the DictionarySync.
//
// Returns:
//   - An integer representing the number of elements in the DictionarySync.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	size := dict.Size() // size will be 2
func (c *DictionarySync[K, V]) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.items)
}

// Exists checks if the given key exists in the DictionarySync.
//
// Parameters:
//   - key: The key of type K to check for in the DictionarySync.
//
// Returns:
//   - A boolean indicating whether the key exists in the DictionarySync.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	exists := dict.Exists("a") // exists will be true
//	exists = dict.Exists("c")  // exists will be false
func (c *DictionarySync[K, V]) Exists(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, exists := c.items[key]
	return exists
}

// Find returns a slice of values from the DictionarySync that satisfy the given predicate function.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and returns a boolean.
//     The function should return true for the values that should be included in the result.
//
// Returns:
//   - A slice of values of type V that satisfy the predicate function.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//	result := dict.Find(func(k string, v int) bool { return v > 1 })
//	// result will be [2, 3]
func (c *DictionarySync[K, V]) Find(predicate func(K, V) bool) []V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	filter := []V{}
	for k, v := range c.items {
		if predicate(k, v) {
			filter = append(filter, v)
		}
	}
	return filter
}

// FindOne searches for the first key-value pair in the DictionarySync that satisfies the given predicate function.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and returns a boolean.
//     The function should return true for the first pair that matches the search criteria.
//
// Returns:
//   - A pointer to the value of type V if a matching key-value pair is found, or nil if not found.
//   - A boolean indicating whether a match was found (true if found, false otherwise).
//
// Example usage:
//
//	dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//	value, found := dict.FindOne(func(k string, v int) bool { return v == 2 })
//	// value will be a pointer to 2, found will be true
//	value, found = dict.FindOne(func(k string, v int) bool { return v == 4 })
//	// value will be nil, found will be false
func (c *DictionarySync[K, V]) FindOne(predicate func(K, V) bool) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for k, v := range c.items {
		if predicate(k, v) {
			return v, true
		}
	}
	var zero V
	return zero, false
}

// Get retrieves the value associated with the given key in the DictionarySync.
// It returns a pointer to the value if the key exists, and a boolean indicating whether the key was found.
//
// Parameters:
//   - key: The key of type K whose associated value is to be retrieved.
//
// Returns:
//   - A pointer to the value of type V associated with the key, or nil if the key does not exist.
//   - A boolean indicating whether the key was found in the DictionarySync (true if found, false otherwise).
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	value, found := dict.Get("a") // value will be a pointer to 1, found will be true
//	value, found = dict.Get("c")  // value will be nil, found will be false
func (c *DictionarySync[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, exists := c.items[key]
	return value, exists
}

// Put adds a key-value pair to the DictionarySync, updating the value if the key already exists.
// It returns the old value associated with the key, if any, and a boolean indicating whether
// the key already existed in the DictionarySync (true if it existed, false otherwise).
//
// Parameters:
//   - key: The key of type K to associate with the given value.
//   - item: The value of type V to be associated with the key.
//
// Returns:
//   - A pointer to the old value associated with the key, or nil if the key did not exist.
//   - A boolean indicating whether the key was already present in the DictionarySync (true if it existed).
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	oldValue, exists := dict.Put("a", 3) // oldValue will be a pointer to 1, exists will be true
//	oldValue, exists = dict.Put("c", 4)  // oldValue will be nil, exists will be false
func (c *DictionarySync[K, V]) Put(key K, item V) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	old, exists := c.items[key]
	c.items[key] = item
	return old, exists
}

// PutIfAbsent adds a key-value pair to the DictionarySync only if the key does not already exist.
// If the key is already present, it does nothing and returns the existing value associated with the key,
// along with a boolean indicating that the key was already present.
//
// Parameters:
//   - key: The key of type K to associate with the given value.
//   - item: The value of type V to be associated with the key if the key is absent.
//
// Returns:
//   - A pointer to the old value associated with the key, or nil if the key was not found.
//   - A boolean indicating whether the key was already present in the DictionarySync (true if it existed, false if it was absent).
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	oldValue, exists := dict.PutIfAbsent("a", 3) // oldValue will be a pointer to 1, exists will be true
//	oldValue, exists = dict.PutIfAbsent("c", 4)  // oldValue will be nil, exists will be false
func (c *DictionarySync[K, V]) PutIfAbsent(key K, item V) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	old, exists := c.items[key]
	if !exists {
		c.items[key] = item
	}
	return old, exists
}

// PutAll adds all key-value pairs from another map to the DictionarySync
// overwriting any existing values for the keys that already exist in the DictionarySync.
//
// Parameters:
//   - items: A map of type map[K]V containing the key-value pairs to add to the DictionarySync.
//
// Returns:
//   - The DictionarySync itself, with all the new key-value pairs added.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	otherMap := map[string]int{"b": 3, "c": 4}
//	dict.PutAll(otherMap) // dict will contain {"a": 1, "b": 3, "c": 4}
func (c *DictionarySync[K, V]) PutAll(items map[K]V) IDictionary[K, V] {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key := range items {
		c.items[key] = items[key]
	}
	return c
}

// Merge combines all key-value pairs from another DictionarySync into the current DictionarySync
// overwriting any existing values for the keys that already exist.
//
// Parameters:
//   - other: A DictionarySync of type DictionarySync[K, V] to merge into the current DictionarySync.
//
// Returns:
//   - The Dictionary itself, with the key-value pairs from the other DictionarySync added.
//
// Example usage:
//
//	dict1 := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	dict2 := DictionarySyncFromMap(map[string]int{"b": 3, "c": 4})
//	dict1.Merge(dict2) // dict1 will contain {"a": 1, "b": 3, "c": 4}
func (c *DictionarySync[K, V]) Merge(other IDictionary[K, V]) IDictionary[K, V] {
	return c.PutAll(other.Collect())
}

// Filter creates a new DictionarySync by filtering the key-value pairs in the current DictionarySync
// based on the provided predicate function. It iterates over all key-value pairs and retains
// those that satisfy the condition defined in the predicate.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and returns a boolean.
//     The function should return true for the key-value pairs that should be kept in the result.
//
// Returns:
//   - A new DictionarySync itself, after filtering out the key-value pairs that do not satisfy the predicate.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//	filtered := dict.Filter(func(k string, v int) bool { return v > 1 })
//	// filtered will contain {"b": 2, "c": 3}
func (c *DictionarySync[K, V]) Filter(predicate func(K, V) bool) IDictionary[K, V] {
	c.mu.RLock()

	filter := map[K]V{}
	for key, v := range c.items {
		if predicate(key, v) {
			filter[key] = v
		}
	}

	c.mu.RUnlock()

	return DictionarySyncFromMap(filter)
}

// FilterSelf filters the key-value pairs in the current DictionarySync based on the provided predicate function.
// It updates the DictionarySync itself, removing key-value pairs that do not satisfy the condition defined in the predicate.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and returns a boolean.
//     The function should return true for the key-value pairs that should be retained in the Dictionary.
//
// Returns:
//   - The DictionarySync itself, with only the key-value pairs that satisfy the predicate.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//	dict.FilterSelf(func(k string, v int) bool { return v > 1 })
//	// dict will contain {"b": 2, "c": 3}
func (c *DictionarySync[K, V]) FilterSelf(predicate func(K, V) bool) IDictionary[K, V] {
	c.mu.Lock()
	defer c.mu.Unlock()

	filter := map[K]V{}
	for key, v := range c.items {
		if predicate(key, v) {
			filter[key] = v
		}
	}
	c.items = filter
	return c
}

// Remove deletes a key-value pair from the DictionarySync by the provided key.
// It returns the old value associated with the key, if it exists, along with a boolean
// indicating whether the key was found and removed from the DictionarySync.
//
// Parameters:
//   - key: The key of type K to remove from the Dictionary.
//
// Returns:
//   - A pointer to the old value associated with the key, or nil if the key was not found.
//   - A boolean indicating whether the key was present and removed (true if removed, false if not).
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	oldValue, exists := dict.Remove("a", 1) // oldValue will be a pointer to 1, exists will be true
//	oldValue, exists = dict.Remove("c", 3)  // oldValue will be nil, exists will be false
func (c *DictionarySync[K, V]) Remove(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	old, exists := c.items[key]
	delete(c.items, key)
	return old, exists
}

// ForEach iterates over all key-value pairs in the DictionarySync, applying the provided predicate function to each pair.
// The predicate is called with each key and value, allowing side effects or custom actions for every entry in the DictionarySync.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and performs an action or operation.
//
// Returns:
//   - The DictionarySync itself, allowing for method chaining.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	dict.ForEach(func(k string, v int) { fmt.Println(k, v) })
//	// Output:
//	// a 1
//	// b 2
func (c *DictionarySync[K, V]) ForEach(predicate func(K, V)) IDictionary[K, V] {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for k, v := range c.items {
		predicate(k, v)
	}
	return c
}

// Map transforms the values in the DictionarySync by applying the provided predicate function to each key-value pair.
//
// Parameters:
//   - predicate: A function that takes a key of type K and a value of type V, and returns a new value of type V.
//
// Returns:
//   - The DictionarySync itself, with the transformed values.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	dict.Map(func(k string, v int) int { return v * 2 })
//	// dict will contain {"a": 2, "b": 4}
func (c *DictionarySync[K, V]) Map(predicate func(K, V) V) IDictionary[K, V] {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.items {
		c.items[k] = predicate(k, v)
	}
	return c
}

// Clean removes all key-value pairs from the DictionarySync, effectively clearing its contents.
//
// Returns:
//   - The DictionarySync itself, now empty, allowing for method chaining.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	dict.Clean() // dict will be empty: {}
func (c *DictionarySync[K, V]) Clean() IDictionary[K, V] {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[K]V)
	return c
}

// Clone creates a shallow copy of the DictionarySync, including all key-value pairs.
// The new DictionarySync will have the same keys and values as the original, but modifications to one
// will not affect the other.
//
// Returns:
//   - A new DictionarySync that is a clone of the current DictionarySync.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	clonedDict := dict.Clone() // clonedDict is a new DictionarySync with the same contents as dict
func (c *DictionarySync[K, V]) Clone() IDictionary[K, V] {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cloned := make(map[K]V)
	maps.Copy(cloned, c.items)
	return DictionarySyncFromMap(cloned)
}

// Max returns the key-value pair from the DictionarySync that yields the maximum
// score when evaluated with the provided predicate function.
//
// The predicate function is applied to each key-value pair in the DictionarySync.
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
//   - A boolean indicating whether the DictionarySync was non-empty.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"go": 14, "rust": 11, "zig": 3})
//	pair, score, ok := dict.Max(func(k string, v int) int { return v })
//	// pair.key == "go", pair.value == 14, score == 92, ok == true
func (c *DictionarySync[K, V]) Max(predicate func(k K, v V) int) (Pair[K, V], int, bool) {
	if len(c.items) == 0 {
		var zero_key K
		var zero_val V
		return NewPair(zero_key, zero_val), 0, false
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

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

// Min returns the key-value pair from the DictionarySync that yields the minimum
// score when evaluated with the provided predicate function.
//
// The predicate function is applied to each key-value pair in the DictionarySync.
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
//   - A boolean indicating whether the DictionarySync was non-empty.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"go": 90, "rust": 85, "zig": 92})
//	pair, score, ok := dict.Min(func(k string, v int) int { return v })
//	// pair.key == "rust", pair.value == 85, score == 85, ok == true
func (c *DictionarySync[K, V]) Min(
	predicate func(k K, v V) int,
) (Pair[K, V], int, bool) {
	if len(c.items) == 0 {
		var zero_key K
		var zero_val V
		return Pair[K, V]{key: zero_key, value: zero_val}, 0, false
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

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

// Keys returns a slice of all the keys in the DictionarySync. The keys are returned in no specific order.
//
// Returns:
//   - A slice of type []K containing all the keys in the DictionarySync.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//	keys := dict.Keys() // keys will contain []string{"a", "b", "c"}
func (c *DictionarySync[K, V]) Keys() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := make([]K, 0, len(c.items))
	for key := range c.items {
		keys = append(keys, key)
	}
	return keys
}

// KeysVector returns a Vector containing all the keys in the DictionarySync.
//
// Returns:
//   - A Vector[K] containing all the keys from the DictionarySync.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//	keysVector := dict.KeysVector() // keysVector will be a Vector containing ["a", "b", "c"]
func (c *DictionarySync[K, V]) KeysVector() *Vector[K] {
	return VectorFromList(c.Keys())
}

// Values returns a slice containing all the values in the DictionarySync. The values are returned in no specific order.
//
// Returns:
//   - A slice of type []V containing all the values in the DictionarySync.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//	values := dict.Values() // values will contain []int{1, 2, 3}
func (c *DictionarySync[K, V]) Values() []V {
	c.mu.RLock()
	defer c.mu.RUnlock()

	values := make([]V, 0, len(c.items))
	for key := range c.items {
		values = append(values, c.items[key])
	}
	return values
}

// ValuesVector returns a Vector containing all the values in the DictionarySync.
//
// Returns:
//   - A Vector[V] containing all the values from the Dictionary.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//	valuesVector := dict.ValuesVector() // valuesVector will be a Vector containing [1, 2, 3]
func (c *DictionarySync[K, V]) ValuesVector() *Vector[V] {
	return VectorFromList(c.Values())
}

// Pairs returns a slice of key-value pairs in the DictionarySync, where each pair is represented as a Pair[K, V].
// The pairs are returned in no specific order.
//
// Returns:
//   - A slice of type []Pair[K, V] containing all key-value pairs from the DictionarySync.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//	pairs := dict.Pairs()
//	// pairs will contain [{a 1}, {b 2}, {c 3}], where each Pair holds a key-value pair
func (c *DictionarySync[K, V]) Pairs() []Pair[K, V] {
	c.mu.RLock()
	defer c.mu.RUnlock()

	pairs := make([]Pair[K, V], 0, len(c.items))
	for k, v := range c.items {
		pairs = append(pairs, NewPair(k, v))
	}
	return pairs
}

// Collect returns an instance of map containing all the key-value pairs in the DictionarySync.
//
// Returns:
//   - A map of type map[K]V containing all key-value pairs in the DictionarySync.
//
// Example usage:
//
//	dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//	collectedMap := dict.Collect() // collectedMap will be map[string]int{"a": 1, "b": 2}
func (c *DictionarySync[K, V]) Collect() map[K]V {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return maps.Clone(c.items)
}

// DictionarySyncMap creates a new DictionarySync by applying the provided predicate function to each key-value pair in the original IDictionary.
// The predicate function is applied to each key and value, and its result is used as the new value in the returned DictionarySync.
//
// Parameters:
//   - c: A pointer to the IDictionary[K, V] from which the key-value pairs will be transformed.
//   - predicate: A function that takes a key of type K and a value of type V, and returns a new value of type E. This function is applied to each key-value pair.
//
// Returns:
//   - A new Dictionary[K, E] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//	newDict := DictionaryMap(dict, func(k string, v int) string { return fmt.Sprintf("%d", v) })
//	// newDict will contain {"a": "1", "b": "2"}, where the values are transformed to strings
func DictionarySyncMap[K comparable, V, E any](c IDictionary[K, V], predicate func(K, V) E) IDictionary[K, E] {
	return MapToDictionarySync(c.Collect(), predicate)
}

// MapToDictionary creates a new DictionarySync by applying the provided predicate function to each key-value pair in the provided map.
// The predicate function is applied to each key and value, and its result is used as the new value in the returned DictionarySync.
//
// Parameters:
//   - c: A map[K]V from which the key-value pairs will be transformed.
//   - predicate: A function that takes a key of type K and a value of type V, and returns a new value of type E. This function is applied to each key-value pair.
//
// Returns:
//   - A new Dictionary[K, E] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//
//	dict := map[string]int{"a": 1, "b": 2}
//	newDict := MapToDictionary(dict, func(k string, v int) string { return fmt.Sprintf("%d", v) })
//	// newDict will contain {"a": "1", "b": "2"}, where the values are transformed to strings
func MapToDictionarySync[K comparable, V, E any](c map[K]V, predicate func(K, V) E) IDictionary[K, E] {
	return MapToIDictionary(c, predicate, MakeDictionarySync)
}

// VectorMapToDictionarySync applies the given predicate function to each element in the IVector,
// transforming each element of type K into an tuple of types E, that implements comparable, and V, then returns
// a new DictionarySync with the transformed elements.
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
//	transformed := VectorMapToDictionarySync(vec, func(v int) (string, int) { return fmt.Sprintf("Item %d", v), v })
//	// transformed will be a new Vector with elements: {"Item 1": 1, "Item 2": 2, "Item 3": 3, "Item 4": 4}
func VectorMapToDictionarySync[K, V any, E comparable](c IVector[K], predicate func(K) (E, V)) IDictionary[E, V] {
	return ListMapToDictionarySync(c.Collect(), predicate)
}

// ListMapToDictionarySync applies the given predicate function to each element in the slice,
// transformng each element of type K into an tuple of types E, that implements comparable, and V, then returns
// a new DictionarySync with the transformed elements.
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
//	transformed := ListMapToDictionarySync(slc, func(v int) (string, int) { return fmt.Sprintf("Item %d", v), v })
//	// transformed will be a new Vector with elements: {"Item 1": 1, "Item 2": 2, "Item 3": 3, "Item 4": 4}
func ListMapToDictionarySync[K, V any, E comparable](c []K, predicate func(K) (E, V)) IDictionary[E, V] {
	return ListMapToIDictionary(c, predicate, MakeDictionarySync)
}
