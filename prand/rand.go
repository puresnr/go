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

// Probs is a cumulative probability distribution.
// It is an opaque type and can only be created by NewProbs.
type Probs struct {
	cumulative []float64
}

// RandIdx returns a random index based on the cumulative probability distribution.
// It uses binary search and has a time complexity of O(log N).
func (p *Probs) RandIdx() (int, error) {
	if p == nil {
		return 0, ErrorInvalidProbs
	}

	lenp := len(p.cumulative)
	value := rand.Float64()

	// sort.Search finds the smallest index i for which p.cumulative[i] > value.
	// This is the definition of which bucket the random value falls into.
	idx := sort.Search(lenp, func(i int) bool { return p.cumulative[i] > value })

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
func NewProbs(rawprobs []float64) (*Probs, error) {
	lenProbs := len(rawprobs)

	if lenProbs == 0 {
		return nil, ErrorEmptyProbs
	}

	cumulative := make([]float64, lenProbs)
	cumulative[0] = rawprobs[0]
	for i := 1; i < lenProbs; i++ {
		cumulative[i] = cumulative[i-1] + rawprobs[i]
	}

	if math.Abs(cumulative[lenProbs-1]-1.0) > Epsilon {
		return nil, ErrorInvalidSumProbs
	}
	return &Probs{cumulative: cumulative}, nil
}