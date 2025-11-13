package collection

import (
	"fmt"
	"testing"

	"github.com/Rafael24595/go-collections/collection"
)

func TestDictionaryMap(t *testing.T) {
	dict := collection.DictionaryFromMap(map[string]int{
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
		"11": 11,
	})

	mapped := collection.DictionaryMap(dict, func(k string, v int) string {
		return fmt.Sprintf("value-%d", v)
	}, collection.MakeDictionary)

	if value, ok := mapped.Get("1"); !ok || *value != "value-1" {
		t.Errorf("Expected %s but got %s", "value-1", *value)
	}
}
