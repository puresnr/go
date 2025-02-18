package cast

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCastInt(t *testing.T) {
	utar, err := CastInt[uint, uint](0)
	assert.Equal(t, uint(0), utar)
	assert.Nil(t, err)

	itar, err := CastInt[uint, int](0)
	assert.Equal(t, 0, itar)
	assert.Nil(t, err)

	itar, _ = CastInt[uint, int](math.MaxUint32)
	assert.Equal(t, math.MaxUint32, itar)

	utar, _ = CastInt[uint, uint](math.MaxUint64)
	assert.Equal(t, uint(math.MaxUint64), utar)

	_, err = CastInt[uint, int](math.MaxUint64)
	assert.Equal(t, ECastIntOverflow, err)

	_, err = CastInt[uint, uint32](math.MaxUint64)
	assert.Equal(t, ECastIntOverflow, err)

	itar, err = CastInt[int, int](0)
	assert.Equal(t, 0, itar)
	assert.Nil(t, err)

	utar, err = CastInt[int, uint](0)
	assert.Equal(t, uint(0), utar)
	assert.Nil(t, err)

	itar, _ = CastInt[int, int](-math.MaxUint32)
	assert.Equal(t, -math.MaxUint32, itar)

	_, err = CastInt[int, uint](-math.MaxUint32)
	assert.Equal(t, ECastIntOverflow, err)

	_, err = CastInt[int, int8](-math.MaxUint32)
	assert.Equal(t, ECastIntOverflow, err)

	itar, _ = CastInt[int, int](math.MaxUint32)
	assert.Equal(t, math.MaxUint32, itar)

	utar, _ = CastInt[int, uint](math.MaxUint32)
	assert.Equal(t, uint(math.MaxUint32), utar)

	_, err = CastInt[int, int8](math.MaxUint32)
	assert.Equal(t, ECastIntOverflow, err)

	_, err = CastInt[int, uint8](math.MaxUint32)
	assert.Equal(t, ECastIntOverflow, err)
}
