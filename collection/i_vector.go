package collection

type VectorConstructor[I any] func([]I) IVector[I]

type IVector[I any] interface {
	Size() int
	Contains(predicate func(I) bool) bool
	IndexOf(predicate func(I) bool) int
	Find(predicate func(I) bool) []I
	FindOne(predicate func(I) bool) (I, bool)
	Get(index int) (I, bool)
	First() (I, bool)
	Last() (I, bool)
	Append(items ...I) *Vector[I]
	Set(index int, item I) (I, bool)
	AppendIfAbsent(predicate func(I, I) bool, items ...I) *Vector[I]
	Merge(other Vector[I]) *Vector[I]
	Filter(predicate func(I) bool) *Vector[I]
	FilterSelf(predicate func(I) bool) *Vector[I]
	Remove(index int) (I, bool)
	Slice(start, end int) *Vector[I]
	SliceSelf(start, end int) *Vector[I]
	Unshift(items ...I) *Vector[I]
	Shift() (I, bool)
	JoinBy(indexer func(I) string, predicate func(i, j I) I) *Vector[I]
	ForEach(predicate func(int, I)) *Vector[I]
	Map(predicate func(int, I) I) *Vector[I]
	Clean() *Vector[I]
	Clone() *Vector[I]
	Sort(less func(i, j I) bool) *Vector[I]
	Max(predicate func(I) int) (I, int, bool)
	Min(predicate func(I) int) (I, int, bool)
	Collect() []I
	Join(separator string) string
	Pages(size int) int
	Page(page, size int) *Vector[I]
}

// IVectorMap applies the given predicate function to each element in the IVector,
// transforming each element of type I into an element of type K, and returns
// a new Vector with the transformed elements.
//
// Parameters:
//   - c: The source IVector containing elements of type I.
//   - predicate: A function that takes an element of type I and transforms it into an element of type K.
//   - constructor: A function that instance a new IVector implementation, and return it with the mapped values.
//
// Returns:
//   - A new IVector containing the transformed elements of type K.
//
// Example usage:
//
//	vec := VectorFromList([]int{1, 2, 3, 4})
//	transformed := IVectorMap(vec, func(v int) string { return fmt.Sprintf("Item %d", v) }, MakeVector)
//	// transformed will be a new Vector with elements: ["Item 1", "Item 2", "Item 3", "Item 4"]
func IVectorMap[I, K any](c *Vector[I], predicate func(I) K, constructor VectorConstructor[K]) IVector[K] {
	return MapToIVector(c.Collect(), predicate, constructor)
}

// MapToVector applies the given predicate function to each element in the slice,
// transforming each element of type I into an element of type K, and returns
// a IVector with the transformed elements.
//
// Parameters:
//   - c: The source slice containing elements of type I.
//   - predicate: A function that takes an element of type I and transforms it into an element of type K.
//   - constructor: A function that instance a new IVector implementation, and return it with the mapped values.
//
// Returns:
//   - A new IVector containing the transformed elements of type K.
//
// Example usage:
//
//	slc := []int{1, 2, 3, 4}
//	transformed := MapToIVector(slc, func(v int) string { return fmt.Sprintf("Item %d", v) }, MakeVector)
//	// transformed will be a new Vector with elements: ["Item 1", "Item 2", "Item 3", "Item 4"]
func MapToIVector[I, K any](c []I, predicate func(I) K, constructor VectorConstructor[K]) IVector[K] {
	mapped := make([]K, len(c))
	for i, item := range c {
		mapped[i] = predicate(item)
	}
	return constructor(mapped)
}
