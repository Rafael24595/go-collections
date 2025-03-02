package collection

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Vector represents a dynamically-sized array-like collection that holds elements of type T.
// It provides various methods to interact with the collection, such as adding, removing, and accessing elements.
//
// Type parameters:
//   - T: The type of elements stored in the Vector.
//
// Fields:
//   - items: A slice that holds the elements of type T in the Vector.
//
// Example usage:
//     vec := Vector[int]{items: []int{1, 2, 3}}
//     vec.Append(4)
//     value, exists := vec.Get(2) // value will be 3, exists will be true
type Vector[T any] struct {
	items []T
}

// VectorFromList creates a new Vector from a given slice of elements.
// It takes a slice of type T and returns a pointer to a new Vector that holds the same elements.
//
// Parameters:
//   - items: A slice of elements of type T that will be used to populate the Vector.
//
// Returns:
//   - A pointer to a new Vector[T] containing the same elements as the provided slice.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3})
//     // vec will be a Vector containing [1, 2, 3]
func VectorFromList[T any](items []T) *Vector[T] {
	return &Vector[T]{
		items,
	}
}

// VectorEmpty creates and returns an empty Vector of type T.
// It initializes a new Vector with no elements, essentially a Vector with a slice of zero length.
//
// Returns:
//   - A pointer to a new empty Vector[T].
//
// Example usage:
//     emptyVec := VectorEmpty[int]() // emptyVec will be a Vector with no elements
func VectorEmpty[T any]() *Vector[T] {
	return VectorFromList(make([]T, 0))
}

// Size returns the number of elements currently stored in the Vector.
//
// Returns:
//   - The number of elements in the Vector as an integer (len(c.items)).
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3})
//     size := vec.Size() // size will be 3
func (c *Vector[T]) Size() int {
	return len(c.items)
}

// Contains checks whether any element in the Vector satisfies the given predicate function.
// It returns true if there is at least one element that matches the predicate, and false otherwise.
//
// Parameters:
//   - predicate: A function that takes an element of type T and returns a boolean indicating 
//                whether the element meets the condition.
//
// Returns:
//   - A boolean indicating whether any element in the Vector satisfies the predicate.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3})
//     containsEven := vec.Contains(func(v int) bool { return v%2 == 0 }) // containsEven will be true
//     containsGreaterThanFive := vec.Contains(func(v int) bool { return v > 5 }) // containsGreaterThanFive will be false
func (c *Vector[T]) Contains(predicate func(T) bool) bool {
	_, ok := c.FindOne(predicate)
	return ok
}

// IndexOf finds the index of the first element in the Vector that satisfies the given predicate function.
// It returns the index of the first matching element and a boolean indicating whether such an element exists.
//
// Parameters:
//   - predicate: A function that takes an element of type T and returns a boolean indicating whether the element meets the condition.
//
// Returns:
//   - The index of the first element that satisfies the predicate, or -1 if no element satisfies it.
//   - A boolean indicating whether an element was found (true if found, false if not).
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4})
//     index, found := vec.IndexOf(func(v int) bool { return v == 3 }) // index will be 2, found will be true
//     index, found := vec.IndexOf(func(v int) bool { return v == 5 }) // index will be -1, found will be false
func (c *Vector[T]) IndexOf(predicate func(T) bool) (int, bool) {
	for i, item := range c.items {
		if predicate(item) {
			return i, true
		}
	}
	return -1, false
}

// Find returns a slice of all elements in the Vector that satisfy the given predicate function.
// It applies the predicate to each element and returns all matching elements in a new slice.
//
// Parameters:
//   - predicate: A function that takes an element of type T and returns a boolean indicating whether the element meets the condition.
//
// Returns:
//   - A slice containing all elements from the Vector that satisfy the predicate.
//     If no elements match, it returns an empty slice.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4})
//     evenNumbers := vec.Find(func(v int) bool { return v%2 == 0 }) // evenNumbers will be [2, 4]
//     greaterThanFive := vec.Find(func(v int) bool { return v > 5 }) // greaterThanFive will be []
func (c *Vector[T]) Find(predicate func(T) bool) []T {
	filter := []T{}
	for _, v := range c.items {
		if predicate(v) {
			filter = append(filter, v)
		}
	}
	return filter
}

