package collection

import (
	"testing"

	"github.com/Rafael24595/go-collections/collection"
)

func TestDictionaryLimitedPut(t *testing.T) {
	dict := collection.DictionaryLimitedFromMap(10, map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"10": 10,
		"11": 11,
	})
	if dict.Size() > 10 {
        t.Errorf("Expected %d but got %d", 10, dict.Size())
    }

	dict.Put("12", 12)
	if dict.Size() > 10 {
        t.Errorf("Expected %d but got %d", 10, dict.Size())
    }
}

func TestDictionaryLimitedPutAll(t *testing.T) {
	dict := collection.DictionaryLimitedFromMap(10, map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"10": 10,
		"11": 11,
	})
	if dict.Size() > 10 {
        t.Errorf("Expected %d but got %d", 10, dict.Size())
    }

	dict.PutAll(map[string]int{
		"12": 12,
		"13": 13,
		"14": 14,
	})
	if dict.Size() > 10 {
        t.Errorf("Expected %d but got %d", 10, dict.Size())
    }

	_, exists12 := dict.Get("12")
	_, exists13 := dict.Get("13")
	_, exists14 := dict.Get("14")

	if !exists12 || !exists13 || !exists14 {
		t.Errorf("Last inserted elements are removed")
	}
}

func TestDictionaryLimitedPutIfAbsent(t *testing.T) {
	dict := collection.DictionaryLimitedFromMap(10, map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"10": 10,
		"11": 11,
	})
	if dict.Size() > 10 {
        t.Errorf("Expected %d but got %d", 10, dict.Size())
    }

	dict.PutIfAbsent("12", 12)
	if dict.Size() > 10 {
        t.Errorf("Expected %d but got %d", 10, dict.Size())
    }
	if dict.Size() > 10 {
        t.Errorf("Expected %d but got %d", 10, dict.Size())
    }

	_, exists := dict.Get("12")

	if !exists {
		t.Errorf("Last inserted element are removed")
	}
}
