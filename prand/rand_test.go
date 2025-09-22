package prand

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewWeightedRand(t *testing.T) {
	testCases := []struct {
		name     string
		values   []any
		rawprobs []float64
		want     *WeightedRand[any]
		wantErr  error
	}{
		{
			name:     "valid probabilities with integers",
			values:   []any{10, 20, 70},
			rawprobs: []float64{0.1, 0.2, 0.7},
			want: &WeightedRand[any]{
				values:     []any{10, 20, 70},
				cumulative: []float64{0.1, 0.3, 1.0},
			},
			wantErr: nil,
		},
		{
			name:     "valid probabilities with strings",
			values:   []any{"a", "b", "c"},
			rawprobs: []float64{0.1, 0.2, 0.7},
			want: &WeightedRand[any]{
				values:     []any{"a", "b", "c"},
				cumulative: []float64{0.1, 0.3, 1.0},
			},
			wantErr: nil,
		},
		{
			name:     "mismatched lengths",
			values:   []any{1, 2},
			rawprobs: []float64{0.1, 0.2, 0.7},
			wantErr:  ErrorMismatchedLen,
		},
		{
			name:     "empty probabilities",
			values:   []any{},
			rawprobs: []float64{},
			wantErr:  ErrorEmptyProbs,
		},
		{
			name:     "invalid sum of probabilities",
			values:   []any{1, 2},
			rawprobs: []float64{0.1, 0.2},
			wantErr:  ErrorInvalidSumProbs,
		},
		{
			name:     "single element, valid",
			values:   []any{"hello"},
			rawprobs: []float64{1.0},
			want: &WeightedRand[any]{
				values:     []any{"hello"},
				cumulative: []float64{1.0},
			},
			wantErr: nil,
		},
		{
			name:     "negative probability",
			values:   []any{1, 2, 3},
			rawprobs: []float64{0.5, -0.1, 0.6},
			wantErr:  ErrorNegativeProb,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NewWeightedRand(tc.values, tc.rawprobs)
			assert.Equal(t, tc.wantErr, err)

			if tc.wantErr != nil {
				return
			}
			require.NotNil(t, got)
			assert.Equal(t, tc.want.values, got.values)
			assert.Equal(t, len(tc.want.cumulative), len(got.cumulative))
			for i := range got.cumulative {
				assert.InDelta(t, tc.want.cumulative[i], got.cumulative[i], Epsilon, "cumulative probabilities at index %d do not match", i)
			}
		})
	}
}

func TestWeightedRand_Rand(t *testing.T) {
	values := []string{"a", "b", "c"}
	rawprobs := []float64{0.2, 0.3, 0.5}
	p, err := NewWeightedRand(values, rawprobs)
	require.NoError(t, err)

	const iterations = 100000
	counts := make(map[string]int)
	for i := 0; i < iterations; i++ {
		v, err := p.Rand()
		require.NoError(t, err)
		counts[v]++
	}

	for i, val := range values {
		got := float64(counts[val]) / float64(iterations)
		want := rawprobs[i]
		assert.InDelta(t, want, got, 0.01, fmt.Sprintf("Statistical distribution for value %q is out of tolerance", val))
	}
}

func TestWeightedRand_RandIdx(t *testing.T) {
	t.Run("nil receiver", func(t *testing.T) {
		var p *WeightedRand[int]
		_, err := p.RandIdx()
		assert.Equal(t, ErrorInvalidProbs, err)
	})

	t.Run("valid probs - check bounds", func(t *testing.T) {
		values := []int{0, 1, 2, 3}
		rawprobs := []float64{0.1, 0.2, 0.3, 0.4}
		p, err := NewWeightedRand(values, rawprobs)
		require.NoError(t, err)

		for i := 0; i < 1000; i++ {
			idx, err := p.RandIdx()
			require.NoError(t, err)
			assert.GreaterOrEqual(t, idx, 0)
			assert.Less(t, idx, len(values))
		}
	})

	t.Run("statistical distribution", func(t *testing.T) {
		values := []int{0, 1, 2}
		rawprobs := []float64{0.2, 0.3, 0.5}
		p, err := NewWeightedRand(values, rawprobs)
		require.NoError(t, err)

		const iterations = 100000
		counts := make([]int, len(values))
		for i := 0; i < iterations; i++ {
			idx, _ := p.RandIdx()
			counts[idx]++
		}

		for i, count := range counts {
			got := float64(count) / float64(iterations)
			want := rawprobs[i]
			assert.InDelta(t, want, got, 0.01, "Statistical distribution for index %d is out of tolerance", i)
		}
	})
}

func TestWeightedRand_Rand_Empty(t *testing.T) {
	p := &WeightedRand[any]{
		values:     nil,
		cumulative: []float64{},
	}
	_, err := p.Rand()
	assert.Error(t, err)
}