// FindOne searches for the first element in the Vector that satisfies the given predicate function.
// It returns a pointer to the first matching element and a boolean indicating whether such an element was found.
//
// Parameters:
//   - predicate: A function that takes an element of type T and returns a boolean indicating whether the element meets the condition.
//
// Returns:
//   - A pointer to the first element that satisfies the predicate, or nil if no element matches.
//   - A boolean indicating whether a matching element was found (true if found, false if not).
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4})
//     value, found := vec.FindOne(func(v int) bool { return v == 3 }) // value will be 3, found will be true
//     value, found := vec.FindOne(func(v int) bool { return v == 5 }) // value will be nil, found will be false
func (c *Vector[T]) FindOne(predicate func(T) bool) (*T, bool) {
	for _, v := range c.items {
		if predicate(v) {
			return &v, true
		}
	}
	return nil, false
}

// Get retrieves the element at the specified index in the Vector.
// It returns a pointer to the element and a boolean indicating whether the element exists at the given index.
//
// Parameters:
//   - index: The index of the element to retrieve.
//
// Returns:
//   - A pointer to the element of type T at the specified index, or nil if the index is out of bounds.
//   - A boolean indicating whether the element exists at the given index (true if valid, false if out of bounds).
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3})
//     value, exists := vec.Get(1) // value will be 2, exists will be true
//     value, exists := vec.Get(5) // value will be nil, exists will be false (index out of bounds)
func (c *Vector[T]) Get(index int) (*T, bool) {
	if index >= 0 && index < len(c.items) {
		return &c.items[index], true
	}
	return nil, false
}

// First returns the first element in the Vector.
// It calls the Get method with index 0 and returns the result.
//
// Returns:
//   - A pointer to the first element in the Vector, or nil if the Vector is empty.
//   - A boolean indicating whether the element exists (true if the Vector has at least one element, false if empty).
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3})
//     value, exists := vec.First() // value will be 1, exists will be true
//     emptyVec := VectorEmpty[int]()
//     value, exists := emptyVec.First() // value will be nil, exists will be false
func (c *Vector[T]) First() (*T, bool) {
	return c.Get(0)
}

// Last returns the last element in the Vector.
// It calls the Get method with the index of the last element (Size() - 1) and returns the result.
//
// Returns:
//   - A pointer to the last element in the Vector, or nil if the Vector is empty.
//   - A boolean indicating whether the element exists (true if the Vector has at least one element, false if empty).
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3})
//     value, exists := vec.Last() // value will be 3, exists will be true
//     emptyVec := VectorEmpty[int]()
//     value, exists := emptyVec.Last() // value will be nil, exists will be false
func (c *Vector[T]) Last() (*T, bool) {
	return c.Get(c.Size() - 1)
}

// Append adds one or more elements to the end of the Vector.
// It modifies the Vector by appending the provided items and returns the updated Vector.
//
// Parameters:
//   - items: One or more elements of type T to be added to the end of the Vector.
//
// Returns:
//   - The updated Vector with the appended elements.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2})
//     vec.Append(3) // vec will now contain [1, 2, 3]
//     vec.Append(4, 5) // vec will now contain [1, 2, 3, 4, 5]
func (c *Vector[T]) Append(items ...T) *Vector[T] {
	c.items = append(c.items, items...)
	return c
}

// Set replaces the element at the specified index in the Vector with a new value and returns a pointer 
// to the previous element along with a boolean indicating whether the operation was successful.
//
// Parameters:
//   - index: The position of the element to be replaced in the Vector.
//   - item: The new value to set at the specified index.
//
// Returns:
//   - A pointer to the previous element at the given index, or nil if the index is invalid.
//   - A boolean indicating whether the replacement was successful.
//
// Example usage:
//     vec := VectorFromList([]int{10, 20, 30})
//     old, ok := vec.Set(1, 25) // old = &20, ok = true, vec will be modified to [10, 25, 30]
//     old, ok = vec.Set(5, 50)  // old = nil, ok = false (index out of bounds)
func (c *Vector[T]) Set(index int,item T) (*T, bool) {
	if index < 0 || index > len(c.items)-1{
		return nil, false
	}

	old, exists := c.Get(index)

	c.items[index] = item

	return old, exists
}

