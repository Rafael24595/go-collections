package collection

import (
	"testing"
	"github.com/Rafael24595/go-collections/collection"
)

func TestVectorSize(t *testing.T) {
	vector := collection.VectorEmpty[int]()
	vector.Append(0)

	result := vector.Size()
    expected := 1

    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}

func TestContains(t *testing.T) {
	vector := collection.VectorEmpty[int]()
	vector.Append(0)

	resultTrue := vector.Contains(func(i int) bool {
		return i == 0
	})

	resultFalse := vector.Contains(func(i int) bool {
		return i == 9
	})

    if !resultTrue {
        t.Errorf("Vector does not contains %d but it is added.", 0)
    }

	if resultFalse {
        t.Errorf("Vector contains %d but it is not added.", 9)
    }
}