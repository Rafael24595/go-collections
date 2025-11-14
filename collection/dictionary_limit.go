package collection

// DictionaryLimit is a thread-safe key-value store with a fixed size limit. It extends DictionarySync
// and maintains a timeline of inserted keys to track the order of entries.
//
// When the size limit is reached, older entries may be removed based on insertion order.
//
// Type Parameters:
//   - T: The type of the keys, which must be comparable.
//   - K: The type of the values.
//
// Fields:
//   - DictionarySync[T, K]: The underlying synchronized dictionary storing key-value pairs.
//   - size: The maximum number of entries allowed in the dictionary.
//   - timeline: A Vector[T] maintaining the order of key insertions for eviction tracking.
//
// Example usage:
//     dict := DictionaryLimit[string, int]{size: 3}
//     dict.Put("a", 1)
//     dict.Put("b", 2)
//     dict.Put("c", 3)
//     dict.Put("d", 4) // "a" is removed if using FIFO eviction
//
//     value, ok := dict.Get("b") // value = 2, ok = true
type DictionaryLimit[T comparable, K any] struct {
	DictionarySync[T, K]
	size     int
	timeline Vector[T]
}

// DictionaryLimitFromMap creates a new DictionaryLimit instance from an existing map while enforcing a size limit.
// It initializes the dictionary with up to `size` key-value pairs from the given map.
//
// Parameters:
//   - size: The maximum number of entries allowed in the dictionary.
//   - items: A map containing initial key-value pairs.
//
// Returns:
//   - A pointer to a new IDictionary instance containing up to `size` elements.
//
// Notes:
//   - If the provided map has more elements than `size`, only the first `size` elements (based on map iteration order) will be included.
//   - The timeline vector tracks the order of inserted keys for potential eviction policies.
//
// Example usage:
//     data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//     dict := DictionaryLimitFromMap(3, data) // dict will contain up to 3 items
func MakeDictionaryLimit[T comparable, K any](size int, items map[T]K) IDictionary[T, K] {
	return DictionaryLimitFromMap(size, items)
}

// DictionaryLimitFromMap creates a new DictionaryLimit instance from an existing map while enforcing a size limit.
// It initializes the dictionary with up to `size` key-value pairs from the given map.
//
// Parameters:
//   - size: The maximum number of entries allowed in the dictionary.
//   - items: A map containing initial key-value pairs.
//
// Returns:
//   - A pointer to a new DictionaryLimit instance containing up to `size` elements.
//
// Notes:
//   - If the provided map has more elements than `size`, only the first `size` elements (based on map iteration order) will be included.
//   - The timeline vector tracks the order of inserted keys for potential eviction policies.
//
// Example usage:
//     data := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
//     dict := DictionaryLimitFromMap(3, data) // dict will contain up to 3 items
func DictionaryLimitFromMap[T comparable, K any](size int, items map[T]K) *DictionaryLimit[T, K] {
	instance := &DictionaryLimit[T, K]{
		size:     size,
		timeline: *VectorEmpty[T](),
	}

	instance.items = make(map[T]K)

	count := 0
	for k, v := range items {
		if count == size || instance.timeline.Size() > size {
			break
		}
		instance.items[k] = v
		instance.timeline.Append(k)
		count++
	}

	return instance
}

// DictionaryLimitEmpty creates a new empty DictionaryLimit instance with a specified size limit.
//
// Parameters:
//   - size: The maximum number of entries allowed in the dictionary.
//
// Returns:
//   - A pointer to a new DictionaryLimit instance with no initial elements.
//
// Example usage:
//     dict := DictionaryLimitEmpty[string, int](5) //dictionary with a max size of 5
func DictionaryLimitEmpty[T comparable, K any](size int) *DictionaryLimit[T, K] {
	return DictionaryLimitFromMap(size, make(map[T]K))
}

// DictionaryLimitFromVector creates a new DictionaryLimit instance from a Vector.
// It applies a mapping function to transform each element of the Vector into a key.
//
// Parameters:
//   - size: The maximum number of entries allowed in the dictionary.
//   - vector: A Vector containing elements to be added as values in the dictionary.
//   - mapper: A function that takes an element of type K and returns a key of type T.
//
// Returns:
//   - A pointer to a new DictionaryLimit instance containing up to `size` elements.
//
// Notes:
//   - The `mapper` function is used to generate unique keys for each element in the Vector.
//   - If the Vector has more elements than `size`, only the first `size` elements will be included.
//
// Example usage:
//     vec := VectorFromList([]int{10, 20, 30})
//     dict := DictionaryLimitFromVector(2, vec, func(v int) string { return fmt.Sprintf("key_%d", v) })
//     // The dictionary will contain at most 2 key-value pairs like {"key_10": 10, "key_20": 20}
func DictionaryLimitFromVector[T comparable, K any](size int, vector Vector[K], mapper func(K) T) IDictionary[T, K] {
	return DictionaryLimitFromList(size, vector.items, mapper)
}