// AppendIfAbsent adds one or more elements to the end of the Vector, but only if the element does not already exist
// based on the provided predicate function. The predicate is used to check whether an element already exists in the Vector.
// If the element is absent, it will be appended; if present, it will be ignored.
//
// Parameters:
//   - predicate: A function that takes two elements of type T and returns a boolean indicating whether the elements
//                are considered equal or "matching" according to the condition defined in the predicate.
//   - items: One or more elements of type T to be appended to the Vector if they are not already present.
//
// Returns:
//   - The updated Vector with the new elements appended (if they were absent).
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3})
//     vec.AppendIfAbsent(func(a, b int) bool { return a == b }, 4, 5) // vec will now contain [1, 2, 3, 4, 5]
//     vec.AppendIfAbsent(func(a, b int) bool { return a == b }, 2) // vec will remain [1, 2, 3, 4, 5], 2 is not added again
func (c *Vector[T]) AppendIfAbsent(predicate func(T, T) bool, items ...T) *Vector[T] {
    for _, v := range items {
        if c.Contains(func(t T) bool {
            return predicate(t, v)
        }) {
            c.items = append(c.items, v)
        }
    }
	return c
}

// Merge combines the elements of another Vector with the current Vector.
// It appends all elements from the provided Vector to the end of the current Vector and returns the updated Vector.
//
// Parameters:
//   - other: The Vector whose elements will be appended to the current Vector.
//
// Returns:
//   - The updated Vector with elements from both the current Vector and the provided Vector.
//
// Example usage:
//     vec1 := VectorFromList([]int{1, 2, 3})
//     vec2 := VectorFromList([]int{4, 5, 6})
//     vec1.Merge(vec2) // vec1 will now contain [1, 2, 3, 4, 5, 6]
func (c *Vector[T]) Merge(other Vector[T]) *Vector[T] {
	c.items = append(c.items, other.items...)
	return c
}

// Filter creates a new Vector containing only the elements that satisfy the given predicate function.
// It applies the predicate to each element in the Vector and returns a new Vector with only those elements that match the condition.
//
// Parameters:
//   - predicate: A function that takes an element of type T and returns a boolean indicating whether the element meets the condition.
//
// Returns:
//   - A new Vector containing only the elements from the original Vector that satisfy the predicate.
//     If no elements match, the returned Vector will be empty.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4, 5})
//     evenNumbers := vec.Filter(func(v int) bool { return v%2 == 0 }) // evenNumbers will be [2, 4]
//     greaterThanThree := vec.Filter(func(v int) bool { return v > 3 }) // greaterThanThree will be [4, 5]
func (c *Vector[T]) Filter(predicate func(T) bool) *Vector[T] {
	filter := []T{}
	for _, v := range c.items {
		if predicate(v) {
			filter = append(filter, v)
		}
	}
	return VectorFromList(filter)
}

// FilterSelf modifies the current Vector by retaining only the elements that satisfy the given predicate function.
// It applies the predicate to each element in the Vector and updates the Vector to include only the matching elements.
//
// Parameters:
//   - predicate: A function that takes an element of type T and returns a boolean indicating whether the element meets the condition.
//
// Returns:
//   - The updated Vector with only the elements that satisfy the predicate. The original Vector is directly modified.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4, 5})
//     vec.FilterSelf(func(v int) bool { return v%2 == 0 }) // vec will be modified to [2, 4]
//     vec.FilterSelf(func(v int) bool { return v > 3 }) // vec will be modified to [4]
func (c *Vector[T]) FilterSelf(predicate func(T) bool) *Vector[T] {
	filter := []T{}
	for _, v := range c.items {
		if predicate(v) {
			filter = append(filter, v)
		}
	}
    c.items = filter
	return c
}

