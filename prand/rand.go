package prand

import (
	"errors"
	"math/rand/v2"
)

var (
	ErrorEmptyProbs      = errors.New("empty probs")
	ErrorInvalidSumProbs = errors.New("sum probs should be 1")
	ErrorInvalidProbs    = errors.New("invalid probs")
)

type Probs []float64

func (p Probs) RandIdx() (int, error) {
	if lenp := len(p); lenp == 0 || p[lenp-1] != 1 {
		return 0, ErrorInvalidProbs
	}

	value := rand.Float64()

	for idx, v := range p {
		if value < v {
			return idx, nil
		}
	}

	// should not be reached
	return 0, ErrorInvalidProbs
}

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

	if probs[lenProbs-1] != 1 {
		return nil, ErrorInvalidSumProbs
	}
	return probs, nil
}

