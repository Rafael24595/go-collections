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
	})

	if value, ok := mapped.Get("1"); !ok || value != "value-1" {
		t.Errorf("Expected %s but got %s", "value-1", value)
	}
}

func TestDictionaryMax(t *testing.T) {
	dict := collection.DictionaryFromMap(map[string]int{"a": 1, "b": 3, "c": 2})

	pair, value, ok := dict.Max(func(k string, v int) int { return v })

	if !ok {
		t.Fatal("expected ok == true")
	}

	expected_key := "b"
	expected_val, _ := dict.Get(expected_key)

	if pair.Key() != expected_key || pair.Value() != expected_val || value != expected_val {
		t.Errorf("expected (%s, %d, %d), got (%s, %d, %d)", expected_key, expected_val, expected_val, pair.Key(), pair.Value(), value)
	}
}

func TestDictionaryMin(t *testing.T) {
	dict := collection.DictionaryFromMap(map[string]int{"a": 1, "b": 3, "c": 2})

	pair, value, ok := dict.Min(func(k string, v int) int { return v })

	if !ok {
		t.Fatal("expected ok == true")
	}

	expected_key := "a"
	expected_val, _ := dict.Get(expected_key)

	if pair.Key() != expected_key || pair.Value() != expected_val || value != expected_val {
		t.Errorf("expected (%s, %d, %d), got (%s, %d, %d)", expected_key, expected_val, expected_val, pair.Key(), pair.Value(), value)
	}
}

func TestDictionaryMaxEmpty(t *testing.T) {
	dict := collection.DictionaryFromMap(map[string]int{})

	_, _, ok := dict.Max(func(k string, v int) int { return v })

	if ok {
		t.Errorf("expected ok == false for empty dictionary")
	}
}

func TestDictionaryMinEmpty(t *testing.T) {
	dict := collection.DictionaryFromMap(map[string]int{})

	_, _, ok := dict.Min(func(k string, v int) int { return v })

	if ok {
		t.Errorf("expected ok == false for empty dictionary")
	}
}

func TestDictionaryMaxWithPredicate(t *testing.T) {
	dict := collection.DictionaryFromMap(map[string]LangTest{
		"go":   {"Golang", 30},
		"rust": {"Rust", 25},
		"zig":  {"Zig", 40},
	})

	pair, value, ok := dict.Max(func(k string, v LangTest) int { return v.score })

	if !ok {
		t.Fatal("expected ok == true")
	}

	expected_key := "zig"
	expected_val, _ := dict.Get(expected_key)

	if pair.Key() != expected_key || pair.Value().name != expected_val.name || value != expected_val.score {
		t.Errorf("expected (%s, %s, %d), got (%s, %s, %d)", expected_key, expected_val.name, expected_val.score, pair.Key(), pair.Value().name, value)
	}
}

func TestDictionaryMinWithPredicate(t *testing.T) {
	dict := collection.DictionaryFromMap(map[string]LangTest{
		"go":   {"Golang", 30},
		"rust": {"Rust", 25},
		"zig":  {"Zig", 40},
	})

	pair, value, ok := dict.Min(func(k string, v LangTest) int {
		return v.score
	})

	if !ok {
		t.Fatal("expected ok == true")
	}

	expected_key := "rust"
	expected_val, _ := dict.Get(expected_key)

	if pair.Key() != expected_key || pair.Value().name != expected_val.name || value != expected_val.score {
		t.Errorf("expected (%s, %s, %d), got (%s, %s, %d)", expected_key, expected_val.name, expected_val.score, pair.Key(), pair.Value().name, value)
	}
}
