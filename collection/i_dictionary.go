package collection

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