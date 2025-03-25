package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSorted(t *testing.T) {
	instances := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.True(t, IsSorted(instances))

	instances = []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 9}
	assert.False(t, IsSorted(instances))
}