// Remove deletes the element at the specified index from the Vector and returns a pointer to the removed element 
// along with a boolean indicating whether the element existed. If the index is out of bounds, it returns nil and false.
//
// Parameters:
//   - index: The position of the element to be removed in the Vector.
//
// Returns:
//   - A pointer to the removed element, or nil if the index is invalid.
//   - A boolean indicating whether the element was successfully removed.
//
// Example usage:
//     vec := VectorFromList([]int{10, 20, 30, 40})
//     removed, ok := vec.Remove(2) // removed = &30, ok = true, vec will be modified to [10, 20, 40]
//     removed, ok = vec.Remove(5)  // removed = nil, ok = false (index out of bounds)
func (c *Vector[T]) Remove(index int) (*T, bool) {
	if index < 0 || index > len(c.items)-1{
		return nil, false
	}

	old, exists := c.Get(index)

	c.items = append(c.items[:index], c.items[index:]...)

	return old, exists
}

// Slice creates a new Vector from a portion of the current Vector, defined by the start and end indices.
// It slices the Vector's elements from the `start` index (inclusive) to the `end` index (exclusive), adjusting
// the indices if they are out of bounds. If the start or end index is out of range, it will be clamped to valid values.
// A new Vector containing the sliced elements is returned.
//
// Parameters:
//   - start: The index to begin slicing from (inclusive). If out of bounds, it will be adjusted to 0 or the start of the Vector.
//   - end: The index to end slicing at (exclusive). If out of bounds, it will be adjusted to the length of the Vector.
//
// Returns:
//   - A new Vector containing the sliced elements from the original Vector. The original Vector remains unchanged.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4, 5})
//     slicedVec := vec.Slice(1, 4) // slicedVec will contain [2, 3, 4]
//     slicedVec2 := vec.Slice(0, 2) // slicedVec2 will contain [1, 2]
//     slicedVec3 := vec.Slice(6, 10) // slicedVec3 will contain []
func (c *Vector[T]) Slice(start, end int) *Vector[T] {
	if start < 0 {
		start = 0
	}
	if start > len(c.items)-1 {
		start = len(c.items)
	}
	if end > len(c.items)-1 {
		end = len(c.items)
	}
	return VectorFromList(c.items[start:end])
}

// SliceSelf modifies the current Vector from a portion of the current Vector, defined by the start and end indices.
// It slices the Vector's elements from the `start` index (inclusive) to the `end` index (exclusive), adjusting the indices
// if they are out of bounds. If the start or end index is out of range, it will be clamped to valid values.
//
// Parameters:
//   - start: The index to begin slicing from (inclusive). If out of bounds, it will be adjusted to 0 or the start of the Vector.
//   - end: The index to end slicing at (exclusive). If out of bounds, it will be adjusted to the length of the Vector.
//
// Returns:
//   - The updated Vector, containing the sliced elements from the original Vector. The original Vector is directly modified.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4, 5})
//     vec.Clone().Slice(1, 4) // vec will be modified to [2, 3, 4]
//     vec.Clone().Slice(0, 2) // vec will be modified to [1, 2]
//     vec.Clone().Slice(6, 10) // vec will be modified to []
func (c *Vector[T]) SliceSelf(start, end int) *Vector[T] {
	if start < 0 {
		start = 0
	}
	if start > len(c.items)-1 {
		start = len(c.items)
	}
	if end > len(c.items)-1 {
		end = len(c.items)
	}
	c.items = c.items[start:end]
	return c
}

// Shift removes and returns the first element of the Vector, shifting all remaining elements left.
// If the Vector is empty, it returns nil and false.
//
// Returns:
//   - A pointer to the removed element, or nil if the Vector is empty.
//   - A boolean indicating whether the operation was successful.
//
// Example usage:
//     vec := VectorFromList([]int{10, 20, 30})
//     first, ok := vec.Shift() // first = &10, ok = true, vec will be modified to [20, 30]
//     first, ok = vec.Shift()  // first = &20, ok = true, vec will be modified to [30]
//     first, ok = vec.Shift()  // first = &30, ok = true, vec will be modified to []
//     first, ok = vec.Shift()  // first = nil, ok = false (empty Vector)
func (c *Vector[T]) Shift() (*T, bool) {
	if len(c.items) == 0 {
		return nil, false
	}

	first := c.items[0]
	c.items = c.items[1:]

	return &first, true
}

