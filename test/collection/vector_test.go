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

func TestVectorSet(t *testing.T) {
	vector := collection.VectorFromList([]int{
		1, 2, 3,
	})

	position, ok := vector.Get(1)
	if !ok || *position != 2 {
        t.Errorf("Vector does not contains %d but it is added.", 2)
    }

	vector.Set(1, 4)
	position, ok = vector.Get(1)
	if !ok || *position != 4 {
        t.Errorf("Vector does not contains %d but it is added.", 2)
    }
}

func TestVectorRemove(t *testing.T) {
	vector := collection.VectorFromList([]int{
		1, 2, 3,
	})

	position, ok := vector.Remove(1)
	if !ok || *position != 2 {
        t.Errorf("Vector does not contains %d but it is added.", 2)
    }

	if len := vector.Size(); len < 2 {
        t.Errorf("Expected %d but got %d", 2, len)
    }
}

func TestVectorShift(t *testing.T) {
	vector := collection.VectorFromList([]int{
		1, 2, 3,
	})

	position, ok := vector.Shift()
	if !ok || *position != 1 {
        t.Errorf("Vector does not contains %d but it is added.", 2)
    }

	if len := vector.Size(); len < 2 {
        t.Errorf("Expected %d but got %d", 2, len)
    }
}

func TestVectorContains(t *testing.T) {
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

func TestVectorAppendIfAbsent(t *testing.T) {
	vector := collection.VectorFromList([]int{
		0, 1, 2, 3,
	})

	vector.AppendIfAbsent(func(i1, i2 int) bool {
		return i1 != i2
	}, 2)

	if vector.Size() > 4 {
        t.Errorf("Vector size is %d but %d expected", vector.Size(), 4)
    }

	vector.AppendIfAbsent(func(i1, i2 int) bool {
		return i1 != i2
	}, 4, 5)

	if vector.Size() > 6 {
        t.Errorf("Vector size is %d but %d expected", vector.Size(), 6)
    }
}
