package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPivot(t *testing.T) {
	instances := GenerateRandomSlices(30)
	pivot := findPivot(instances, 0, len(instances)-1)

	t.Run("asc", func(t *testing.T) {
		for _, v := range instances[:pivot] {
			assert.True(t, v <= instances[pivot])
		}
	})

	t.Run("desc", func(t *testing.T) {
		for i := pivot + 1; i < len(instances); i++ {
			assert.True(t, instances[i] >= instances[pivot])
		}
	})
}

func TestQuickSort(t *testing.T) {
	instances := GenerateRandomSlices(30)
	quickSort(instances, 0, len(instances)-1)
	assert.True(t, IsSorted(instances))
}