// JoinBy groups elements in the Vector based on a key generated by the indexer function, 
// and combines the grouped elements using the provided predicate function.
// If multiple elements share the same key, the predicate function is used to merge them 
// into a single element. The original Vector is modified to reflect the grouped and merged elements.
//
// Parameters:
//   - indexer: A function that extracts a key of type string from an element of type T.
//   - predicate: A function that takes two elements of type T and merges them into one element of type T.
//
// Returns:
//   - The modified Vector, containing the merged elements, allowing for method chaining.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 2, 3, 3, 3})
//     vec.JoinBy(func(v int) string { return fmt.Sprintf("key-%d", v) }, 
//                func(i, j int) int { return i + j })
//     // vec will be modified to [1, 4, 9], where values 2 and 3 have been merged as per the predicate
func (c *Vector[T]) JoinBy(indexer func(T) string, predicate func(i, j T) T) *Vector[T] {
	dict := map[string]T{}
	for _, item := range c.items {
		key := indexer(item)
		aux := item
		if found, ok := dict[key]; ok {
			aux = predicate(found, item)
		}
		dict[key] = aux
	}

	c.items = make([]T, 0)
	for _, v := range dict {
		c.items = append(c.items, v)
	}

	return c
}

// ForEach applies the given predicate function to each element in the Vector, passing both the index and the element itself.
// It allows you to perform operations on each element of the Vector, such as printing, modifying external state, or aggregating data.
// The original Vector is not modified.
//
// Parameters:
//   - predicate: A function that takes the index of the element (int) and the element itself (T) as arguments.
//     This function will be executed for each element in the Vector.
//
// Returns:
//   - The current Vector, allowing for method chaining.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4})
//     vec.ForEach(func(i, v int) {
//         fmt.Println(i, v) // Prints each index and value (0 1, 1 2, 2 3, 3 4)
//     })
func (c *Vector[T]) ForEach(predicate func(int, T)) *Vector[T] {
	for i, v := range c.items {
		predicate(i, v)
	}
	return c
}

// Map transforms each element in the Vector by applying the given predicate function to it.
// The predicate function takes both the index (int) and the element (T) as arguments, 
// and returns a transformed element of the same type T. This method directly modifies 
// the original Vector with the transformed elements.
//
// Parameters:
//   - predicate: A function that takes the index (int) and an element of type T, 
//     and returns a transformed element of type T.
//
// Returns:
//   - The current Vector with the transformed elements, allowing for method chaining.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4})
//     vec.Map(func(i, v int) int { return v * (i + 1) }) 
//     // vec will be modified to [1, 4, 9, 16] (multiplying each element by its index + 1)
func (c *Vector[T]) Map(predicate func(int, T) T) *Vector[T] {
	for i, item := range c.items {
		c.items[i] = predicate(i, item)
	}
	return c
}

// Clean clears all elements in the Vector, resetting it to an empty state.
// This method modifies the original Vector, and returns the same Vector instance for method chaining.
//
// Returns:
//   - The current, empty Vector, allowing for method chaining.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4})
//     vec.Clean() // vec will be modified to an empty Vector []
func (c *Vector[T]) Clean() *Vector[T] {
	c.items = make([]T, 0)
	return c
}

// Clone creates a new Vector that is a shallow copy of the original Vector.
// It duplicates all the elements in the current Vector, ensuring that the new Vector
// is independent of the original one, with no shared references.
//
// Returns:
//   - A new Vector that is a clone of the original Vector.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4})
//     clonedVec := vec.Clone() // clonedVec will be a new Vector with the same elements: [1, 2, 3, 4]
func (c *Vector[T]) Clone() *Vector[T] {
	cloned := make([]T, len(c.items))
	copy(cloned, c.items)
	return VectorFromList(cloned)
}

// Sort sorts the elements of the Vector in-place using the provided comparison function.
// The comparison function should return true if the element at index i should be ordered before
// the element at index j, and false otherwise.
//
// Parameters:
//   - less: A comparison function that takes two elements of type T (i and j), and returns a boolean.
//           It should return true if i should come before j in the sorted order.
//
// Returns:
//   - The current Vector with its elements sorted, allowing for method chaining.
//
// Example usage:
//     vec := VectorFromList([]int{4, 1, 3, 2})
//     vec.Sort(func(i, j int) bool { return i < j }) // vec will be sorted to [1, 2, 3, 4]
func (c *Vector[T]) Sort(less func(i, j T) bool) *Vector[T] {
	sort.Slice(c.items, func(i, j int) bool {
		return less(c.items[i], c.items[j])
	})
	return c
}

