package prand

import (
	"math"
	
	"testing"
)

func TestNewProbs(t *testing.T) {
	testCases := []struct {
		name     string
		rawprobs []float64
		want     Probs
		wantErr  error
	}{
		{
			name:     "valid probabilities",
			rawprobs: []float64{0.1, 0.2, 0.7},
			want:     Probs{0.1, 0.3, 1.0},
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
			want:     Probs{1.0},
			wantErr:  nil,
		},
		{
			name:     "single element, invalid sum",
			rawprobs: []float64{0.9},
			want:     nil,
			wantErr:  ErrorInvalidSumProbs,
		},
		{
			name:     "probabilities with zero",
			rawprobs: []float64{0.5, 0, 0.5},
			want:     Probs{0.5, 0.5, 1.0},
			wantErr:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewProbs(tc.rawprobs)
			if err != tc.wantErr {
				t.Errorf("NewProbs() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			// Use a tolerance for float comparison
			if len(got) != len(tc.want) {
				t.Errorf("NewProbs() len = %d, want len %d", len(got), len(tc.want))
			}
			for i := range got {
				if math.Abs(got[i]-tc.want[i]) > 1e-9 {
					t.Errorf("NewProbs() = %v, want %v", got, tc.want)
					break
				}
			}
		})
	}
}

func TestRandIdx(t *testing.T) {
	t.Run("invalid probs - empty", func(t *testing.T) {
		p := Probs{}
		_, err := p.RandIdx()
		if err != ErrorInvalidProbs {
			t.Errorf("RandIdx() with empty probs should return ErrorInvalidProbs, got %v", err)
		}
	})

	t.Run("invalid probs - sum not 1", func(t *testing.T) {
		p := Probs{0.1, 0.5}
		_, err := p.RandIdx()
		if err != ErrorInvalidProbs {
			t.Errorf("RandIdx() with invalid sum should return ErrorInvalidProbs, got %v", err)
		}
	})

	t.Run("valid probs - check bounds", func(t *testing.T) {
		p, err := NewProbs([]float64{0.1, 0.2, 0.3, 0.4})
		if err != nil {
			t.Fatalf("Failed to create probs for testing: %v", err)
		}
		for i := 0; i < 1000; i++ {
			idx, err := p.RandIdx()
			if err != nil {
				t.Fatalf("RandIdx() returned an unexpected error: %v", err)
			}
			if idx < 0 || idx >= len(p) {
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
