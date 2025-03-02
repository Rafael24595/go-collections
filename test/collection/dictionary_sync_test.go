package collection

import (
	"strconv"
	"sync"
	"testing"

	"github.com/Rafael24595/go-collections/collection"
)

func TestDictionarySyncStress(t *testing.T) {
	dict := collection.DictionarySyncEmpty[string, int]()

	var wg sync.WaitGroup
	n := 10000

	wg.Add(n*3)

	for i := range n {
		key := strconv.Itoa(i)
		go func(i int, key string) {
			defer wg.Done()
			dict.Put(key, i)
		}(i, key)
		go func(i int, key string) {
			defer wg.Done()
			dict.Map(func(s string, i int) int {
				return i + 1
			})
		}(i, key)
		go func(i int, key string) {
			defer wg.Done()
			dict.Remove(key)
		}(i, key)
	}

	wg.Wait()
}
