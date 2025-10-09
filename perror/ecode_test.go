package perror

import (
	"errors"
	"fmt"
	"testing"

	"github.com/puresnr/go/perror/ecode"
	"github.com/stretchr/testify/assert"
)

// anotherError is a distinct error type for testing against Ecode.
type anotherError struct{ error }

func TestEcode(t *testing.T) {
	t.Run("t-nil s-nil unwrap same-type", func(t *testing.T) {
		var serr, terr *ecode.Ecode
		assert.True(t, errors.Is(serr, terr), "t-nil s-nil unwrap same-type should be equal")
	})

	t.Run("t-nil s-nil unwrap diff-type", func(t *testing.T) {
		var serr *ecode.Ecode
		var terr error
		assert.False(t, errors.Is(serr, terr), "t-nil s-nil unwrap diff-type should not be equal")
	})

	t.Run("t-nil s-not-nil wrap same-type", func(t *testing.T) {
		var terr *ecode.Ecode
		serr := Wrap(ecode.New("1001"))
		assert.False(t, errors.Is(serr, terr), "t-nil s-not-nil wrap same-type should not be equal")
	})

	t.Run("t-nil s-not-nil wrap diff-type", func(t *testing.T) {
		var terr error
		serr := Wrap(ecode.New("1001"))
		assert.False(t, errors.Is(serr, terr), "t-nil s-not-nil wrap diff-type should not be equal")
	})

	t.Run("t-nil s-not-nil unwrap same-type", func(t *testing.T) {
		var terr *ecode.Ecode
		serr := ecode.New("1001")
		assert.False(t, errors.Is(serr, terr), "t-nil s-not-nil unwrap same-type should not be equal")
	})

	t.Run("t-nil s-not-nil unwrap diff-type", func(t *testing.T) {
		var terr error
		serr := ecode.New("1001")
		assert.False(t, errors.Is(serr, terr), "t-nil s-not-nil unwrap diff-type should not be equal")
	})

	t.Run("t--not-nil s-nil unwrap same-type", func(t *testing.T) {
		var serr *ecode.Ecode
		terr := ecode.New("1001")
		assert.False(t, errors.Is(serr, terr), "t--not-nil s-nil unwrap same-type should not be equal")
	})

		t.Run("t--not-nil s-nil unwrap diff-type", func(t *testing.T) {
		var serr *ecode.Ecode
		terr := errors.New("1001")
		assert.False(t, errors.Is(serr, terr), "t--not-nil s-nil unwrap diff-type should not be equal")
	})

			t.Run("t--not-nil s-not-nil wrap same-type same-value same-code", func(t *testing.T) {
				tmp := ecode.New("1001")
		serr := Wrap(tmp)
		terr := tmp
		assert.True(t, errors.Is(serr, terr), "t--not-nil s-not-nil wrap same-type same-value same-code should be equal")
	})

				t.Run("t--not-nil s-not-nil wrap same-type diff-value same-code", func(t *testing.T) {
		serr := Wrap(ecode.New("1001"))
		terr := ecode.New("1001")
		assert.True(t, errors.Is(serr, terr), "t--not-nil s-not-nil wrap same-type diff-value same-code should be equal")
	})

				t.Run("t--not-nil s-not-nil wrap same-type diff-value diff-code", func(t *testing.T) {
serr := Wrap(ecode.New("1001"))
		terr := ecode.New("1002")
		assert.False(t, errors.Is(serr, terr), "t--not-nil s-not-nil wrap same-type diff-value diff-code should be equal")
	})

					t.Run("t--not-nil s-not-nil wrap diff-type diff-value same-code", func(t *testing.T) {
		serr := Wrap(ecode.New("1001"))
		terr := errors.New("1001")
		assert.False(t, errors.Is(serr, terr), "t--not-nil s-not-nil wrap diff-type diff-value same-code should not be equal")
	})

				t.Run("t--not-nil s-not-nil wrap diff-type diff-value diff-code", func(t *testing.T) {
serr := Wrap(ecode.New("1001"))
		terr := errors.New("1002")
		assert.False(t, errors.Is(serr, terr), "t--not-nil s-not-nil wrap diff-type diff-value diff-code should not be equal")
	})

				t.Run("t--not-nil s-not-nil unwrap same-type same-value same-code", func(t *testing.T) {
				tmp := ecode.New("1001")
		serr := tmp
		terr := tmp
		assert.True(t, errors.Is(serr, terr), "t--not-nil s-not-nil unwrap same-type same-value same-code should be equal")
	})

				t.Run("t--not-nil s-not-nil unwrap same-type diff-value same-code", func(t *testing.T) {
		serr := ecode.New("1001")
		terr := ecode.New("1001")
		assert.True(t, errors.Is(serr, terr), "t--not-nil s-not-nil unwrap same-type diff-value same-code should be equal")
	})

				t.Run("t--not-nil s-not-nil unwrap same-type diff-value diff-code", func(t *testing.T) {
serr := ecode.New("1001")
		terr := ecode.New("1002")
		assert.False(t, errors.Is(serr, terr), "t--not-nil s-not-nil unwrap same-type diff-value diff-code should be equal")
	})

					t.Run("t--not-nil s-not-nil unwrap diff-type diff-value same-code", func(t *testing.T) {
		serr := ecode.New("1001")
		terr := errors.New("1001")
		assert.False(t, errors.Is(serr, terr), "t--not-nil s-not-nil unwrap diff-type diff-value same-code should not be equal")
	})

				t.Run("t--not-nil s-not-nil unwrap diff-type diff-value diff-code", func(t *testing.T) {
serr := ecode.New("1001")
		terr := errors.New("1002")
		assert.False(t, errors.Is(serr, terr), "t--not-nil s-not-nil unwrap diff-type diff-value diff-code should not be equal")
	})

	t.Run("Same Ecode", func(t *testing.T) {
		e1 := ecode.New("1001")
		e2 := ecode.New("1001")
		assert.True(t, errors.Is(e1, e2), "Two Ecode instances with the same code should be equal")
	})

	t.Run("Different Ecode", func(t *testing.T) {
		e1 := ecode.New("1001")
		e2 := ecode.New("1002")
		assert.False(t, errors.Is(e1, e2), "Two Ecode instances with different codes should not be equal")
	})

	t.Run("Compare with other error types", func(t *testing.T) {
		e1 := ecode.New("1001")
		e2 := errors.New("1001")
		assert.False(t, errors.Is(e1, e2), "Ecode should not be equal to a standard error with the same message")
	})

	t.Run("Wrapped Ecode", func(t *testing.T) {
		e1 := ecode.New("1001")
		e2 := ecode.New("1001")
		wrappedErr := fmt.Errorf("wrapped error: %w", e1)
		assert.True(t, errors.Is(wrappedErr, e2), "errors.Is should find an Ecode within a wrapped error")
	})

	t.Run("Wrapped different Ecode", func(t *testing.T) {
		e1 := ecode.New("1001")
		e2 := ecode.New("1002")
		wrappedErr := fmt.Errorf("wrapped error: %w", e1)
		assert.False(t, errors.Is(wrappedErr, e2), "errors.Is should not find a different Ecode within a wrapped error")
	})

	t.Run("Nil handling", func(t *testing.T) {
		var typedNil *ecode.Ecode
		e1 := ecode.New("1001")

		// errors.Is(err, nil) is defined as err == nil
		assert.False(t, errors.Is(e1, nil), "A non-nil error should not be equal to an untyped nil")
		assert.True(t, errors.Is(nil, nil), "An untyped nil should be equal to an untyped nil")

		// A typed nil is not equal to an untyped nil
		assert.False(t, errors.Is(typedNil, nil), "A typed nil should not be equal to an untyped nil")

		// Test our Is method's logic with typed nils
		assert.False(t, errors.Is(e1, typedNil), "A non-nil Ecode should not be equal to a typed nil Ecode")
		assert.False(t, errors.Is(typedNil, e1), "A typed nil Ecode should not be equal to a non-nil Ecode")
		assert.True(t, errors.Is(typedNil, typedNil), "A typed nil Ecode should be equal to another typed nil Ecode")
	})

	t.Run("Error method", func(t *testing.T) {
		e := ecode.New("my-error-code")
		assert.Equal(t, "my-error-code", e.Error(), "Error() method should return the code string")
	})

	t.Run("Two different typed nils", func(t *testing.T) {
		var ecodeNil *ecode.Ecode
		var anotherErrNil *anotherError

		// errors.Is should return false because the types are different,
		// and the Is method on Ecode will fail the type assertion.
		assert.False(t, errors.Is(ecodeNil, anotherErrNil))
		assert.False(t, errors.Is(anotherErrNil, ecodeNil))
	})
}
