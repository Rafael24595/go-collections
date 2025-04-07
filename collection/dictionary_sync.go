package collection

import (
	"sync"
	"maps"
)

// DictionarySync is a thread-safe generic key-value store where each key is of type T and each value is of type K.
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
//   - items: A map storing the actual key-value pairs. The keys are of type T, and the values are of type K.
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2})
//     dict.Put("c", 3)
//     value, exists := dict.Get("a") // value will be 1, exists will be true
type DictionarySync[T comparable, K any] struct {
	mu sync.RWMutex
	items map[T]K
}

// DictionarySyncFromMap creates a new DictionarySync from a given map.
// It takes a map with keys of type T and values of type K and 
// returns a pointer to a DictionarySync containing the same items.
//
// T must be a comparable type to be used as a map key.
// K can be any type.
//
// Example usage:
//     myMap := map[string]int{"a": 1, "b": 2}
//     dict := DictionarySyncFromMap(myMap)
func DictionarySyncFromMap[T comparable, K any](items map[T]K) *DictionarySync[T, K] {
	return &DictionarySync[T, K]{
		items: items,
	}
}

// DictionarySyncEmpty creates and returns a new, empty DictionarySync.
//
// T must be a comparable type to be used as a map key.
// K can be any type.
//
// Example usage:
//     emptyDict := DictionarySyncEmpty[string, int]()
func DictionarySyncEmpty[T comparable, K any]() *DictionarySync[T, K] {
	return DictionarySyncFromMap(make(map[T]K))
}

// DictionarySyncFromVector creates a DictionarySync from a Vector by applying a mapping function.
//
// T must be a comparable type to be used as a dictionary key.
// K can be any type.
//
// Parameters:
//   - vector: A Vector containing values of type K.
//   - mapper: A function that converts a value of type K into a key of type T.
//
// Returns:
//   - A pointer to a DictionarySync[T, K].
//
// Example usage:
//     vec := VectorFromList([]int{10, 20, 30})
//     dict := DictionaryFromVector(vec, func(i int) string { return fmt.Sprintf("key-%d", i) })
func DictionarySyncFromVector[T comparable, K any](vector Vector[K], mapper func(K) T) *DictionarySync[T, K] {
	return DictionarySyncFromList(vector.items, mapper)
}

// DictionarySyncFromList creates a DictionarySync from a slice by applying a mapping function.
//
// T must be a comparable type to be used as a dictionary key.
// K can be any type.
//
// Parameters:
//   - vector: A slice of values of type K.
//   - mapper: A function that converts a value of type K into a key of type T.
//
// Returns:
//   - A pointer to a DictionarySync[T, K] containing the mapped key-value pairs.
//
// Example usage:
//     list := []int{10, 20, 30}
//     dict := DictionarySyncFromList(list, func(i int) string { return fmt.Sprintf("key-%d", i) 
func DictionarySyncFromList[T comparable, K any](vector []K, mapper func(K) T) *DictionarySync[T, K] {
	mapp := DictionarySyncEmpty[T, K]()
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
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     size := dict.Size() // size will be 2
func (c *DictionarySync[T, K]) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.items)
}

// Exists checks if the given key exists in the DictionarySync.
//
// Parameters:
//   - key: The key of type T to check for in the DictionarySync.
//
// Returns:
//   - A boolean indicating whether the key exists in the DictionarySync.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     exists := dict.Exists("a") // exists will be true
//     exists = dict.Exists("c")  // exists will be false
func (c *DictionarySync[T, K]) Exists(key T) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, exists := c.items[key]
	return exists
}

