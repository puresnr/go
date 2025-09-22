package prand

import (
	"errors"
	"math"
	"math/rand/v2"
	"sort"
)

// Epsilon defines the tolerance for floating-point comparisons.
const Epsilon = 1e-9

var (
	ErrorEmptyProbs      = errors.New("empty probs")
	ErrorInvalidSumProbs = errors.New("sum probs should be 1")
	ErrorInvalidProbs    = errors.New("invalid probs")
)

type Probs []float64

// RandIdx returns a random index based on the cumulative probability distribution.
// It uses binary search and has a time complexity of O(log N).
func (p Probs) RandIdx() (int, error) {
	lenp := len(p)
	// Check if the distribution is valid. The sum (last element) must be 1.0 within tolerance.
	if lenp == 0 || math.Abs(p[lenp-1]-1.0) > Epsilon {
		return 0, ErrorInvalidProbs
	}

	value := rand.Float64()

	// sort.Search finds the smallest index i for which p[i] > value.
	// This is the definition of which bucket the random value falls into.
	idx := sort.Search(lenp, func(i int) bool { return p[i] > value })

	// If idx is lenp, it means the value is >= the last element in Probs
	// (which can happen if the sum is slightly less than 1.0).
	// In that case, the value belongs to the last bucket.
	if idx == lenp {
		return lenp - 1, nil
	}

	return idx, nil
}

// NewProbs creates a cumulative probability distribution from a slice of raw probabilities.
// The sum of rawprobs must be equal to 1.0 within a tolerance of Epsilon.
func NewProbs(rawprobs []float64) (Probs, error) {
	lenProbs := len(rawprobs)

	if lenProbs == 0 {
		return nil, ErrorEmptyProbs
	}

	probs := make(Probs, lenProbs)
	probs[0] = rawprobs[0]
	for i := 1; i < lenProbs; i++ {
		probs[i] = probs[i-1] + rawprobs[i]
	}

	if math.Abs(probs[lenProbs-1]-1.0) > Epsilon {
		return nil, ErrorInvalidSumProbs
	}
	return probs, nil
}
