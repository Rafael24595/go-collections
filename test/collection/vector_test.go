package collection

import (
	"testing"

	"github.com/Rafael24595/go-collections/collection"
)

type LangTest struct {
	name  string
	score int
}

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
	if !ok || position != 2 {
		t.Errorf("Vector does not contains %d but it is added.", 2)
	}

	vector.Set(1, 4)
	position, ok = vector.Get(1)
	if !ok || position != 4 {
		t.Errorf("Vector does not contains %d but it is added.", 2)
	}
}

func TestVectorRemove(t *testing.T) {
	vector := collection.VectorFromList([]int{
		1, 2, 3,
	})

	position, ok := vector.Remove(1)
	if !ok || position != 2 {
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
	if !ok || position != 1 {
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

func TestVectorMax(t *testing.T) {
	vec := collection.VectorFromList([]int{4, 1, 3, 2})

	item, value, ok := vec.Max(func(i int) int {
		return i
	})

	if !ok {
		t.Fatal("expected ok == true")
	}

	expected, _ := vec.Get(0)

	value_expected := 4

	if item != expected || value != value_expected {
		t.Errorf("expected (%d, %d), got (%d, %d)", expected, value_expected, item, value)
	}
}

func TestVectorMaxWithPredicate(t *testing.T) {
	vec := collection.VectorFromList([]LangTest{
		{"Golang", 30},
		{"Rust", 25},
		{"Zig", 40},
	})

	item, value, ok := vec.Max(func(l LangTest) int {
		return l.score
	})

	if !ok {
		t.Fatal("expected ok == true")
	}

	expected, _ := vec.Get(2)

	name_expexted := expected.name
	value_expected := expected.score

	if item.name != name_expexted || value != value_expected {
		t.Errorf("expected (Zig, 40), got (%s, %d)", item.name, value)
	}
}

func TestVectorMin(t *testing.T) {
	vec := collection.VectorFromList([]int{4, 1, 3, 2})

	item, value, ok := vec.Min(func(i int) int {
		return i
	})

	if !ok {
		t.Fatal("expected ok == true")
	}

	expected, _ := vec.Get(1)

	value_expected := 1

	if item != expected || value != value_expected {
		t.Errorf("expected (%d, %d), got (%d, %d)", expected, value_expected, item, value)
	}
}

func TestVectorMinWithPredicate(t *testing.T) {
	vec := collection.VectorFromList([]LangTest{
		{"Golang", 30},
		{"Rust", 25},
		{"Zig", 40},
	})

	item, value, ok := vec.Min(func(l LangTest) int {
		return l.score
	})

	if !ok {
		t.Fatal("expected ok == true")
	}

	expected, _ := vec.Get(1)

	name_expexted := expected.name
	value_expected := expected.score

	if item.name != name_expexted || value != value_expected {
		t.Errorf("expected (%s, %d), got (%s, %d)", name_expexted, value_expected, item.name, value)
	}
}

func TestVectorMaxEmpty(t *testing.T) {
	vec := collection.VectorFromList([]int{})

	_, _, ok := vec.Max(func(i int) int {
		return i
	})

	if ok {
		t.Fatal("expected ok == false")
	}
}

func TestVectorMinEmpty(t *testing.T) {
	vec := collection.VectorFromList([]int{})

	_, _, ok := vec.Min(func(i int) int {
		return i
	})

	if ok {
		t.Fatal("expected ok == false")
	}
}
