package filters_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superlinkx/go-skeleton/services/messages/internal/filters"
	"golang.org/x/exp/constraints"
)

type TestCase[T constraints.Integer] struct {
	Name  string
	Input []T
	Want  []T
}

func TestFilterToOddIntegers(t *testing.T) {
	var (
		testCases = []TestCase[int64]{
			{
				Name:  "empty",
				Input: []int64{},
				Want:  []int64{},
			},
			{
				Name:  "allOdd",
				Input: []int64{1, 3, 5, 7, 9},
				Want:  []int64{1, 3, 5, 7, 9},
			},
			{
				Name:  "allEven",
				Input: []int64{2, 4, 6, 8, 10},
				Want:  []int64{},
			},
			{
				Name:  "sequential",
				Input: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				Want:  []int64{1, 3, 5, 7, 9},
			},
			{
				Name:  "nonSequential",
				Input: []int64{10, 6, 2, 34, 5, 6, 2, 34, 1213, 13, 4, 56},
				Want:  []int64{5, 1213, 13},
			},
		}
	)

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			filtered := filters.FilterToOddIntegers(tc.Input)
			assert.Exactly(t, tc.Want, filtered)
		})
	}
}