// Find returns a slice of values from the DictionarySync that satisfy the given predicate function.
//
// Parameters:
//   - predicate: A function that takes a key of type T and a value of type K, and returns a boolean.
//                The function should return true for the values that should be included in the result.
//
// Returns:
//   - A slice of values of type K that satisfy the predicate function.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     result := dict.Find(func(k string, v int) bool { return v > 1 })
//     // result will be [2, 3]
func (c *DictionarySync[T, K]) Find(predicate func(T, K) bool) []K {
	c.mu.RLock()
	defer c.mu.RUnlock()

	filter := []K{}
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
//   - predicate: A function that takes a key of type T and a value of type K, and returns a boolean.
//                The function should return true for the first pair that matches the search criteria.
//
// Returns:
//   - A pointer to the value of type K if a matching key-value pair is found, or nil if not found.
//   - A boolean indicating whether a match was found (true if found, false otherwise).
//
// Example usage:
//     dict := DictionaryFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     value, found := dict.FindOne(func(k string, v int) bool { return v == 2 })
//     // value will be a pointer to 2, found will be true
//     value, found = dict.FindOne(func(k string, v int) bool { return v == 4 })
//     // value will be nil, found will be false
func (c *DictionarySync[T, K]) FindOne(predicate func(T, K) bool) (*K, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for k, v := range c.items {
		if predicate(k, v) {
			return &v, true
		}
	}
	return nil, false
}

// Get retrieves the value associated with the given key in the DictionarySync.
// It returns a pointer to the value if the key exists, and a boolean indicating whether the key was found.
//
// Parameters:
//   - key: The key of type T whose associated value is to be retrieved.
//
// Returns:
//   - A pointer to the value of type K associated with the key, or nil if the key does not exist.
//   - A boolean indicating whether the key was found in the DictionarySync (true if found, false otherwise).
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     value, found := dict.Get("a") // value will be a pointer to 1, found will be true
//     value, found = dict.Get("c")  // value will be nil, found will be false
func (c *DictionarySync[T, K]) Get(key T) (*K, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, exists := c.items[key]
	return &value, exists
}

// Put adds a key-value pair to the DictionarySync, updating the value if the key already exists.
// It returns the old value associated with the key, if any, and a boolean indicating whether
// the key already existed in the DictionarySync (true if it existed, false otherwise).
//
// Parameters:
//   - key: The key of type T to associate with the given value.
//   - item: The value of type K to be associated with the key.
//
// Returns:
//   - A pointer to the old value associated with the key, or nil if the key did not exist.
//   - A boolean indicating whether the key was already present in the DictionarySync (true if it existed).
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     oldValue, exists := dict.Put("a", 3) // oldValue will be a pointer to 1, exists will be true
//     oldValue, exists = dict.Put("c", 4)  // oldValue will be nil, exists will be false
func (c *DictionarySync[T, K]) Put(key T, item K) (*K, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	old, exists := c.items[key]
	c.items[key] = item
	return &old, exists
}

// PutIfAbsent adds a key-value pair to the DictionarySync only if the key does not already exist.
// If the key is already present, it does nothing and returns the existing value associated with the key,
// along with a boolean indicating that the key was already present.
//
// Parameters:
//   - key: The key of type T to associate with the given value.
//   - item: The value of type K to be associated with the key if the key is absent.
//
// Returns:
//   - A pointer to the old value associated with the key, or nil if the key was not found.
//   - A boolean indicating whether the key was already present in the DictionarySync (true if it existed, false if it was absent).
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     oldValue, exists := dict.PutIfAbsent("a", 3) // oldValue will be a pointer to 1, exists will be true
//     oldValue, exists = dict.PutIfAbsent("c", 4)  // oldValue will be nil, exists will be false
func (c *DictionarySync[T, K]) PutIfAbsent(key T, item K) (*K, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	old, exists :=  c.items[key]
	if !exists {
		c.items[key] = item
	}
	return &old, exists
}

// PutAll adds all key-value pairs from another map to the DictionarySync
// overwriting any existing values for the keys that already exist in the DictionarySync.
//
// Parameters:
//   - items: A map of type map[T]K containing the key-value pairs to add to the DictionarySync.
//
// Returns:
//   - The DictionarySync itself, with all the new key-value pairs added.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     otherMap := map[string]int{"b": 3, "c": 4}
//     dict.PutAll(otherMap) // dict will contain {"a": 1, "b": 3, "c": 4}
func (c *DictionarySync[T, K]) PutAll(items map[T]K) IDictionary[T, K] {
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
//   - other: A DictionarySync of type DictionarySync[T, K] to merge into the current DictionarySync.
//
// Returns:
//   - The Dictionary itself, with the key-value pairs from the other DictionarySync added.
//
// Example usage:
//     dict1 := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     dict2 := DictionarySyncFromMap(map[string]int{"b": 3, "c": 4})
//     dict1.Merge(dict2) // dict1 will contain {"a": 1, "b": 3, "c": 4}
func (c *DictionarySync[T, K]) Merge(other IDictionary[T, K]) IDictionary[T, K] {
	return c.PutAll(other.Collect())
}

// Filter creates a new DictionarySync by filtering the key-value pairs in the current DictionarySync
// based on the provided predicate function. It iterates over all key-value pairs and retains
// those that satisfy the condition defined in the predicate.
//
// Parameters:
//   - predicate: A function that takes a key of type T and a value of type K, and returns a boolean.
//                The function should return true for the key-value pairs that should be kept in the result.
//
// Returns:
//   - A new DictionarySync itself, after filtering out the key-value pairs that do not satisfy the predicate.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     filtered := dict.Filter(func(k string, v int) bool { return v > 1 })
//     // filtered will contain {"b": 2, "c": 3}
func (c *DictionarySync[T, K]) Filter(predicate func(T, K) bool) IDictionary[T, K] {
	c.mu.RLock()

	filter := map[T]K{}
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
//   - predicate: A function that takes a key of type T and a value of type K, and returns a boolean.
//                The function should return true for the key-value pairs that should be retained in the Dictionary.
//
// Returns:
//   - The DictionarySync itself, with only the key-value pairs that satisfy the predicate.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     dict.FilterSelf(func(k string, v int) bool { return v > 1 })
//     // dict will contain {"b": 2, "c": 3}
func (c *DictionarySync[T, K]) FilterSelf(predicate func(T, K) bool) IDictionary[T, K] {
	c.mu.Lock()
	defer c.mu.Unlock()

	filter := map[T]K{}
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
//   - key: The key of type T to remove from the Dictionary.
//
// Returns:
//   - A pointer to the old value associated with the key, or nil if the key was not found.
//   - A boolean indicating whether the key was present and removed (true if removed, false if not).
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     oldValue, exists := dict.Remove("a", 1) // oldValue will be a pointer to 1, exists will be true
//     oldValue, exists = dict.Remove("c", 3)  // oldValue will be nil, exists will be false
func (c *DictionarySync[T, K]) Remove(key T) (*K, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	old, exists := c.items[key]
	delete(c.items, key)
	return &old, exists
}

// ForEach iterates over all key-value pairs in the DictionarySync, applying the provided predicate function to each pair.
// The predicate is called with each key and value, allowing side effects or custom actions for every entry in the DictionarySync.
//
// Parameters:
//   - predicate: A function that takes a key of type T and a value of type K, and performs an action or operation.
//
// Returns:
//   - The DictionarySync itself, allowing for method chaining.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     dict.ForEach(func(k string, v int) { fmt.Println(k, v) })
//     // Output:
//     // a 1
//     // b 2
func (c *DictionarySync[T, K]) ForEach(predicate func(T, K)) IDictionary[T, K] {
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
//   - predicate: A function that takes a key of type T and a value of type K, and returns a new value of type K.
//
// Returns:
//   - The DictionarySync itself, with the transformed values.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     dict.Map(func(k string, v int) int { return v * 2 })
//     // dict will contain {"a": 2, "b": 4}
func (c *DictionarySync[T, K]) Map(predicate func(T, K) K) IDictionary[T, K] {
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
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     dict.Clean() // dict will be empty: {}
func (c *DictionarySync[T, K]) Clean() IDictionary[T, K] {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[T]K)
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
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     clonedDict := dict.Clone() // clonedDict is a new DictionarySync with the same contents as dict
func (c *DictionarySync[T, K]) Clone() IDictionary[T, K] {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cloned := make(map[T]K)
	maps.Copy(cloned, c.items)
	return DictionarySyncFromMap(cloned)
}

// Keys returns a slice of all the keys in the DictionarySync. The keys are returned in no specific order.
//
// Returns:
//   - A slice of type []T containing all the keys in the DictionarySync.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     keys := dict.Keys() // keys will contain []string{"a", "b", "c"}
func (c *DictionarySync[T, K]) Keys() []T {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	keys := make([]T, 0, len(c.items))
	for key := range c.items {
		keys = append(keys, key)
	}
	return keys
}

// KeysVector returns a Vector containing all the keys in the DictionarySync.
//
// Returns:
//   - A Vector[T] containing all the keys from the DictionarySync.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     keysVector := dict.KeysVector() // keysVector will be a Vector containing ["a", "b", "c"]
func (c *DictionarySync[T, K]) KeysVector() *Vector[T] {
	return VectorFromList(c.Keys())
}

// Values returns a slice containing all the values in the DictionarySync. The values are returned in no specific order.
//
// Returns:
//   - A slice of type []K containing all the values in the DictionarySync.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     values := dict.Values() // values will contain []int{1, 2, 3}
func (c *DictionarySync[T, K]) Values() []K {
	c.mu.RLock()
	defer c.mu.RUnlock()

	values := make([]K, 0, len(c.items))
	for key := range c.items {
		values = append(values, c.items[key])
	}
	return values
}

// ValuesVector returns a Vector containing all the values in the DictionarySync.
//
// Returns:
//   - A Vector[K] containing all the values from the Dictionary.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     valuesVector := dict.ValuesVector() // valuesVector will be a Vector containing [1, 2, 3]
func (c *DictionarySync[T, K]) ValuesVector() *Vector[K] {
	return VectorFromList(c.Values())
}

// Pairs returns a slice of key-value pairs in the DictionarySync, where each pair is represented as a Pair[T, K].
// The pairs are returned in no specific order.
//
// Returns:
//   - A slice of type []Pair[T, K] containing all key-value pairs from the DictionarySync.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2, "c": 3})
//     pairs := dict.Pairs() 
//     // pairs will contain [{a 1}, {b 2}, {c 3}], where each Pair holds a key-value pair
func (c *DictionarySync[T, K]) Pairs() []Pair[T, K] {
	c.mu.RLock()
	defer c.mu.RUnlock()

	pairs := make([]Pair[T, K], 0, len(c.items))
	for k, v := range c.items {
		pairs = append(pairs, NewPair(k, v))
	}
	return pairs
}

// Collect returns an instance of map containing all the key-value pairs in the DictionarySync.
//
// Returns:
//   - A map of type map[T]K containing all key-value pairs in the DictionarySync.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     collectedMap := dict.Collect() // collectedMap will be map[string]int{"a": 1, "b": 2}
func (c *DictionarySync[T, K]) Collect() map[T]K {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return maps.Clone(c.items)
}

// DictionarySyncMap creates a new DictionarySync by applying the provided predicate function to each key-value pair in the original DictionarySync.
// The predicate function is applied to each key and value, and its result is used as the new value in the returned DictionarySync.
//
// Parameters:
//   - c: A pointer to the DictionarySync[T, K] from which the key-value pairs will be transformed.
//   - predicate: A function that takes a key of type T and a value of type K, and returns a new value of type E. This function is applied to each key-value pair.
//
// Returns:
//   - A new DictionarySync[T, E] where the keys remain the same, but the values are the result of applying the predicate function.
//
// Example usage:
//     dict := DictionarySyncFromMap(map[string]int{"a": 1, "b": 2})
//     newDict := DictionarySyncMap(dict, func(k string, v int) string { return fmt.Sprintf("%d", v) })
//     // newDict will contain {"a": "1", "b": "2"}, where the values are transformed to strings
func DictionarySyncMap[T comparable, K, E any](c *DictionarySync[T, K], predicate func(T, K) E) *DictionarySync[T, E] {
	mapped := map[T]E{}
	for key, item := range c.Collect() {
		mapped[key] = predicate(key, item)
	}
	return &DictionarySync[T, E]{
		items: mapped,
	}
}