// DictionaryLimitFromList creates a new DictionaryLimit instance from a list of values.
// It maps each element in the list to a key using the provided mapping function, and inserts them
// into the dictionary while enforcing the specified size limit.
//
// Parameters:
//   - size: The maximum number of entries allowed in the dictionary.
//   - vector: A slice of values that will be mapped to keys and added to the dictionary.
//   - mapper: A function that takes an element of type K and returns a key of type T.
//
// Returns:
//   - A pointer to a new DictionaryLimit instance containing up to `size` key-value pairs.
//
// Notes:
//   - The `mapper` function is used to generate unique keys for each element in the list.
//   - If the list contains more elements than the specified `size`, only the first `size` elements will be included.
//
// Example usage:
//     values := []int{10, 20, 30, 40}
//     dict := DictionaryLimitFromList(3, values, func(v int) string { return fmt.Sprintf("key_%d", v) })
//     // The dictionary will contain at most 3 key-value pairs like {"key_10": 10, "key_20": 20, "key_30": 30}
func DictionaryLimitFromList[T comparable, K any](size int, vector []K, mapper func(K) T) IDictionary[T, K] {
	mapp := DictionaryLimitEmpty[T, K](size)
	count := 0
	for _, v := range vector {
		if count == size || mapp.timeline.Size() > size {
			break
		}
		key := mapper(v)
		mapp.Put(key, v)
		mapp.timeline.Append(key)
		count++
	}
	return mapp
}

// Put adds a key-value pair to the DictionaryLimit, updating the value if the key already exists.
// It returns the old value associated with the key, if any, and a boolean indicating whether
// the key already existed in the DictionaryLimit (true if it existed, false otherwise).
//
// Parameters:
//   - key: The key of type T to associate with the given value.
//   - item: The value of type K to be associated with the key.
//
// Returns:
//   - A pointer to the old value associated with the key, or nil if the key did not exist.
//   - A boolean indicating whether the key was already present in the DictionaryLimit (true if it existed).
//
// Example usage:
//     dict := DictionaryLimitFromMap(map[string]int{"a": 1, "b": 2})
//     oldValue, exists := dict.Put("a", 3) // oldValue will be a pointer to 1, exists will be true
//     oldValue, exists = dict.Put("c", 4)  // oldValue will be nil, exists will be false
func (c *DictionaryLimit[T, K]) Put(key T, item K) (*K, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	old, exists := c.items[key]
	c.items[key] = item

	if index := c.timeline.IndexOf(func(t T) bool {
		return key == t
	}); index != -1 {
		c.timeline.Remove(index)
	}

	c.timeline.Append(key)

	if c.timeline.Size() > 0 && c.timeline.Size() > c.size {
		first, _ := c.timeline.Shift()
		delete(c.items, *first)
	}

	return &old, exists
}

// PutIfAbsent adds a key-value pair to the DictionaryLimit only if the key does not already exist.
// If the key is already present, it does nothing and returns the existing value associated with the key,
// along with a boolean indicating that the key was already present.
//
// Parameters:
//   - key: The key of type T to associate with the given value.
//   - item: The value of type K to be associated with the key if the key is absent.
//
// Returns:
//   - A pointer to the old value associated with the key, or nil if the key was not found.
//   - A boolean indicating whether the key was already present in the DictionaryLimit (true if it existed, false if it was absent).
//
// Example usage:
//     dict := DictionaryLimitFromMap(map[string]int{"a": 1, "b": 2})
//     oldValue, exists := dict.PutIfAbsent("a", 3) // oldValue will be a pointer to 1, exists will be true
//     oldValue, exists = dict.PutIfAbsent("c", 4)  // oldValue will be nil, exists will be false
func (c *DictionaryLimit[T, K]) PutIfAbsent(key T, item K) (*K, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	old, exists := c.items[key]
	if exists {
		return &old, exists
	}

	c.items[key] = item

	if index := c.timeline.IndexOf(func(t T) bool {
		return key == t
	}); index != -1 {
		c.timeline.Remove(index)
	}
	c.timeline.Append(key)

	if c.timeline.Size() > 0 && c.timeline.Size() > c.size {
		first, _ := c.timeline.Shift()
		delete(c.items, *first)
	}

	return &old, exists
}

// PutAll adds all key-value pairs from another map to the DictionaryLimit
// overwriting any existing values for the keys that already exist in the DictionaryLimit.
//
// Parameters:
//   - items: A map of type map[T]K containing the key-value pairs to add to the DictionaryLimit.
//
// Returns:
//   - The DictionaryLimit itself, with all the new key-value pairs added.
//
// Example usage:
//     dict := DictionaryLimitFromMap(map[string]int{"a": 1, "b": 2})
//     otherMap := map[string]int{"b": 3, "c": 4}
//     dict.PutAll(otherMap) // dict will contain {"a": 1, "b": 3, "c": 4}
func (c *DictionaryLimit[T, K]) PutAll(items map[T]K) IDictionary[T, K] {
	c.mu.Lock()
	defer c.mu.Unlock()

	count := 0
	for key := range items {
		if count == c.size {
			break
		}
		c.items[key] = items[key]
		c.timeline.Append(key)
		count++
	}

	for range count {
		if c.timeline.Size() > 0 && c.timeline.Size() > c.size {
			first, _ := c.timeline.Shift()
			delete(c.items, *first)
		}
	}

	return c
}