// Collect returns a slice containing all the elements in the Vector.
// This method does not modify the original Vector; it simply gives direct access to the internal slice, allowing the caller to interact with it as a regular, allowing the caller to interact with it as a regular map.
//
// Returns:
//   - A slice of type T containing all elements in the Vector.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4})
//     items := vec.Collect() // items will be a slice: [1, 2, 3, 4]
func (c Vector[T]) Collect() []T {
	return c.items
}

// Join combines all elements of the Vector into a single string, separated by the specified separator.
// If the elements of the Vector are already strings, it uses the strings.Join function to join them.
// Otherwise, it converts each element into a string using fmt.Sprintf and then joins them.
//
// Parameters:
//   - separator: A string that separates the elements in the resulting joined string.
//
// Returns:
//   - A single string containing all the elements of the Vector, separated by the provided separator.
//
// Example usage:
//     vec := VectorFromList([]string{"apple", "banana", "cherry"})
//     result := vec.Join(", ") // result will be "apple, banana, cherry"
//
//     vec2 := VectorFromList([]int{1, 2, 3})
//     result2 := vec2.Join(" - ") // result2 will be "1 - 2 - 3"
func (c *Vector[T]) Join(separator string) string {
	if items, ok := interface{}(c.items).([]string); ok {
		return strings.Join(items, separator)
	}
	return VectorMap(c, func(i T) string {
		return fmt.Sprintf("%v", i)
	}).Join(separator)
}

// Pages calculates the number of pages required to hold all the elements of the Vector,
// given the specified page size. It uses the ceiling function to round up to the next whole page 
// if there are leftover items that don't fill an entire page.
//
// Parameters:
//   - size: The maximum number of elements allowed on each page.
//
// Returns:
//   - The total number of pages required to hold all the elements, rounded up.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4, 5, 6, 7})
//     pages := vec.Pages(3) // pages will be 3, as the items are split across three pages of 3 elements each
func (c *Vector[T]) Pages(size int) int {
	len := float64(len(c.items))
	fSize := float64(size)
	return int(math.Ceil(len / fSize))
}

// Page returns a subset (or "page") of elements from the Vector, based on the specified page number and page size.
// It uses the provided page number to calculate the start and end indices, and then returns the corresponding slice.
// If the page number is 0, it defaults to page 1.
//
// Parameters:
//   - page: The 1-based index of the page to retrieve.
//   - size: The maximum number of elements per page.
//
// Returns:
//   - A new Vector containing the elements from the specified page.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
//     page1 := vec.Page(1, 3) // page1 will contain [1, 2, 3]
//     page2 := vec.Page(2, 3) // page2 will contain [4, 5, 6]
//     page3 := vec.Page(3, 3) // page3 will contain [7, 8, 9]
//     page4 := vec.Page(4, 3) // page4 will contain [10]
func (c *Vector[T]) Page(page, size int) *Vector[T] {
	if page == 0 {
		page = 1
	}
	start := (page - 1) * size
	end := page * size
	return c.Slice(start, end)
}

// VectorMap applies the given predicate function to each element in the Vector, 
// transforming each element of type T into an element of type K, and returns 
// a new Vector with the transformed elements.
//
// Parameters:
//   - c: The source Vector containing elements of type T.
//   - predicate: A function that takes an element of type T and transforms it into an element of type K.
//
// Returns:
//   - A new Vector containing the transformed elements of type K.
//
// Example usage:
//     vec := VectorFromList([]int{1, 2, 3, 4})
//     transformed := VectorMap(vec, func(v int) string { return fmt.Sprintf("Item %d", v) })
//     // transformed will be a new Vector with elements: ["Item 1", "Item 2", "Item 3", "Item 4"]
func VectorMap[T, K any](c *Vector[T], predicate func(T) K) *Vector[K] {
	var mapped []K
	for _, item := range c.items {
		mapped = append(mapped, predicate(item))
	}
	return &Vector[K]{
		items: mapped,
	}
}
