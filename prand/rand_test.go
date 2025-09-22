package prand

import (
	"math"
	"testing"
)

func TestNewProbs(t *testing.T) {
	testCases := []struct {
		name     string
		rawprobs []float64
		want     *Probs
		wantErr  error
	}{
		{
			name:     "valid probabilities",
			rawprobs: []float64{0.1, 0.2, 0.7},
			want:     &Probs{cumulative: []float64{0.1, 0.3, 1.0}},
			wantErr:  nil,
		},
		{
			name:     "empty probabilities",
			rawprobs: []float64{},
			want:     nil,
			wantErr:  ErrorEmptyProbs,
		},
		{
			name:     "invalid sum of probabilities",
			rawprobs: []float64{0.1, 0.2},
			want:     nil,
			wantErr:  ErrorInvalidSumProbs,
		},
		{
			name:     "single element, valid",
			rawprobs: []float64{1.0},
			want:     &Probs{cumulative: []float64{1.0}},
			wantErr:  nil,
		},
		{
			name:     "sum slightly under 1, within tolerance",
			rawprobs: []float64{0.5, 0.5 - Epsilon/2},
			want:     &Probs{cumulative: []float64{0.5, 1.0 - Epsilon/2}},
			wantErr:  nil,
		},
		{
			name:     "negative probability",
			rawprobs: []float64{0.5, -0.1, 0.6},
			want:     nil,
			wantErr:  ErrorNegativeProb,
		},
		{
			name:     "negative probability at start",
			rawprobs: []float64{-0.1, 0.5, 0.6},
			want:     nil,
			wantErr:  ErrorNegativeProb,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewProbs(tc.rawprobs)
			if err != tc.wantErr {
				t.Errorf("NewProbs() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if tc.wantErr != nil {
				return
			}
			if got == nil {
				t.Fatalf("NewProbs() = nil, want %v", tc.want)
			}
			if len(got.cumulative) != len(tc.want.cumulative) {
				t.Fatalf("NewProbs() len = %d, want len %d", len(got.cumulative), len(tc.want.cumulative))
			}
			for i := range got.cumulative {
				if math.Abs(got.cumulative[i]-tc.want.cumulative[i]) > Epsilon {
					t.Errorf("NewProbs() = %v, want %v", got, tc.want)
					break
				}
			}
		})
	}
}

func TestRandIdx(t *testing.T) {
	t.Run("nil probs", func(t *testing.T) {
		var p *Probs
		_, err := p.RandIdx()
		if err != ErrorInvalidProbs {
			t.Errorf("RandIdx() with nil probs should return ErrorInvalidProbs, got %v", err)
		}
	})

	t.Run("valid probs - check bounds", func(t *testing.T) {
		rawprobs := []float64{0.1, 0.2, 0.3, 0.4}
		p, err := NewProbs(rawprobs)
		if err != nil {
			t.Fatalf("Failed to create probs for testing: %v", err)
		}
		for i := 0; i < 1000; i++ {
			idx, err := p.RandIdx()
			if err != nil {
				t.Fatalf("RandIdx() returned an unexpected error: %v", err)
			}
			if idx < 0 || idx >= len(rawprobs) {
				t.Fatalf("RandIdx() returned an out-of-bounds index: %d", idx)
			}
		}
	})

	t.Run("statistical distribution", func(t *testing.T) {
		rawprobs := []float64{0.2, 0.3, 0.5}
		p, err := NewProbs(rawprobs)
		if err != nil {
			t.Fatalf("Failed to create probs for testing: %v", err)
		}

		const iterations = 100000
		counts := make([]int, len(rawprobs))
		for i := 0; i < iterations; i++ {
			idx, _ := p.RandIdx()
			counts[idx]++
		}

		for i, count := range counts {
			got := float64(count) / float64(iterations)
			want := rawprobs[i]
			// Allow a tolerance for randomness
			tolerance := 0.01
			if got < want-tolerance || got > want+tolerance {
				t.Errorf("Statistical distribution for index %d is out of tolerance. got: %f, want: %f", i, got, want)
			}
		}
	})
}