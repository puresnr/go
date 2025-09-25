package pslice

import (
	"reflect"
	"testing"

	"github.com/puresnr/go/deepcopy/constraint"
)

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
	t.Run("nil slice", func(t *testing.T) {
		var s1 []int
		s2 := DeepcopyBasic(s1)
		if s2 != nil {
			t.Errorf("DeepcopyBasic(nil) should be nil, got %v", s2)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		s1 := []int{}
		s2 := DeepcopyBasic(s1)
		if len(s2) != 0 {
			t.Errorf("Clone of empty slice should be an empty slice, got %v", s2)
		}
	})

	t.Run("int slice", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := DeepcopyBasic(s1)

		if !reflect.DeepEqual(s1, s2) {
			t.Errorf("Cloned slice %v is not equal to original %v", s2, s1)
		}

		// Modify original and check if clone is affected
		s1[0] = 99
		if s2[0] == 99 {
			t.Errorf("Clone is not a deep copy; it was modified when the original changed.")
		}
	})

	t.Run("string slice", func(t *testing.T) {
		s1 := []string{"a", "b", "c"}
		s2 := DeepcopyBasic(s1)

		if !reflect.DeepEqual(s1, s2) {
			t.Errorf("Cloned slice %v is not equal to original %v", s2, s1)
		}

		// Modify original and check if clone is affected
		s1[0] = "z"
		if s2[0] == "z" {
			t.Errorf("Clone is not a deep copy; it was modified when the original changed.")
		}
	})
}

func TestDeepcopy(t *testing.T) {
	t.Run("Deepcopyable struct slice", func(t *testing.T) {
		ptr1 := new(int)
		*ptr1 = 10
		s1 := []MyStruct{
			{Value: 1, Ptr: ptr1},
			{Value: 2, Ptr: new(int)},
		}
		*s1[1].Ptr = 20

		s2 := Deepcopy(s1)

		if !reflect.DeepEqual(s1, s2) {
			t.Errorf("Deepcopied slice %v is not equal to original %v", s2, s1)
		}

		// Modify original and check if clone is affected
		s1[0].Value = 99
		*s1[0].Ptr = 999

		if s2[0].Value == 99 || *s2[0].Ptr == 999 {
			t.Errorf("Deepcopy is not a deep copy; it was modified when the original changed.")
		}
	})
}
