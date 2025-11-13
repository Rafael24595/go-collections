package collection

type VectorConstructor[T any] func([]T) IVector[T]

type IVector[T any] interface {
	Size() int
	Contains(predicate func(T) bool) bool
	IndexOf(predicate func(T) bool) (int, bool)
	Find(predicate func(T) bool) []T
	FindOne(predicate func(T) bool) (*T, bool)
	Get(index int) (*T, bool)
	First() (*T, bool)
	Last() (*T, bool)
	Append(items ...T) *Vector[T]
	Set(index int, item T) (*T, bool)
	AppendIfAbsent(predicate func(T, T) bool, items ...T) *Vector[T]
	Merge(other Vector[T]) *Vector[T]
	Filter(predicate func(T) bool) *Vector[T]
	FilterSelf(predicate func(T) bool) *Vector[T]
	Remove(index int) (*T, bool)
	Slice(start, end int) *Vector[T]
	SliceSelf(start, end int) *Vector[T]
	Shift() (*T, bool)
	JoinBy(indexer func(T) string, predicate func(i, j T) T) *Vector[T]
	ForEach(predicate func(int, T)) *Vector[T]
	Map(predicate func(int, T) T) *Vector[T]
	Clean() *Vector[T]
	Clone() *Vector[T]
	Sort(less func(i, j T) bool) *Vector[T]
	Collect() []T
	Join(separator string) string
	Pages(size int) int
	Page(page, size int) *Vector[T]
}

// VectorMap applies the given predicate function to each element in the Vector,
// transforming each element of type T into an element of type K, and returns
// a new Vector with the transformed elements.
//
// Parameters:
//   - c: The source Vector containing elements of type T.
//   - predicate: A function that takes an element of type T and transforms it into an element of type K.
//   - constructor: A function that instance a new IVector implementation, and return it with the mapped values.
//
// Returns:
//   - A new Vector containing the transformed elements of type K.
//
// Example usage:
//
//	vec := VectorFromList([]int{1, 2, 3, 4})
//	transformed := VectorMap(vec, func(v int) string { return fmt.Sprintf("Item %d", v) })
//	// transformed will be a new Vector with elements: ["Item 1", "Item 2", "Item 3", "Item 4"]
func VectorMap[T, K any](c *Vector[T], predicate func(T) K, constructor VectorConstructor[K]) IVector[K] {
	return MapToVector(c.Collect(), predicate, constructor)
}

// MapToVector applies the given predicate function to each element in the slice,
// transforming each element of type T into an element of type K, and returns
// a Vector with the transformed elements.
//
// Parameters:
//   - c: The source Vector containing elements of type T.
//   - predicate: A function that takes an element of type T and transforms it into an element of type K.
//   - constructor: A function that instance a new IVector implementation, and return it with the mapped values.
//
// Returns:
//   - A new Vector containing the transformed elements of type K.
//
// Example usage:
//
//	vec := VectorFromList([]int{1, 2, 3, 4})
//	transformed := VectorMap(vec, func(v int) string { return fmt.Sprintf("Item %d", v) })
//	// transformed will be a new Vector with elements: ["Item 1", "Item 2", "Item 3", "Item 4"]
func MapToVector[T, K any](c []T, predicate func(T) K, constructor VectorConstructor[K]) IVector[K] {
	mapped := make([]K, len(c))
	for i, item := range c {
		mapped[i] = predicate(item)
	}
	return constructor(mapped)
}
