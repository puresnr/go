package basic

import (
	"reflect"
	"testing"
)

func TestDeepcopy(t *testing.T) {
	t.Run("nil slice", func(t *testing.T) {
		var s1 []int
		s2 := Deepcopy(s1)
		if s2 != nil {
			t.Errorf("Deepcopy(nil) should be nil, got %v", s2)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		s1 := []int{}
		s2 := Deepcopy(s1)
		if len(s2) != 0 {
			t.Errorf("Clone of empty slice should be an empty slice, got %v", s2)
		}
	})

	t.Run("int slice", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := Deepcopy(s1)

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
		s2 := Deepcopy(s1)

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
