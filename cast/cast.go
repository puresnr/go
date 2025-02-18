package cast

import (
	"errors"
	"golang.org/x/exp/constraints"
)

var ECastIntOverflow = errors.New("overflow")

func CastInt[S, T constraints.Integer](src S) (T, error) {
	tar := T(src)

	if S(tar) != src || (src < 0 && tar > 0) || (src > 0 && tar < 0) {
		return 0, ECastIntOverflow
	}
	return tar, nil
}
