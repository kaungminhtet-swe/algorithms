package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testcases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "merge two elements",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "merge three elements",
			input:    []int{3, 2, 1},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range testcases {
		actual := MergeSort(tc.input)

		assert.Equal(t, tc.expected, actual)
	}
}

func TestMergeSortWithRandomData(t *testing.T) {
	instances := GenerateRandomSlices(100000)
	assert.True(t, IsSorted(MergeSort(instances)))
}
