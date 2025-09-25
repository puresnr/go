package pmap

import (
	"reflect"
	"testing"

	"github.com/puresnr/go/deepcopy/constraint"
)

// MyStruct for testing Deepcopyable values in maps
type MyStruct struct {
	Value int
	Ptr   *int
}

func (m MyStruct) Deepcopy() MyStruct {
	newPtr := new(int)
	*newPtr = *m.Ptr
	return MyStruct{
		Value: m.Value,
		Ptr:   newPtr,
	}
}

// Ensure MyStruct implements Deepcopyable
var _ constraint.Deepcopyable[MyStruct] = MyStruct{}

func TestDeepcopyBasic(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		var m1 map[string]int
		m2 := DeepcopyBasic(m1)
		if m2 != nil {
			t.Errorf("DeepcopyBasic(nil) should be nil, got %v", m2)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		m1 := map[string]int{}
		m2 := DeepcopyBasic(m1)
		if m2 != nil {
			t.Errorf("DeepcopyBasic of empty map should be nil, got %v", m2)
		}
	})

	t.Run("int map", func(t *testing.T) {
		m1 := map[string]int{"a": 1, "b": 2}
		m2 := DeepcopyBasic(m1)

		if !reflect.DeepEqual(m1, m2) {
			t.Errorf("Cloned map %v is not equal to original %v", m2, m1)
		}

		// Modify original and check if clone is affected
		m1["a"] = 99
		if m2["a"] == 99 {
			t.Errorf("Clone is not a deep copy; it was modified when the original changed.")
		}
	})
}

func TestDeepcopy(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		var m1 map[string]MyStruct
		m2 := Deepcopy(m1)
		if m2 != nil {
			t.Errorf("Deepcopy(nil) should be nil, got %v", m2)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		m1 := map[string]MyStruct{}
		m2 := Deepcopy(m1)
		if m2 != nil {
			t.Errorf("Deepcopy of empty map should be nil, got %v", m2)
		}
	})

	t.Run("Deepcopyable struct map", func(t *testing.T) {
		ptr1 := new(int)
		*ptr1 = 10
		ptr2 := new(int)
		*ptr2 = 20

		m1 := map[string]MyStruct{
			"key1": {Value: 1, Ptr: ptr1},
			"key2": {Value: 2, Ptr: ptr2},
		}

		m2 := Deepcopy(m1)

		if !reflect.DeepEqual(m1, m2) {
			t.Errorf("Deepcopied map %v is not equal to original %v", m2, m1)
		}

		// Modify original and check if clone is affected
		modifiedStruct := m1["key1"]
		modifiedStruct.Value = 99
		*modifiedStruct.Ptr = 999
		m1["key1"] = modifiedStruct

		if m2["key1"].Value == 99 || *m2["key1"].Ptr == 999 {
			t.Errorf("Deepcopy is not a deep copy; it was modified when the original changed.")
		}
	})
}
